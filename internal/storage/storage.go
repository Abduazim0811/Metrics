package storage

import (
	"database/sql"
	"library/internal/http/api/handler"
	"library/internal/infrastructura/repository/postgres"
	"library/internal/service"
	"log"
)

func OpenSql() (*sql.DB, error) {
	db, err := sql.Open("postgres://postgres:Abdu0811@localhost:5432/n9?sslmode=disable", "postgres")
	if err != nil {
		log.Println("failed to open sql:", err)
		return nil, err
	}

	if err = db.Ping(); err != nil {
		log.Println("unable to connect database:", err)
		return nil, err
	}
	return db, err
}

func Handler() *handler.BookHandler {
	db, err := OpenSql()
	if err != nil {
		log.Println(err)
		return nil
	}

	repo := postgres.NewBookPostgres(db)
	service := service.NewBookService(repo)
	handler := handler.NewBookHandler(*service)

	return handler
}
