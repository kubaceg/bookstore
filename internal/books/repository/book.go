package repository

import (
	"context"
	"errors"
)

var BookNotFound = errors.New("book not found")
var BookAlreadyExists = errors.New("book already exists")

type BookEntity struct {
	Id     string `json:"id,omitempty"`
	Title  string `json:"title,omitempty"`
	Author string `json:"author,omitempty"`
	Isbn   string `json:"isbn,omitempty"`
}

type BookListParams struct {
	perPage int32
	page    int32
}

type BookRepository interface {
	AddBook(ctx context.Context, book BookEntity) (id *string, err error)
	GetBook(ctx context.Context, id string) (book *BookEntity, err error)
	GetBookList(ctx context.Context, params BookListParams) (list []BookEntity, err error)
}
