package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/sirupsen/logrus"
	"gitub.com/urlshortener/proto"
)

// URLModel ...
type URLModel struct {
	gorm.Model
	// proto.Url
}

// test http://localhost:5050/Tr-8mZNJ
func main() {
	const addrs = "postgres://root:admin@127.0.0.1:53232/urls?sslmode=require"
	db, err := gorm.Open("postgres", addrs)
	if err != nil {
		logrus.Fatal(err)
	} else {
		logrus.Info("Connected to cockroach db")
	}

	defer db.Close()
	db.LogMode(false)
	db.AutoMigrate(&URLModel{})

	router := gin.Default()
	router.GET("/:url", func(ctx *gin.Context) {
		// TODO:  find url
		ctx.JSON(http.StatusNotImplemented, gin.H{
			"message": "Not Implemented",
		})
	})

	router.PUT("/save", func(ctx *gin.Context) {
		url := &proto.Url{}
		ctx.BindJSON(url)
		db.Create(url)
		ctx.JSON(http.StatusOK, gin.H{
			"generated": url.GetGenerated(),
			"ip_v4":     url.GetIpV4(),
			"real":      url.GetReal(),
		})
	})

	if os.Getenv("GIN_MODE") == "release" {
		gin.SetMode(gin.ReleaseMode)
	}
	port := "5050"
	if os.Getenv("SERVICE_PORT") != "" {
		port = os.Getenv("SERVICE_PORT")
	}
	logrus.Fatal(router.Run(fmt.Sprintf(":%v", port)))
}
