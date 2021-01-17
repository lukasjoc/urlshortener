package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// Url ...
type Url struct {
	gorm.Model
	Real, Generated string
}

func main() {

	const addr = "postgres://urlshortener@localhost:26257/urlshortener?sslmode=require"
	db, err := gorm.Open("postgres", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	db.LogMode(false)
	db.AutoMigrate(&Url{})

	router := gin.Default()

	router.GET("/save", func(ctx *gin.Context) {
		db.Create(&Url{})
	})

	logrus.Fatal(router.Run(fmt.Sprintf(":%d", 5050)))
}
