package server

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"strconv"
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
	Likes  string `json:"likes"`
	Steps  []Step
}

type ArticleOneStep struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Likes  string `json:"likes"`
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
	rows, err := db.Query("select article_id,title,author,likes from articles order by article_id")

	if err != nil {
		return nil, errors.WithStack(err)
	}

	var result []Title

	for rows.Next() {
		var title Title
		err = rows.Scan(&title.ArticleId, &title.Title, &title.Author, &title.Likes)
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
	err = db.QueryRow("select title,author,likes from articles where article_id=$1", articleId).Scan(&article.Title, &article.Author, &article.Likes)
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
	err = db.QueryRow("select title,author,likes from articles where article_id=$1", articleId).Scan(&article.Title, &article.Author, &article.Likes)
	if err != nil {
		return ArticleOneStep{}, errors.WithStack(err)
	}
	rows, err := db.Query(`select code_id,code_file_name,code_content,article_content from 
		(select title,author,step_primary_key,step_id,articles.article_id,article_content from articles
		join steps on articles.article_id = steps.article_id where articles.article_id=$1 and steps.step_id=$2) as article_steps 
		join codes on article_steps.step_primary_key = codes.step_primary_key`, articleId, stepId)

	if err != nil {
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

		code := Code{CodeContent: codeContent, CodeFileName: codeFileName}
		codes = append(codes, code)

	}
	article.Step = Step{Codes: codes, Content: articleContent}

	return article, nil
}

func addArticle(postJson ArticleAllSteps, articleId int, tx *sql.Tx) error {

	_, err := tx.Exec("insert into articles values ($1,$2,$3,$4)", articleId, postJson.Title, postJson.Author, 0)
	if err != nil {
		return errors.WithStack(err)
	}

	var stepPrimaryKey int
	err = tx.QueryRow("select max(step_primary_key)+1 from steps").Scan(&stepPrimaryKey)
	if err != nil {
		return errors.WithStack(err)
	}

	for i, v := range postJson.Steps {
		_, err = tx.Exec("insert into steps values ($1,$2,$3,$4)", stepPrimaryKey+i, articleId, i, v.Content)
		if err != nil {
			return errors.WithStack(err)
		}
		for i2, v2 := range v.Codes {
			_, err = tx.Exec("insert into codes values ($1,$2,$3,$4)", stepPrimaryKey+i, i2, v2.CodeFileName, v2.CodeContent)
			if err != nil {
				return errors.WithStack(err)
			}

		}

	}
	return err
}

func AddNewArticle(postJson ArticleAllSteps) error {
	db, err := dbOpen()
	if err != nil {
		return errors.WithStack(err)
	}
	defer db.Close()

	var articleId int
	err = db.QueryRow("select max(article_id)+1 from articles").Scan(&articleId)
	if err != nil {
		return errors.WithStack(err)
	}

	tx, err := db.Begin()
	if err != nil {
		return errors.WithStack(err)
	}
	err = addArticle(postJson, articleId, tx)

	if err != nil {
		return err
	}

	if err == nil {
		tx.Commit()
	} else {
		tx.Rollback()
	}
	return err
}

func EditArticle(postJson ArticleAllSteps, articleId int) error {
	db, err := dbOpen()
	if err != nil {
		return errors.WithStack(err)
	}
	defer db.Close()

	err = deleteArticle(db, articleId)
	if err != nil {
		return err
	}

	tx, err := db.Begin()
	if err != nil {
		return errors.WithStack(err)
	}
	addArticle(postJson, articleId, tx)

	if err == nil {
		tx.Commit()
	} else {
		tx.Rollback()
	}

	return err
}

func deleteArticle(db *sql.DB, articleId int) error {
	rows, err := db.Query("select step_primary_key from steps where article_id=$1", articleId)
	if err != nil {
		return errors.WithStack(err)
	}

	var stepPrimaryKeys []string

	for rows.Next() {
		var stepPrimaryKey string
		rows.Scan(&stepPrimaryKey)
		stepPrimaryKeys = append(stepPrimaryKeys, stepPrimaryKey)
	}
	tx, err := db.Begin()
	if err != nil {
		return errors.WithStack(err)
	}

	for _, v := range stepPrimaryKeys {
		_, err = tx.Exec("delete from codes where step_primary_key=$1", v)
		if err != nil {
			return errors.WithStack(err)
			tx.Rollback()
		}
	}
	_, err = tx.Exec("delete from steps where article_id=$1", articleId)
	if err != nil {
		return errors.WithStack(err)
		tx.Rollback()
	}
	_, err = tx.Exec("delete from articles where article_id=$1", articleId)
	if err != nil {
		return errors.WithStack(err)
		tx.Rollback()
	}

	if err == nil {
		tx.Commit()
	}
	return nil
}

func AddLike(articleId string) error {
	a, err := strconv.Atoi(articleId)
	if err != nil {
		return errors.WithStack(err)
	}
	db, err := dbOpen()
	if err != nil {
		return errors.WithStack(err)
	}
	defer db.Close()

	_, err = db.Exec("update articles set likes=(select likes from articles where article_id=$1)+1 where article_id=$1", a)
	return errors.WithStack(err)
}
