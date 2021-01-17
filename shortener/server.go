package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"gitub.com/urlshortener/proto"
)

// URLModel
type URLModel struct {
	gorm.Model
	proto.Url
}

// test http://localhost:5050/Tr-8mZNJ
func main() {
	router := gin.Default()
	router.GET("/:url", func(ctx *gin.Context) {
		// TODO:  find url
		ctx.JSON(http.StatusNotImplemented, gin.H{
			"message": "Not Implemented",
		})
	})
	router.PUT("/save", func(ctx *gin.Context) {
		// TODO: save to db
		url := &proto.Url{}
		ctx.BindJSON(url)
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
