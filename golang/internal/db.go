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
	sqlState := fmt.Sprintf("select * from articles where article_id = %s", articleId)
	err = db.QueryRow(sqlState).Scan(&article.ArticleID, &article.Title, &article.Content, &article.Author)
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
