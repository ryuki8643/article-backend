package server

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type Titles struct {
	ArticleID int    `json:"id"`
	Title     string `json:"title"`
}

type ReceiveJson struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

type ARTICLE struct {
	Titles
	Content string `json:"content"`
	Author  string `json:"author"`
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

func SelectAllArticle() ([]ARTICLE, error) {
	db, err := dbOpen()
	defer db.Close()
	rows, err := db.Query("select * from articles")

	if err != nil {
		return nil, err
	}

	var result []ARTICLE

	for rows.Next() {
		var article ARTICLE
		rows.Scan(&article.ArticleID, &article.Title, &article.Content, &article.Author)
		result = append(result, article)

	}
	return result, err
}

func SelectOneArticle(articleId string) (ARTICLE, error) {
	db, err := dbOpen()
	if err != nil {
		return ARTICLE{}, err
	}
	defer db.Close()
	var article ARTICLE

	err = db.QueryRow("select * from articles where article_id = $1", articleId).Scan(&article.ArticleID, &article.Title, &article.Content, &article.Author)
	if err != nil {

		return ARTICLE{}, err
	}
	return article, nil
}

func SelectAllTitleAndID() ([]Titles, error) {
	db, err := dbOpen()
	defer db.Close()
	rows, err := db.Query("select article_id,title from articles")

	if err != nil {
		return nil, err
	}

	var result []Titles

	for rows.Next() {
		var article Titles
		rows.Scan(&article.ArticleID, &article.Title)
		result = append(result, article)

	}
	return result, err
}

func InsertNewArticle(postJson ReceiveJson) error {
	db, err := dbOpen()
	defer db.Close()
	if err != nil {
		return err
	}
	title := postJson.Title
	if title == "" {
		title = "no_title"
	}
	content := postJson.Content
	if content == "" {
		content = "no_content"
	}
	author := postJson.Author
	if author == "" {
		author = "no_author"
	}

	_, err = db.Exec("insert into articles  (title,article_content,author) values ($1,$2,$3)", title, content, author)
	return err
}

func EditArticle(articleId string, postJson ReceiveJson) error {
	db, err := dbOpen()
	defer db.Close()
	if err != nil {
		return err
	}

	title := postJson.Title
	if title == "" {
		title = "no_title"
	}
	content := postJson.Content
	if content == "" {
		content = "no_content"
	}
	author := postJson.Author
	if author == "" {
		author = "no_author"
	}
	_, err = db.Exec("update articles Set title=$1,article_content=$2,author=$3 WHERE article_id=$4", title, content, author, articleId)
	return err
}
