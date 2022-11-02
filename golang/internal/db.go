package server

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type ARTICLE struct {
	ArticleID int    `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Author    string `json:"author"`
}

const (
	host     = "postgres"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "pgweb"
)

func SelectAllDb() ([]ARTICLE, error) {
	psqlConn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlConn)
	defer db.Close()

	if err != nil {

		return nil, err
	}

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
