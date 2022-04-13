package main

import (
	"net/http"
	"log"

	"github.com/gin-gonic/gin"
)

type SaveData struct {
	Session int
	ID      int
	Content string
}

func initializeRouter() *gin.Engine {
	//Router setup
	router := gin.Default()
	router.SetTrustedProxies(nil)
	router.LoadHTMLGlob("templates/*")
	router.Static("/static", "./static")
	//Endpoint setup
	router.GET("/", func(ctx *gin.Context) { //Homepage
		ctx.HTML(http.StatusOK, "homepage.html", nil)
	})
	router.POST("/save", func(ctx *gin.Context) {
		var saveData SaveData
		ctx.BindJSON(&saveData)
	})
	return router
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}