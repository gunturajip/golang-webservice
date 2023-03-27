package controllers

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "db-go-sql"
)

var (
	db  *sql.DB
	err error
)

func StartDB() {
	config := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err = sql.Open("postgres", config)
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected to database")
}

type Book struct {
	BookID string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Descr  string `json:"desc"`
}

func GetBooks(ctx *gin.Context) {
	var books = []Book{}

	sqlStatement := `SELECT * FROM books`
	defer func() {
		if err := recover(); err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
			return
		}
	}()
	rows, _ := db.Query(sqlStatement)

	for rows.Next() {
		var book = Book{}
		err := rows.Scan(&book.BookID, &book.Title, &book.Author, &book.Descr)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
			return
		}
		books = append(books, book)
	}
	ctx.JSON(http.StatusOK, books)
}

func GetBook(ctx *gin.Context) {
	BookID := ctx.Param("bookID")

	var book = Book{}

	sqlStatement := `SELECT * FROM books WHERE id = $1`
	defer func() {
		if err := recover(); err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
			return
		}
	}()
	err := db.QueryRow(sqlStatement, BookID).
		Scan(&book.BookID, &book.Title, &book.Author, &book.Descr)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, book)
}

func CreateBook(ctx *gin.Context) {
	asdf := `SELECT * FROM books`
	_, asdfsfd := db.Query(asdf)
	if asdfsfd != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, "hola")
	}

	var book = Book{}

	type BookInput struct {
		Title  string `json:"title" binding:"required"`
		Author string `json:"author" binding:"required"`
		Descr  string `json:"desc" binding:"required"`
	}
	var input BookInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	sqlStatement := `
		INSERT INTO books (title, author, descr)
		VALUES ($1, $2, $3)
		Returning *
	`
	defer func() {
		if err := recover(); err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
			return
		}
	}()
	err := db.QueryRow(sqlStatement, input.Title, input.Author, input.Descr).
		Scan(&book.BookID, &book.Title, &book.Author, &book.Descr)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, "Created")
}

func UpdateBook(ctx *gin.Context) {
	BookID := ctx.Param("bookID")

	type BookInput struct {
		Title  string `json:"title" binding:"required"`
		Author string `json:"author" binding:"required"`
		Descr  string `json:"desc" binding:"required"`
	}
	var input BookInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	sqlStatement := `
		UPDATE books SET title = $2, author = $3, descr = $4
		WHERE id = $1;
	`
	defer func() {
		if err := recover(); err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
			return
		}
	}()
	_, err := db.Exec(sqlStatement, BookID, input.Title, input.Author, input.Descr)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, "Updated")
}

func DeleteBook(ctx *gin.Context) {
	BookID := ctx.Param("bookID")

	sqlStatement := `
		DELETE FROM books
		WHERE id = $1;
	`
	defer func() {
		if err := recover(); err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
			return
		}
	}()
	_, err := db.Exec(sqlStatement, BookID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, "Deleted")
}
