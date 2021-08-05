package service

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/stretchr/testify/mock"

	"github.com/kubaceg/bookstore/internal/books/repository"
	"github.com/kubaceg/bookstore/internal/books/repository/test"
	"github.com/kubaceg/bookstore/internal/common/genproto/book"
)

func TestBookGrpcService_AddBook(t *testing.T) {
	bookId := "1234"

	tests := []struct {
		name               string
		givenBook          book.Book
		repoMock           func(entity repository.BookEntity) *test.BookRepositoryMock
		expectedBookEntity repository.BookEntity
		want               *book.BookId
		wantErr            bool
	}{
		{
			name: "book already exist error",
			givenBook: book.Book{
				Id:     "1234",
				Title:  "Test Book",
				Author: "Jakub",
				Isbn:   "1234",
				State:  book.State_AVAILABLE,
			},
			repoMock: func(entity repository.BookEntity) *test.BookRepositoryMock {
				repo := test.BookRepositoryMock{}
				repo.On("AddBook", mock.Anything, entity).Return(nil, repository.BookAlreadyExists)

				return &repo
			},
			expectedBookEntity: repository.BookEntity{
				Id:     "1234",
				Title:  "Test Book",
				Author: "Jakub",
				Isbn:   "1234",
				State:  repository.Available,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "creates book",
			givenBook: book.Book{
				Id:     "1234",
				Title:  "Test Book",
				Author: "Jakub",
				Isbn:   "1234",
				State:  book.State_NOT_AVAILABLE,
			},
			repoMock: func(entity repository.BookEntity) *test.BookRepositoryMock {
				repo := test.BookRepositoryMock{}
				repo.On("AddBook", mock.Anything, entity).Return(&bookId, nil)

				return &repo
			},
			expectedBookEntity: repository.BookEntity{
				Id:     "1234",
				Title:  "Test Book",
				Author: "Jakub",
				Isbn:   "1234",
				State:  repository.NotAvailable,
			},
			want:    &book.BookId{Id: "1234"},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repoMock := tt.repoMock(tt.expectedBookEntity)
			b := NewBookGrpcService(repoMock)

			got, err := b.AddBook(context.TODO(), &tt.givenBook)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddBook() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddBook() got = %v, want %v", got, tt.want)
			}

			repoMock.AssertExpectations(t)
		})
	}
}

func TestBookGrpcService_GetBook(t *testing.T) {
	bookId := "1234"

	tests := []struct {
		name               string
		ctx                context.Context
		bookId             book.BookId
		repoMock           func(ctx context.Context, entity *repository.BookEntity) *test.BookRepositoryMock
		expectedBookEntity *repository.BookEntity
		want               *book.Book
		wantErr            bool
	}{
		{
			name: "book not found",
			ctx:  context.TODO(),
			bookId: book.BookId{
				Id: bookId,
			},
			repoMock: func(ctx context.Context, entity *repository.BookEntity) *test.BookRepositoryMock {
				repo := test.BookRepositoryMock{}
				repo.On("GetBook", ctx, bookId).Return(entity, repository.BookNotFound)

				return &repo
			},
			expectedBookEntity: nil,
			want:               nil,
			wantErr:            true,
		},
		{
			name: "found book with available status",
			ctx:  context.TODO(),
			bookId: book.BookId{
				Id: bookId,
			},
			repoMock: func(ctx context.Context, entity *repository.BookEntity) *test.BookRepositoryMock {
				repo := test.BookRepositoryMock{}
				repo.On("GetBook", ctx, bookId).Return(entity, nil)

				return &repo
			},
			expectedBookEntity: &repository.BookEntity{
				Id:     "1234",
				Title:  "Test Book",
				Author: "Jakub",
				Isbn:   "1234",
				State:  repository.Available,
			},
			want: &book.Book{
				Id:     "1234",
				Title:  "Test Book",
				Author: "Jakub",
				Isbn:   "1234",
				State:  book.State_AVAILABLE,
			},
			wantErr: false,
		},
		{
			name: "found book with reserved status",
			ctx:  context.TODO(),
			bookId: book.BookId{
				Id: bookId,
			},
			repoMock: func(ctx context.Context, entity *repository.BookEntity) *test.BookRepositoryMock {
				repo := test.BookRepositoryMock{}
				repo.On("GetBook", ctx, bookId).Return(entity, nil)

				return &repo
			},
			expectedBookEntity: &repository.BookEntity{
				Id:     "1234",
				Title:  "Test Book",
				Author: "Jakub",
				Isbn:   "1234",
				State:  repository.Reserved,
			},
			want: &book.Book{
				Id:     "1234",
				Title:  "Test Book",
				Author: "Jakub",
				Isbn:   "1234",
				State:  book.State_RESERVED,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repoMock := tt.repoMock(tt.ctx, tt.expectedBookEntity)
			b := NewBookGrpcService(repoMock)

			got, err := b.GetBook(tt.ctx, &tt.bookId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBook() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetBook() got = %v, want %v", got, tt.want)
			}

			repoMock.AssertExpectations(t)
		})
	}
}

