package server

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
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

type ArticleAllSteps struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Steps  []Step
}

type ArticleOneStep struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Step   Step
}

type Step struct {
	Content string `json:"content"`
	Codes   []Code
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
		return nil, errors.WithStack(err)
	}
	return db, errors.WithStack(err)
}

func SelectAllArticle() ([]Title, error) {
	db, err := dbOpen()
	defer db.Close()
	rows, err := db.Query("select article_id,title,author,likes from articles")

	if err != nil {
		return nil, errors.WithStack(err)
	}

	var result []Title

	for rows.Next() {
		var title Title
		err = rows.Scan(&title.ArticleId, &title.Title, &title.Author, &title.Author)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		result = append(result, title)

	}
	return result, err
}

func SelectOneArticle(articleId string) (ArticleAllSteps, error) {
	db, err := dbOpen()
	if err != nil {
		return ArticleAllSteps{}, errors.WithStack(err)
	}
	defer db.Close()
	var article ArticleAllSteps
	err = db.QueryRow("select title,author from articles where article_id=$1", articleId).Scan(&article.Title, &article.Author)
	if err != nil {
		return ArticleAllSteps{}, errors.WithStack(err)
	}
	rows, err := db.Query(`select step_id,code_id,code_file_name,code_content,article_content from 
		(select title,author,step_primary_key,step_id,articles.article_id,article_content from articles
		join steps on articles.article_id = steps.article_id where articles.article_id=$1) as article_steps 
		join codes on article_steps.step_primary_key = codes.step_primary_key`, articleId)

	if err != nil {
		return ArticleAllSteps{}, errors.WithStack(err)
	}
	var steps []Step

	for rows.Next() {
		var stepId int
		var codeId int
		var codeFileName string
		var codeContent string
		var articleContent string

		err = rows.Scan(&stepId, &codeId, &codeFileName, &codeContent, &articleContent)
		if err != nil {
			return ArticleAllSteps{}, errors.WithStack(err)
		}
		log.Println("データベースより取得", stepId, codeId, codeFileName, codeContent)
		if len(steps) > stepId {
			code := Code{CodeContent: codeContent, CodeFileName: codeFileName}
			steps[stepId].Codes = append(steps[stepId].Codes, code)
		} else {
			code := Code{CodeContent: codeContent, CodeFileName: codeFileName}
			step := Step{Codes: []Code{code}, Content: articleContent}
			steps = append(steps, step)
		}
	}
	article.Steps = steps

	return article, nil
}

func SelectOneArticleStep(articleId, stepId string) (ArticleOneStep, error) {
	db, err := dbOpen()
	if err != nil {
		return ArticleOneStep{}, errors.WithStack(err)
	}
	defer db.Close()
	var article ArticleOneStep
	err = db.QueryRow("select title,author from articles where article_id=$1", articleId).Scan(&article.Title, &article.Author)
	if err != nil {
		return ArticleOneStep{}, errors.WithStack(err)
	}
	rows, err := db.Query(`select code_id,code_file_name,code_content,article_content from 
		(select title,author,step_primary_key,step_id,articles.article_id,article_content from articles
		join steps on articles.article_id = steps.article_id where articles.article_id=$1 and steps.step_id=$2) as article_steps 
		join codes on article_steps.step_primary_key = codes.step_primary_key`, articleId, stepId)

	if err != nil {
		log.Println(err)
		return ArticleOneStep{}, errors.WithStack(err)
	}
	var codes []Code
	var articleContent string
	for rows.Next() {
		var codeId int
		var codeFileName string
		var codeContent string

		err = rows.Scan(&codeId, &codeFileName, &codeContent, &articleContent)
		if err != nil {
			return ArticleOneStep{}, errors.WithStack(err)
		}
		log.Println("データベースより取得", stepId, codeId, codeFileName, codeContent)
		code := Code{CodeContent: codeContent, CodeFileName: codeFileName}
		codes = append(codes, code)

	}
	article.Step = Step{Codes: codes, Content: articleContent}

	return article, nil
}

func AddNewArticle(postJson ArticleAllSteps) error {
	db, err := dbOpen()
	if err != nil {
		return errors.WithStack(err)
	}
	defer db.Close()
	tx, err := db.Begin()
	if err != nil {
		return errors.WithStack(err)
	}
	var article_id int
	err = tx.QueryRow("select max(article_id)+1 from articles").Scan(&article_id)
	if err != nil {
		tx.Rollback()
		return errors.WithStack(err)
	}
	_, err = tx.Exec("insert into articles values ($1,$2,$3,$4)", article_id, postJson.Title, postJson.Author, 0)
	if err != nil {
		tx.Rollback()
		return errors.WithStack(err)
	}

	var stepPrimaryKey int
	err = db.QueryRow("select max(step_primary_key)+1 from steps").Scan(&stepPrimaryKey)
	if err != nil {
		tx.Rollback()
		return errors.WithStack(err)
	}

	for i, v := range postJson.Steps {
		_, err = tx.Exec("insert into steps values ($1,$2,$3,$4)", stepPrimaryKey+i, article_id, i, v.Content)
		if err != nil {
			log.Println("aa")
			tx.Rollback()
			return errors.WithStack(err)
		}
		for i2, v2 := range v.Codes {
			_, err = tx.Exec("insert into codes values ($1,$2,$3,$4)", stepPrimaryKey+i, i2, v2.CodeFileName, v2.CodeContent)
			if err != nil {
				tx.Rollback()
				return errors.WithStack(err)
			}

		}

	}

	if err == nil {
		tx.Commit()
	}
	return err
}
