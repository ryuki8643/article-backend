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

func NewHTTPServer() {
	r := gin.Default()
	r.GET("/db", getAllArticles)
	r.GET("/article", allTitle)
	r.GET("/article/:article_id", getOneArticle)
	r.GET("/", hello)

	r.Run()
}
