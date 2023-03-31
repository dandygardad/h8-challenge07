package controller

import (
	"challenge07/helper"
	"challenge07/model/entity"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type BookController interface {
	CreateBook(ctx *gin.Context)
	GetAllBooks(ctx *gin.Context)
	GetBook(ctx *gin.Context)
	UpdateBook(ctx *gin.Context)
	DeleteBook(ctx *gin.Context)
}

func (c Controller) CreateBook(ctx *gin.Context) {
	// Mengisi variabel newBook dari request json
	var newBook entity.Book
	errBind := ctx.ShouldBindJSON(&newBook)
	if errBind != nil {
		helper.ResponseError(ctx, "Invalid json", http.StatusBadRequest)
		return
	}

	// validasi name_book dan author
	if newBook.NameBook == "" || newBook.Author == "" {
		helper.ResponseError(ctx, "name_book/author required", http.StatusBadRequest)
		return
	}

	// Jika dapat request, masukkan dalam service
	result, err := c.service.CreateBook(newBook)
	if err != nil {
		helper.ResponseError(ctx, err.Error(), http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusCreated, result)
}

func (c Controller) GetAllBooks(ctx *gin.Context) {
	books, err := c.service.GetAllBooks()
	if err != nil {
		if err.Error() == "no books" {
			helper.ResponseError(ctx, "Tidak ada buku tersimpan", http.StatusOK)
		} else {
			helper.ResponseError(ctx, "Server error", http.StatusInternalServerError)
		}
		return
	}
	ctx.JSON(http.StatusOK, books)
}

func (c Controller) GetBook(ctx *gin.Context) {
	// Ambil id dari url
	id := ctx.Param("id")
	cvtId, errCvt := strconv.Atoi(id)
	if errCvt != nil {
		fmt.Println(errCvt)
		helper.ResponseError(ctx, "Input bukan tipe nomor", http.StatusBadRequest)
		return
	}

	book, err := c.service.GetBook(cvtId)
	if err != nil {
		if err == sql.ErrNoRows {
			helper.ResponseError(ctx, "Tidak ditemukan", http.StatusNotFound)
		} else {
			helper.ResponseError(ctx, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	ctx.JSON(http.StatusOK, book)
}

func (c Controller) UpdateBook(ctx *gin.Context) {
	// Ambil id dari url
	id := ctx.Param("id")
	cvtId, errCvt := strconv.Atoi(id)
	if errCvt != nil {
		helper.ResponseError(ctx, "Input bukan tipe nomor", 400)
		return
	}

	// Ambil data json dari request
	var newBook entity.Book
	errJson := ctx.ShouldBindJSON(&newBook)
	if errJson != nil {
		helper.ResponseError(ctx, "Invalid json", http.StatusBadRequest)
		return
	}
	// validasi name_book dan author
	if newBook.NameBook == "" || newBook.Author == "" {
		helper.ResponseError(ctx, "name_book/author required", http.StatusBadRequest)
		return
	}

	result, err := c.service.UpdateBook(cvtId, newBook)
	if err != nil {
		if err == sql.ErrNoRows {
			helper.ResponseError(ctx, "Id tidak ditemukan", http.StatusNotFound)
		} else {
			helper.ResponseError(ctx, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	ctx.JSON(http.StatusOK, result)
}

func (c Controller) DeleteBook(ctx *gin.Context) {
	// Ambil id dari url
	id := ctx.Param("id")
	cvtId, errCvt := strconv.Atoi(id)
	if errCvt != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Input bukan tipe nomor",
		})
		return
	}

	err := c.service.DeleteBook(cvtId)
	if err != nil {
		if err.Error() == "no data deleted" {
			helper.ResponseError(ctx, "Book tidak ditemukan", 404)
		} else {
			helper.ResponseError(ctx, err.Error(), 500)
		}
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Book deleted successfully",
	})
}
