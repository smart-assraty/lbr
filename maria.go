package main

import (
	"database/sql"

	"github.com/go-sql-driver/mysql"
)

type connector struct {
	user     string
	password string
	ip       string
	database string
}

func (connector connector) connect(query string) *sql.Rows {
	mysql := mysql.NewConfig()
	mysql.User = connector.user
	mysql.Passwd = connector.password
	mysql.Net = connector.ip
	mysql.DBName = connector.database

	db, errDB := sql.Open("mysql", mysql.FormatDSN())
	checkError("errDB", errDB)

	defer db.Close()

	rows, errQuery := db.Query(query)
	checkError("errQuery", errQuery)
	return rows
}

type Author struct {
	Id          int
	Name        string
	Description string
	Img         string
	ArticlesUrl string
}

func (connector connector) getAuthors() []Author {
	rows := connector.connect("select * from author;")

	var responses []Author
	for rows.Next() {
		var response Author
		checkError("getAuthors", rows.Scan(&response.Id, &response.Name, &response.Description, &response.Img, &response.ArticlesUrl))
		responses = append(responses, response)
	}
	return responses
}

type Article struct {
	Id       int
	Title    string
	AuthorId int
	Url      string
}

func (connector connector) getArticles() []Article {
	rows := connector.connect("select * from article;")
	var responses []Article
	for rows.Next() {
		var response Article
		checkError("getArticles", rows.Scan(&response.Id, &response.Title, &response.AuthorId, &response.Url))
		responses = append(responses, response)
	}
	return responses
}

func (connector connector) getArticlesOfAuthor(authorName string) []Article {
	rows := connector.connect("select * from article where ( select id from author where name='" + authorName + "') = author_id;")

	var responses []Article
	for rows.Next() {
		var response Article
		checkError("getArticlesOfAuthor", rows.Scan(&response.Id, &response.Title, &response.AuthorId, &response.Url))
		responses = append(responses, response)
	}
	return responses
}

type AuthorPage struct {
	Author   Author
	Articles []Article
}

func (connector connector) getAuthorPage(author Author) AuthorPage {
	var authorPage AuthorPage
	authorPage.Author = author
	authorPage.Articles = connector.getArticlesOfAuthor(authorPage.Author.Name)
	return authorPage
}
