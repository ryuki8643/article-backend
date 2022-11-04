package server

import (
	"github.com/gin-gonic/gin"
	"log"
	"strconv"

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
// @Success 200 {array} Title
// @Router /articles [get]
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

// GetOneArticle ...
// @Summary urlパラメータで指定された番号の記事を出力
// @Tags article
// @Produce  json
// @Param article_id path int true "Article ID"
// @Success 200 {object} ArticleAllSteps
// @Router /articles/{article_id} [get]
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

// GetOneArticleStep ...
// @Summary urlパラメータで指定された記事のステップを出力
// @Tags article
// @Produce  json
// @Param article_id path int true "Article ID"
// @Param step_id path int true "Step ID"
// @Success 200 {object} ArticleOneStep
// @Router /articles/{article_id}/{step_id} [get]
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

func insertArticleTest(c *gin.Context) {
	p := ArticleAllSteps{Author: "a", Title: "b", Steps: []Step{{Content: "content", Codes: []Code{{CodeContent: "ss", CodeFileName: "cc"}}}}}
	err := AddNewArticle(p)
	if err != nil {
		log.Printf("%+v", err)
	}
	getAllArticles(c)
}

func EditArticleTest(c *gin.Context) {
	p := ArticleAllSteps{Author: "penguin", Title: "penguin", Steps: []Step{{Content: "content", Codes: []Code{{CodeContent: "ss", CodeFileName: "cc"}}}}}

	i, err := strconv.Atoi(c.Param("article_id"))
	if err != nil {
		log.Printf("%+v", err)
		c.JSONP(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	err = EditArticle(p, i)
	if err != nil {
		log.Printf("%+v", err)
		c.JSONP(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	getAllArticles(c)
}

// AddNewArticle ...
// @Summary 新しい記事の投稿
// @Tags article
// @Produce  json
// @Param article_json body ArticleAllSteps true "Article Json"
// @Success 200 {object} Message
// @Failure 400 {object} Message
// @Router /articles [post]
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

// EditArticle ...
// @Summary 記事の編集
// @Tags article
// @Produce  json
// @Param article_id path int true "Article ID"
// @Param article_json body ArticleAllSteps true "Article Json"
// @Success 200 {object} Message
// @Failure 400 {object} Message
// @Router /articles/{article_id} [put]
func editArticle(c *gin.Context) {
	var postJson ArticleAllSteps
	if err := c.ShouldBindJSON(&postJson); err != nil {
		log.Printf("%+v", err)
		c.JSONP(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	i, err := strconv.Atoi(c.Param("article_id"))
	if err != nil {
		log.Printf("%+v", err)
		c.JSONP(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	err = EditArticle(postJson, i)
	if err != nil {
		log.Printf("%+v", err)
		c.JSONP(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSONP(http.StatusOK, gin.H{
		"message": "put completed",
	})
}

// LikesArticle ...
// @Summary いいね数の追加
// @Tags like
// @Produce  json
// @Param article_id path int true "Article ID"
// @Success 200 {object} Message
// @Failure 400 {object} Message
// @Router /likes/{article_id} [put]
func LikesArticle(c *gin.Context) {
	err := AddLike(c.Param("article_id"))
	if err != nil {
		c.JSONP(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		log.Printf("%+v", err)
		return
	}

	c.JSONP(http.StatusOK, gin.H{
		"message": "put completed",
	})
}

// Swagger ...
// @Summary /swagger/index.html#/にアクセスするとswaggerを返す
// @Tags helloWorld
// @Produce  json
// @Failure 400 {object} Message
// @Router /swagger [get]
func ginSwaggerDoc() func(c *gin.Context) {
	return ginSwagger.WrapHandler(swaggerFiles.Handler)
}

// @title article_api
// @version 2.0
// @license.name ryuki
func NewHTTPServer() {
	r := gin.Default()
	r.Use(ZapLogger)
	r.GET("/articles", getAllArticles)
	r.GET("/articles/:article_id", getOneArticle)
	r.GET("/articles/:article_id/:step_id", getOneArticleStep)

	r.POST("/articles", insertNewArticle)
	r.PUT("/articles/:article_id", editArticle)

	r.GET("/likes/:article_id", LikesArticle)

	r.GET("/check/post", insertArticleTest)
	r.GET("/check/put/:article_id", EditArticleTest)

	r.GET("/", hello)
	r.GET("/swagger/*any", ginSwaggerDoc())

	r.Run()
}
