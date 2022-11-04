package server

import (
	"github.com/gin-gonic/gin"
	"log"

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
// @Router /article [get]
func getAllArticles(c *gin.Context) {
	articles, err := SelectAllArticle()
	if err != nil {
		log.Printf("%+v", err)
		c.JSONP(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, articles)
	}
}

// GetOneArticles ...
// @Summary urlパラメータで指定された番号の記事を出力
// @Tags article
// @Produce  json
// @Param article_id path int true "Article ID"
// @Success 200 {object} Article
// @Router /article/{article_id}/{step_id} [get]
func getOneArticleStep(c *gin.Context) {

	article, err := SelectOneArticleStep(c.Param("article_id"), c.Param("step_id"))
	if err != nil {
		log.Printf("%+v", err)
		c.JSONP(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, article)
	}

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
		log.Printf("%+v", err)
		c.JSONP(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, article)
	}

}

func insertArticleTest(c *gin.Context) {
	p := ArticleAllSteps{Author: "a", Title: "b", Steps: []Step{{Content: "content", Codes: []Code{{CodeContent: "ss", CodeFileName: "cc"}}}}}
	err := AddNewArticle(p)
	if err != nil {
		log.Printf("%+v", err)
	}
	getAllArticles(c)
}

func insertNewArticle(c *gin.Context) {
	var postJson ArticleAllSteps
	if err := c.ShouldBindJSON(&postJson); err != nil {
		log.Printf("%+v", err)
		c.JSONP(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	err := AddNewArticle(postJson)
	if err != nil {
		c.JSONP(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		log.Printf("%+v", err)
		return
	}

	c.JSONP(http.StatusOK, gin.H{
		"message": "post completed",
	})
}

// @title article_api
// @version 2.0
// @license.name ryuki
func NewHTTPServer() {
	r := gin.Default()
	r.GET("/articles", getAllArticles)
	r.GET("/articles/:article_id", getOneArticle)
	r.GET("/articles/:article_id/:step_id", getOneArticleStep)

	r.POST("articles", insertNewArticle)

	r.GET("/insertCheck", insertArticleTest)

	r.GET("/", hello)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run()
}
