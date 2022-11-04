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

	article, err := SelectOneArticleStep(c.Param("article_id"), c.Param("step_id"))
	if err != nil {
		c.JSONP(http.StatusBadRequest, gin.H{
			"message": "no_content",
		})
	} else {
		c.JSON(http.StatusOK, article)
	}

}

// @title article_api
// @version 2.0
// @license.name ryuki
func NewHTTPServer() {
	r := gin.Default()
	r.GET("/article", getAllArticles)

	r.GET("/article/:article_id/:step_id", getOneArticle)

	r.GET("/", hello)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run()
}
