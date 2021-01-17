package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	log "github.com/sirupsen/logrus"
	"github.com/urlshortener/proto"
)

// URLModel defines the url_models schema / table
type URLModel struct {

	// this is a pointer to the default gorm model that implements timestamps and AUTO_INC_ID
	*gorm.Model

	// this is a pointer to the protobuf model that is auto generated with the script in /scripts
	*proto.Url
}

func main() {

	port := "5050"
	host := "localhost"
	if os.Getenv("SERVICE_PORT") != "" {
		port = os.Getenv("SERVICE_PORT")
	}
	if os.Getenv("GIN_MODE") == "release" {
		gin.SetMode(gin.ReleaseMode)
	}
	const addrs = "postgres://root@localhost:26257/urls?sslmode=disable"
	db, err := gorm.Open("postgres", addrs)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Info("Connected to cockroach cluster")
	}

	defer db.Close()
	db.LogMode(false)
	db.AutoMigrate(&URLModel{})

	router := gin.Default()
	router.Use(cors.Default())
	router.GET("/:url", func(ctx *gin.Context) {
		url := &URLModel{}
		ctxURL := fmt.Sprintf("http://%v:%v%v", host, port, ctx.Request.URL.String())
		found := db.Select("real").Where("generated = ?", ctxURL).First(url)
		if found.RowsAffected == 1 {
			ctx.Redirect(http.StatusMovedPermanently, found.Value.(*URLModel).Real)
			return
		}
		ctx.JSON(http.StatusNotFound, gin.H{
			"status": http.StatusNotFound,
			"query":  ctxURL,
		})
		return
	})

	router.POST("/save", func(ctx *gin.Context) {
		url := &URLModel{}
		ctx.BindJSON(url)
		found := db.Select("generated").Where("real=? AND generated=?", url.Real, url.Generated).First(url)
		if found.RowsAffected == 1 {
			ctx.JSON(http.StatusOK, gin.H{
				"status": http.StatusFound,
				"tip":    fmt.Sprintf("try %v to use it", url.Generated),
			})
			return
		}
		db.Create(url)
		ctx.JSON(http.StatusOK, gin.H{
			"generated": url.GetGenerated(),
			"ip_v4":     url.GetIpV4(),
			"real":      url.GetReal(),
		})
	})

	log.Fatal(router.Run(fmt.Sprintf(":%v", port)))
}
