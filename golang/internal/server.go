package server

import (
	"fmt"

	"github.com/gin-gonic/gin"

	_ "github.com/ryuki8643/article-backend/internal/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

// Hello ...
// @Summary helloを返す
// @Tags helloWorld
// @Produce  json
// @Success 200 {object} Message
// @Router / [get]
func hello(c *gin.Context) {
	c.JSONP(http.StatusOK, gin.H{
		"message": "hello",
	})
}

// GetAllArticles ...
// @Summary 全ての記事のデータを返す
// @Tags db
// @Produce  json
// @Success 200 {array} Article
// @Router /db [get]
func getAllArticles(c *gin.Context) {
	articles, err := SelectAllArticle()
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(http.StatusOK, articles)
}

// GetOneArticles ...
// @Summary urlパラメータで指定された番号の記事を出力
// @Tags article
// @Produce  json
// @Param article_id path int true "Article ID"
// @Success 200 {object} Article
// @Router /article/{article_id} [get]
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

// allTitle ...
// @Summary 全ての記事のidとtitleを返す
// @Tags article
// @Produce  json
// @Success 200 {array} Titles
// @Router /article [get]
func allTitle(c *gin.Context) {
	titles, err := SelectAllTitleAndID()
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(http.StatusOK, titles)
}

// AddNewArticle ...
// @Summary 新しい記事の投稿
// @Tags article
// @Produce  json
// @Param article_json body ReceiveJson true "Article Json"
// @Success 200 {object} Article
// @Failure 400 {object} Message
// @Router /article [post]
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

// EditArticle...
// @Summary 既存の記事を編集
// @Tags article
// @Produce  json
// @Param article_id path int true "Article ID"
// @Param article_json body ReceiveJson true "Article Json"
// @Success 200 {object} Article
// @Failure 400 {object} Message
// @Router /article/{article_id} [post]
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

// @title article_api
// @version 2.0
// @license.name ryuki
func NewHTTPServer() {
	r := gin.Default()
	r.GET("/db", getAllArticles)
	r.GET("/article", allTitle)
	r.POST("/article", addNewArticle)
	r.GET("/article/:article_id", getOneArticle)
	r.POST("/article/:article_id")
	r.GET("/", hello)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run()
}
