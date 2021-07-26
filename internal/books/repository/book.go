package repository

import (
	"context"
	"errors"
)

type BookState int32

const (
	Available    BookState = 0
	Reserved               = 1
	NotAvailable           = 2
)

var BookNotFound = errors.New("book not found")
var BookAlreadyExists = errors.New("book already exists")
var BookAlreadyReserved = errors.New("book already reserved")

type BookEntity struct {
	Id     string    `json:"id,omitempty"`
	Title  string    `json:"title,omitempty"`
	Author string    `json:"author,omitempty"`
	Isbn   string    `json:"isbn,omitempty"`
	State  BookState `json:"state,omitempty"`
}

type BookListParams struct {
	perPage int32
	page    int32
}

type BookRepository interface {
	AddBook(ctx context.Context, book BookEntity) (id *string, err error)
	GetBook(ctx context.Context, id string) (book *BookEntity, err error)
	GetBookList(ctx context.Context, params BookListParams) (list []BookEntity, err error)
	ReserveBook(ctx context.Context, id string) (err error)
	ReleaseBook(ctx context.Context, id string) (err error)
}
