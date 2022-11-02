package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func hello(c *gin.Context) {
	c.JSONP(http.StatusOK, gin.H{
		"message": "hello",
	})
}

func allArticles(c *gin.Context) {
	articles, err := SelectAllDb()
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(http.StatusOK, articles)
}

func NewHTTPServer() {
	r := gin.Default()
	r.GET("/db", allArticles)
	r.GET("/", hello)

	r.Run()
}
