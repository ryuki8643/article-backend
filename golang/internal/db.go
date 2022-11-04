package server

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

type Message struct {
	Message string `json:"message"`
}

type Title struct {
	ArticleId string `json:"articleId"`
	Title     string `json:"title"`
	Author    string `json:"author"`
	Likes     int    `json:"likes"`
}

type Article struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Steps  []Step
}

type Step struct {
	Codes []Code
}

type Code struct {
	CodeFileName string `json:"code_file_name"`
	CodeContent  string `json:"code_content"`
}

const (
	host     = "postgres"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "pgweb"
)

func dbOpen() (*sql.DB, error) {
	psqlConn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlConn)

	if err != nil {

		return nil, err
	}
	return db, err
}

func SelectAllArticle() ([]Title, error) {
	db, err := dbOpen()
	defer db.Close()
	rows, err := db.Query("select article_id,title,author,likes from articles")

	if err != nil {
		return nil, err
	}

	var result []Title

	for rows.Next() {
		var title Title
		rows.Scan(&title.ArticleId, &title.Title, &title.Author, &title.Author)
		result = append(result, title)

	}
	return result, err
}

func SelectOneArticleStep(articleId, stepId string) (Article, error) {
	db, err := dbOpen()
	if err != nil {
		return Article{}, err
	}
	defer db.Close()
	var article Article
	err = db.QueryRow("select title,author from articles where article_id=$1", articleId).Scan(&article.Title, &article.Author)
	if err != nil {
		return Article{}, err
	}
	rows, err := db.Query(`select step_id,code_id,code_file_name,code_content from 
		(select title,author,step_primary_key,step_id,articles.article_id from articles
		join steps on articles.article_id = steps.article_id where articles.article_id=$1 and steps.step_id=$2) as article_steps 
		join codes on article_steps.step_primary_key = codes.step_primary_key`, articleId, stepId)

	if err != nil {
		return Article{}, err
	}
	var steps []Step

	for rows.Next() {
		var stepId int
		var codeId int
		var codeFileName string
		var codeContent string
		rows.Scan(&stepId, &codeId, &codeFileName, &codeContent)
		log.Println("データベースより取得", stepId, codeId, codeFileName, codeContent)
		if len(steps) > stepId {
			code := Code{CodeContent: codeContent, CodeFileName: codeFileName}
			steps[stepId].Codes = append(steps[stepId].Codes, code)
		} else {
			code := Code{CodeContent: codeContent, CodeFileName: codeFileName}
			step := Step{Codes: []Code{code}}
			steps = append(steps, step)
		}
	}
	article.Steps = steps

	return article, nil
}
