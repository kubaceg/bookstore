package test

import (
	"context"

	"github.com/stretchr/testify/mock"

	"github.com/kubaceg/bookstore/internal/books/repository"
)

type BookRepositoryMock struct {
	mock.Mock
}

func (m *BookRepositoryMock) AddBook(ctx context.Context, book repository.BookEntity) (id *string, err error) {
	args := m.Called(ctx, book)

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*string), args.Error(1)
}

func (m *BookRepositoryMock) GetBook(ctx context.Context, id string) (book *repository.BookEntity, err error) {
	args := m.Called(ctx, id)

	return args.Get(0).(*repository.BookEntity), args.Error(1)
}

func (m *BookRepositoryMock) GetBookList(ctx context.Context, params repository.BookListParams) (list []repository.BookEntity, err error) {
	args := m.Called(ctx, params)

	return args.Get(0).([]repository.BookEntity), args.Error(1)
}

func (m *BookRepositoryMock) ReserveBook(ctx context.Context, id string) (err error) {
	args := m.Called(ctx, id)

	return args.Error(0)
}

func (m *BookRepositoryMock) ReleaseBook(ctx context.Context, id string) (err error) {
	args := m.Called(ctx, id)

	return args.Error(0)
}
