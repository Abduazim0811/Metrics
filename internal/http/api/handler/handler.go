package handler

import (
	"fmt"
	"library/internal/entity/book"
	"library/internal/service"
	_ "library/docs"
	"net/http"

	"github.com/gin-gonic/gin"
)

// BookHandler handles book-related requests
type BookHandler struct {
	service service.BookService
}

// NewBookHandler creates a new BookHandler
func NewBookHandler(service service.BookService) *BookHandler {
	return &BookHandler{service: service}
}

// CreateBook godoc
// @Summary Create a new book
// @Description Create a new book in the library
// @Tags books
// @Accept  json
// @Produce  json
// @Param book body book.CreateBook true "Book to create"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /books [post]
func (b *BookHandler) CreateBook(c *gin.Context) {
	var book book.CreateBook

	if err := c.ShouldBindJSON(&book); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	err := b.service.Createbook(book)

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Message": "Book created successfully"})
}

// GetByIDBook godoc
// @Summary Get a book by ID
// @Description Get details of a book by its ID
// @Tags books
// @Produce  json
// @Param id path string true "Book ID"
// @Success 200 {object} book.Book
// @Failure 500 {object} string
// @Router /books/{id} [get]
func (b *BookHandler) GetByIDBook(c *gin.Context) {
	id := c.Param("id")

	book, err := b.service.GetbyidBook(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, book)
}

// GetAllBooks godoc
// @Summary Get all books
// @Description Get a list of all books
// @Tags books
// @Produce  json
// @Success 200 {array} book.Book
// @Failure 500 {object} string
// @Router /books [get]
func (b *BookHandler) GetAllBooks(c *gin.Context) {
	books, err := b.service.GetAllbooks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, books)
}

// UpdateBook godoc
// @Summary Update a book
// @Description Update details of an existing book
// @Tags books
// @Accept  json
// @Produce  json
// @Param id path string true "Book ID"
// @Param book body book.Book true "Updated book"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /books/{id} [put]
func (b *BookHandler) UpdateBook(c *gin.Context) {
	id := c.Param("id")

	var req book.Book
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	req.ID = id

	err := b.service.Updatebook(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book updated successfully"})
}

// DeleteBook godoc
// @Summary Delete a book
// @Description Delete a book by its ID
// @Tags books
// @Produce  json
// @Param id path string true "Book ID"
// @Success 200 {object} string
// @Failure 500 {object} string
// @Router /books/{id} [delete]
func (b *BookHandler) DeleteBook(c *gin.Context) {
	id := c.Param("id")
	err := b.service.Deletebook(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book deleted"})
}

func (b *BookHandler) Check(c *gin.Context) {
	fmt.Fprintf(c.Writer, "everything is ok")
}
