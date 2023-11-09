package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Test struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	//ping
	r.GET("/ping", func(ctx *gin.Context) {
		//ctx.String(http.StatusOK, "hello word")
		ctx.JSON(http.StatusOK, gin.H{
			"status": "succes",
			"value":  "hello word",
		})
	})

	v1 := r.Group("v1")
	v1.GET("/user/:name", func(ctx *gin.Context) {
		param := ctx.Param("name")
		ctx.JSON(http.StatusOK, gin.H{
			"status": "succes",
			"value":  param,
		})
	})

	v1.POST("user", func(ctx *gin.Context) {
		var data Test

		ctx.BindJSON(&data)

		ctx.JSON(http.StatusOK, gin.H{
			"status": "succes",
			"value":  data,
		})
	})

	v1.GET("/user", func(ctx *gin.Context) {
		Query := ctx.Query("name")
		ctx.JSON(http.StatusOK, gin.H{
			"status": "succes",
			"value":  Query,
		})
	})

	return r
}

func main() {
	r := setupRouter()

	r.Run(":8080")
}
