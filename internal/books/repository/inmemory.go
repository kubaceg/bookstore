package repository

import (
	"context"
	"fmt"

	"github.com/kubaceg/bookstore/internal/common/log"
)

type InMemoryBookRepository struct {
	books  map[string]BookEntity
	logger log.Logger
}

func NewInMemoryBookRepository(l log.Logger) BookRepository {
	return &InMemoryBookRepository{
		books:  map[string]BookEntity{},
		logger: l,
	}
}

func (i *InMemoryBookRepository) AddBook(ctx context.Context, book BookEntity) (id *string, err error) {
	if _, ok := i.books[book.Id]; ok {
		i.logger.Error(ctx, fmt.Sprintf("%s already exists in db", book.Title))

		return nil, BookAlreadyExists
	}

	i.books[book.Id] = book

	i.logger.Info(ctx, fmt.Sprintf("%s book added to db", book.Title))

	return &book.Id, nil
}

func (i *InMemoryBookRepository) GetBook(_ context.Context, id string) (book *BookEntity, err error) {
	if book, ok := i.books[id]; ok {
		return &book, nil
	}

	return nil, BookNotFound
}

func (i *InMemoryBookRepository) GetBookList(_ context.Context, params BookListParams) (list []BookEntity, err error) {
	list = []BookEntity{}

	for _, book := range i.books {
		list = append(list, book)
	}

	return
}