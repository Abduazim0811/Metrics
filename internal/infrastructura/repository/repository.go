package repository

import "library/internal/entity/book"


type BookRepository interface{
	AddBook(req book.CreateBook)error
	GetByIdBook(id string) (*book.Book, error)
	GetAllBooks() ([]*book.Book, error)
	UpdateBook(req book.Book) error
	Deletebook(id string) error
}