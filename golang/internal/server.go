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

func getAllArticles(c *gin.Context) {
	articles, err := SelectAllArticle()
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(http.StatusOK, articles)
}

func getOneArticle(c *gin.Context) {

	article, err := SelectOneArticle(c.Param("article_id"))
	if err != nil {
		c.JSONP(http.StatusBadRequest, gin.H{
			"message": "no_content",
		})
	} else {
		c.JSON(http.StatusOK, article)
	}

}

func allTitle(c *gin.Context) {
	titles, err := SelectAllTitleAndID()
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(http.StatusOK, titles)
}

func addNewArticle(c *gin.Context) {
	var json ReceiveJson
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSONP(http.StatusBadRequest, gin.H{
			"message": "invalid json scheme",
		})
	}
	if err := InsertNewArticle(json); err != nil {
		c.JSONP(http.StatusBadRequest, gin.H{
			"message": "db error",
		})
	} else {
		c.JSON(http.StatusOK, json)
	}

}

func editArticle(c *gin.Context) {
	var json ReceiveJson
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSONP(http.StatusBadRequest, gin.H{
			"message": "invalid json scheme",
		})
	}
	if err := EditArticle(c.Param("article_id"), json); err != nil {
		c.JSONP(http.StatusBadRequest, gin.H{
			"message": "db error",
		})
	} else {
		c.JSON(http.StatusOK, json)
	}

}
func NewHTTPServer() {
	r := gin.Default()
	r.GET("/db", getAllArticles)
	r.GET("/article", allTitle)
	r.POST("/article", addNewArticle)
	r.GET("/article/:article_id", getOneArticle)
	r.POST("/article/:article_id")
	r.GET("/", hello)

	r.Run()
}
