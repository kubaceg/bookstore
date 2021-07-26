package service

import (
	"context"
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
		wantErr  bool
	}{
		{
			name: "book reservation error",
			ctx:  context.TODO(),
			bookId: book.BookId{
				Id: bookId,
			},
			repoMock: func(ctx context.Context) *test.BookRepositoryMock {
				repo := test.BookRepositoryMock{}
				repo.On("ReserveBook", ctx, bookId).Return(repository.BookAlreadyReserved)

				return &repo
			},
			want:    &book.ReservationStatus{State: false},
			wantErr: true,
		},
		{
			name: "book reservation success",
			ctx:  context.TODO(),
			bookId: book.BookId{
				Id: bookId,
			},
			repoMock: func(ctx context.Context) *test.BookRepositoryMock {
				repo := test.BookRepositoryMock{}
				repo.On("ReserveBook", ctx, bookId).Return(nil)

				return &repo
			},
			want:    &book.ReservationStatus{State: true},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repoMock := tt.repoMock(tt.ctx)
			b := NewBookGrpcService(repoMock)

			got, err := b.ReserveBook(tt.ctx, &tt.bookId)
			if (err != nil) != tt.wantErr {
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

	tests := []struct {
		name     string
		ctx      context.Context
		bookId   book.BookId
		repoMock func(ctx context.Context) *test.BookRepositoryMock
		want     *book.ReservationStatus
		wantErr  bool
	}{
		{
			name: "book release error",
			ctx:  context.TODO(),
			bookId: book.BookId{
				Id: bookId,
			},
			repoMock: func(ctx context.Context) *test.BookRepositoryMock {
				repo := test.BookRepositoryMock{}
				repo.On("ReleaseBook", ctx, bookId).Return(repository.BookNotFound)

				return &repo
			},
			want:    &book.ReservationStatus{State: false},
			wantErr: true,
		},
		{
			name: "book release success",
			ctx:  context.TODO(),
			bookId: book.BookId{
				Id: bookId,
			},
			repoMock: func(ctx context.Context) *test.BookRepositoryMock {
				repo := test.BookRepositoryMock{}
				repo.On("ReleaseBook", ctx, bookId).Return(nil)

				return &repo
			},
			want:    &book.ReservationStatus{State: true},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repoMock := tt.repoMock(tt.ctx)
			b := NewBookGrpcService(repoMock)

			got, err := b.ReleaseBook(tt.ctx, &tt.bookId)
			if (err != nil) != tt.wantErr {
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
