package controller

import (
	"chap2-project/model"
	"chap2-project/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BookController struct {
	bookRepository repository.BookRepository
}

func NewBookController(bookRepository repository.BookRepository) *BookController {
	return &BookController{
		bookRepository: bookRepository,
	}
}
func (bc *BookController) GetBooks(ctx *gin.Context) {

	books, err := bc.bookRepository.Get()

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, books)
	return
}

func (bc *BookController) GetBookById(ctx *gin.Context) {
	idString := ctx.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	uintId := uint(id)
	book, err := bc.bookRepository.GetOne(uintId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, book)
	return
}
func (bc *BookController) AddBook(ctx *gin.Context) {
	newBook := model.Book{}

	if err := ctx.ShouldBindJSON(&newBook); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	savedBook, err := bc.bookRepository.Save(newBook)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, savedBook)
	return

}
func (bc *BookController) UpdateBook(ctx *gin.Context) {
	updatedBook := model.Book{}

	if err := ctx.ShouldBindJSON(&updatedBook); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	idString := ctx.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	uintId := uint(id)
	_, err = bc.bookRepository.GetOne(uintId)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	succesUpdate, err := bc.bookRepository.Update(updatedBook, uintId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, succesUpdate)
	return

}
func (bc *BookController) DeleteBook(ctx *gin.Context) {
	deletedBook := model.Book{}

	idString := ctx.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	uintId := uint(id)
	_, err = bc.bookRepository.GetOne(uintId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := bc.bookRepository.Delete(deletedBook, uintId); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Book deleted successfully",
	})
	return

}
