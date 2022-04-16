package main

import (
	"net/http"

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

	router.GET("/:user/notes/*note", func(ctx *gin.Context) {
		_ = ctx.Param("user")
		note := ctx.Param("note")
		if note == "/" {
			ctx.HTML(http.StatusOK, "notes.html", gin.H{
				"notes": nil,
			})
		}
	})
	return router
}
