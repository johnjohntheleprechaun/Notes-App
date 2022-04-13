package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

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