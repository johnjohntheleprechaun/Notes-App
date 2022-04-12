package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	//Router setup
	router := gin.Default()
	router.SetTrustedProxies(nil)
	router.LoadHTMLGlob("templates/*")
	router.Static("/static", "./static")
	//Endpoint setup
	router.GET("/", func(ctx *gin.Context){
		ctx.HTML(http.StatusOK, "homepage.html", nil)
	})
	router.Run()
}