func TestBookGrpcService_ReserveBook(t *testing.T) {
	bookId := "1234"

	tests := []struct {
		name     string
		ctx      context.Context
		bookId   book.BookId
		repoMock func(ctx context.Context) *test.BookRepositoryMock
		want     *book.ReservationStatus
		wantErr  error
	}{
		{
			name: "book not found",
			ctx:  context.TODO(),
			bookId: book.BookId{
				Id: bookId,
			},
			repoMock: func(ctx context.Context) *test.BookRepositoryMock {
				repo := test.BookRepositoryMock{}
				repo.On("GetBook", ctx, bookId).Return(nil, repository.BookNotFound)

				return &repo
			},
			want:    nil,
			wantErr: repository.BookNotFound,
		},
		{
			name: "book already reserved error",
			ctx:  context.TODO(),
			bookId: book.BookId{
				Id: bookId,
			},
			repoMock: func(ctx context.Context) *test.BookRepositoryMock {
				repo := test.BookRepositoryMock{}
				repo.On("GetBook", ctx, bookId).Return(&repository.BookEntity{State: repository.Reserved}, nil)

				return &repo
			},
			want:    nil,
			wantErr: repository.BookAlreadyReserved,
		},
		{
			name: "book reservation success",
			ctx:  context.TODO(),
			bookId: book.BookId{
				Id: bookId,
			},
			repoMock: func(ctx context.Context) *test.BookRepositoryMock {
				repo := test.BookRepositoryMock{}
				repo.On("GetBook", ctx, bookId).Return(&repository.BookEntity{State: repository.Available}, nil)
				repo.On("UpdateBook", ctx, repository.BookEntity{State: repository.Reserved}).Return(nil)

				return &repo
			},
			want:    &book.ReservationStatus{State: true},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repoMock := tt.repoMock(tt.ctx)
			b := NewBookGrpcService(repoMock)

			got, err := b.ReserveBook(tt.ctx, &tt.bookId)
			if err != tt.wantErr {
				t.Errorf("ReserveBook() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReserveBook() got = %v, want %v", got, tt.want)
			}

			repoMock.AssertExpectations(t)
		})
	}
}

func TestBookGrpcService_ReleaseBook(t *testing.T) {
	bookId := "1234"
	testError := errors.New("test")

	tests := []struct {
		name     string
		ctx      context.Context
		bookId   book.BookId
		repoMock func(ctx context.Context) *test.BookRepositoryMock
		want     *book.ReservationStatus
		wantErr  error
	}{
		{
			name: "book not found",
			ctx:  context.TODO(),
			bookId: book.BookId{
				Id: bookId,
			},
			repoMock: func(ctx context.Context) *test.BookRepositoryMock {
				repo := test.BookRepositoryMock{}
				repo.On("GetBook", ctx, bookId).Return(nil, repository.BookNotFound)

				return &repo
			},
			want:    nil,
			wantErr: repository.BookNotFound,
		},
		{
			name: "book update error",
			ctx:  context.TODO(),
			bookId: book.BookId{
				Id: bookId,
			},
			repoMock: func(ctx context.Context) *test.BookRepositoryMock {
				repo := test.BookRepositoryMock{}
				repo.On("GetBook", ctx, bookId).Return(&repository.BookEntity{Id: "1234"}, nil)
				repo.On("UpdateBook", ctx, repository.BookEntity{Id: "1234", State: repository.Available}).Return(testError)

				return &repo
			},
			want:    nil,
			wantErr: testError,
		},
		{
			name: "book release success",
			ctx:  context.TODO(),
			bookId: book.BookId{
				Id: bookId,
			},
			repoMock: func(ctx context.Context) *test.BookRepositoryMock {
				repo := test.BookRepositoryMock{}
				repo.On("GetBook", ctx, bookId).Return(&repository.BookEntity{State: repository.NotAvailable}, nil)
				repo.On("UpdateBook", ctx, repository.BookEntity{State: repository.Available}).Return(nil)

				return &repo
			},
			want:    &book.ReservationStatus{State: true},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repoMock := tt.repoMock(tt.ctx)
			b := NewBookGrpcService(repoMock)

			got, err := b.ReleaseBook(tt.ctx, &tt.bookId)
			if err != tt.wantErr {
				t.Errorf("ReleaseBook() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReleaseBook() got = %v, want %v", got, tt.want)
			}

			repoMock.AssertExpectations(t)
		})
	}
}
