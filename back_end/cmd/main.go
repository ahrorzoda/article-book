package main

import (
	"github.com/ahrorzoda/article-book/back_end/internal/api"
	"github.com/ahrorzoda/article-book/back_end/internal/cors"
	"github.com/ahrorzoda/article-book/back_end/internal/file"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	rout := gin.Default()
	rout.Use(cors.CorsMiddleware)
	rout.GET("/", func(c *gin.Context) { c.HTML(http.StatusOK, "humanitarian.html", nil) })
	rout.GET("/sc", func(c *gin.Context) { c.HTML(http.StatusOK, "science.html", nil) })
	rout.POST("/humanities", api.AddHumanitarianBook) //TODO: ok
	rout.POST("/science", api.AddScienceBook)         //TODO: ok
	rout.GET("/book/cover", file.FindFile)            //TODO: ok
	err := rout.Run(":8080")
	if err != nil {
		return
	}
}
