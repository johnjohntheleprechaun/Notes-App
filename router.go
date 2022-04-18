package main

import (
	"fmt"
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

	router.GET("/:username/notes/*noteID", func(ctx *gin.Context) {
		username := ctx.Param("username")
		note := ctx.Param("noteID")
		authToken, _ := ctx.Cookie("authToken")
		if note == "/" {
			ctx.HTML(http.StatusOK, "notes.html", gin.H{
				"notes":    nil,
				"username": username,
			})
		} else {
			if authToken == "" {
				ctx.Redirect(http.StatusFound, "/login")
			}
			ctx.JSON(http.StatusOK, gin.H{
				"test":      note,
				"authToken": authToken,
			})
		}
	})

	router.GET("/login", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "login.html", nil)
	})
	router.POST("/login/submit", func(ctx *gin.Context) {
		fmt.Println(ctx.PostForm("username"))
		//Get an auth token
		//Set the cookie
	})
	return router
}
