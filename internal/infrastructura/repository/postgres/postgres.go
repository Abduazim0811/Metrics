package postgres

import (
	"database/sql"
	"library/internal/entity/book"
	"library/internal/infrastructura/repository"
	"log"

	"github.com/Masterminds/squirrel"
)

type BookPostgres struct {
	db *sql.DB
}

func NewBookPostgres(db *sql.DB) repository.BookRepository {
	return &BookPostgres{db: db}
}

func (b *BookPostgres) AddBook(req book.CreateBook) error {
	sql, args, err := squirrel.
		Insert("book").
		Columns("title", "author", "published_date", "isbn").
		Values(req.Title, req.Author, req.Published_Date, req.Isbn).
		PlaceholderFormat(squirrel.Dollar).ToSql()

	if err != nil {
		log.Println("Unable to insert book:", err)
		return err
	}

	_, err = b.db.Exec(sql, args...)
	if err != nil {
		log.Println("Error executing insert:", err)
		return err
	}
	return nil
}

func (b *BookPostgres) GetByIdBook(id string) (*book.Book, error) {
	var book book.Book
	sql, args, err := squirrel.
		Select("*").
		From("book").
		Where(squirrel.Eq{"id": id}).
		PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		log.Println("Unable to generate SQL for GetByIdBook:", err)
		return nil, err
	}

	row := b.db.QueryRow(sql, args...)
	if err := row.Scan(&book.ID, &book.Title, &book.Author, &book.Published_Date, &book.Isbn); err != nil {
		log.Println("Error scanning row in GetByIdBook:", err)
		return nil, err
	}

	return &book, nil
}

func (b *BookPostgres) GetAllBooks() ([]*book.Book, error) {
	var Books []*book.Book

	sql, args, err := squirrel.
		Select("*").
		From("book").
		PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		log.Println("Unable to generate SQL for GetAllBooks:", err)
		return nil, err
	}

	rows, err := b.db.Query(sql, args...)
	if err != nil {
		log.Println("Error executing query in GetAllBooks:", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var book book.Book

		if err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Published_Date, &book.Isbn); err != nil {
			log.Println("Error scanning row in GetAllBooks:", err)
			return nil, err
		}

		Books = append(Books, &book)
	}

	if err := rows.Err(); err != nil {
		log.Println("Row iteration error in GetAllBooks:", err)
		return nil, err
	}

	return Books, nil
}

func (b *BookPostgres) UpdateBook(req book.Book) error {
	sql, args, err := squirrel.
		Update("book").
		Set("title", req.Title).
		Set("author", req.Author).
		Set("published_date", req.Published_Date).
		Set("isbn", req.Isbn).
		Where(squirrel.Eq{"id": req.ID}).
		PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		log.Println("unable to update book")
		return err
	}

	_, err = b.db.Exec(sql, args...)
	if err != nil {
		log.Println("Error executing update", err)
		return err
	}

	return nil
}

func (b *BookPostgres) Deletebook(id string) error {
	sql, args, err := squirrel.
		Delete("book").
		Where(squirrel.Eq{"id": id}).
		PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		log.Println("unable to delete book")
		return err
	}

	_, err = b.db.Exec(sql, args)
	if err != nil {
		log.Println("Error exucting delete: ", err)
		return err
	}

	return nil
}
