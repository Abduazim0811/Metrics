package service

import (
	"library/internal/entity/book"
	"library/internal/infrastructura/repository"
)

type BookService struct{
	repo repository.BookRepository
}

func NewBookService(repo repository.BookRepository) *BookService{
	return &BookService{repo: repo}
}

func (b *BookService) Createbook(req book.CreateBook) error{
	return b.repo.AddBook(req)
}

func (b *BookService) GetbyidBook(id string) (*book.Book, error){
	return b.repo.GetByIdBook(id)
}

func (b *BookService) GetAllbooks()([]*book.Book, error){
	return b.repo.GetAllBooks()
}

func (b *BookService) Updatebook(req *book.Book) error{
	return b.repo.UpdateBook(req)
}

func (b *BookService) Deletebook(id string) error{
	return b.repo.Deletebook(id)
}

