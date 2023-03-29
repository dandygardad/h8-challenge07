package controller

import (
	"challenge07/helper"
	"challenge07/model/entity"
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
		helper.BadRequestError(ctx, "Invalid json")
		return
	}

	// validasi title dan author
	if newBook.Title == "" || newBook.Author == "" {
		helper.BadRequestError(ctx, "title/author required")
		return
	}

	// Jika dapat request, masukkan dalam service
	err := c.service.CreateBook(newBook)
	if err != nil {
		helper.BadRequestError(ctx, err.Error())
		return
	}

	ctx.String(http.StatusCreated, "Created")
}

func (c Controller) GetAllBooks(ctx *gin.Context) {
	books, err := c.service.GetAllBooks()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, books)
}

func (c Controller) GetBook(ctx *gin.Context) {
	// Ambil id dari url
	id := ctx.Param("id")
	cvtId, errCvt := strconv.Atoi(id)
	if errCvt != nil {
		helper.BadRequestError(ctx, "Input bukan tipe nomor")
		return
	}

	book, err := c.service.GetBook(cvtId)
	if err != nil {
		helper.BadRequestError(ctx, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, book)
}

func (c Controller) UpdateBook(ctx *gin.Context) {
	// Ambil id dari url
	id := ctx.Param("id")
	cvtId, errCvt := strconv.Atoi(id)
	if errCvt != nil {
		helper.BadRequestError(ctx, "Input bukan tipe nomor")
		return
	}

	// Ambil data json dari request
	var newBook entity.Book
	errJson := ctx.ShouldBindJSON(&newBook)
	if errJson != nil {
		helper.BadRequestError(ctx, "Invalid json")
		return
	}
	// validasi title dan author
	if newBook.Title == "" || newBook.Author == "" {
		helper.BadRequestError(ctx, "title/author required")
		return
	}

	err := c.service.UpdateBook(cvtId, newBook)
	if err != nil {
		helper.BadRequestError(ctx, err.Error())
		return
	}

	ctx.String(http.StatusOK, "Updated")
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
		helper.BadRequestError(ctx, err.Error())
	}

	ctx.String(http.StatusOK, "Deleted")
}
