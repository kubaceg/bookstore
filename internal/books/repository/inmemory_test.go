package repository

import (
	"context"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/kubaceg/bookstore/internal/common/log/test"
)

func TestInMemoryBookRepository_AddBook(t *testing.T) {
	books := []BookEntity{
		{
			Id:     "1234",
			Title:  "Book 1",
			Author: "Author 1",
			Isbn:   "1234",
		},
		{
			Id:     "1235",
			Title:  "Book 2",
			Author: "Author 2",
			Isbn:   "1234",
		},
	}

	tests := []struct {
		name          string
		books         map[string]BookEntity
		addedBook     BookEntity
		logMocks      map[string][]interface{}
		wantId        *string
		wantErr       error
		wantBooksList map[string]BookEntity
	}{
		{
			name:      "add book to empty database",
			books:     map[string]BookEntity{},
			addedBook: books[0],
			logMocks: map[string][]interface{}{
				"Info": {
					mock.Anything,
					"Book 1 book added to db",
				},
			},
			wantId:        &books[0].Id,
			wantErr:       nil,
			wantBooksList: map[string]BookEntity{"1234": books[0]},
		},
		{
			name:      "add book",
			books:     map[string]BookEntity{"1234": books[0]},
			addedBook: books[1],
			logMocks: map[string][]interface{}{
				"Info": {
					mock.Anything,
					"Book 2 book added to db",
				},
			},
			wantId:        &books[1].Id,
			wantErr:       nil,
			wantBooksList: map[string]BookEntity{"1234": books[0], "1235": books[1]},
		},
		{
			name:      "add existing book error",
			books:     map[string]BookEntity{"1234": books[0]},
			addedBook: books[0],
			logMocks: map[string][]interface{}{
				"Error": {
					mock.Anything,
					"Book 1 already exists in db",
				},
			},
			wantId:        nil,
			wantErr:       BookAlreadyExists,
			wantBooksList: map[string]BookEntity{"1234": books[0]},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logger := test.MockLogger{}
			for name := range tt.logMocks {
				logger.On(name, tt.logMocks[name]...)
			}

			i := &InMemoryBookRepository{
				books:  tt.books,
				logger: &logger,
			}

			gotId, gotErr := i.AddBook(context.TODO(), tt.addedBook)

			assert.Equal(t, tt.wantErr, gotErr)
			assert.Equal(t, tt.wantId, gotId)
			assert.Equal(t, tt.wantBooksList, i.books)

			logger.AssertExpectations(t)
		})
	}
}

func TestInMemoryBookRepository_GetBook(t *testing.T) {
	book := BookEntity{
		Id:     "1234",
		Title:  "Book 1",
		Author: "Author 1",
		Isbn:   "1234",
	}

	books := map[string]BookEntity{
		"1234": {
			Id:     "1234",
			Title:  "Book 1",
			Author: "Author 1",
			Isbn:   "1234",
		},
		"1235": book,
	}

	tests := []struct {
		name     string
		books    map[string]BookEntity
		id       string
		wantBook *BookEntity
		wantErr  error
	}{
		{
			name:     "empty database",
			books:    nil,
			id:       "1234",
			wantBook: nil,
			wantErr:  BookNotFound,
		},
		{
			name:     "product not found",
			books:    books,
			id:       "9999",
			wantBook: nil,
			wantErr:  BookNotFound,
		},
		{
			name:     "product found in database",
			books:    books,
			id:       "1235",
			wantBook: &book,
			wantErr:  nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &InMemoryBookRepository{
				books:  tt.books,
				logger: &test.MockLogger{},
			}
			gotBook, err := i.GetBook(context.TODO(), tt.id)
			if err != tt.wantErr {
				t.Errorf("GetBook() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotBook, tt.wantBook) {
				t.Errorf("GetBook() gotBook = %v, want %v", gotBook, tt.wantBook)
			}
		})
	}
}

func TestInMemoryBookRepository_GetBookList(t *testing.T) {
	books := []BookEntity{
		{
			Id:     "1234",
			Title:  "Book 1",
			Author: "Author 1",
			Isbn:   "1234",
		},
		{
			Id:     "1235",
			Title:  "Book 1",
			Author: "Author 1",
			Isbn:   "1234",
		},
	}

	booksMap := map[string]BookEntity{
		"1234": books[0],
		"1235": books[1],
	}

	tests := []struct {
		name      string
		books     map[string]BookEntity
		wantBooks []BookEntity
	}{
		{
			name:      "empty database",
			books:     nil,
			wantBooks: []BookEntity{},
		},
		{
			name:      "2 products in database",
			books:     booksMap,
			wantBooks: books,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &InMemoryBookRepository{
				books:  tt.books,
				logger: &test.MockLogger{},
			}
			gotBooks, _ := i.GetBookList(context.TODO(), BookListParams{})

			if !reflect.DeepEqual(gotBooks, tt.wantBooks) {
				t.Errorf("GetBookList() gotBook = %v, want %v", gotBooks, tt.wantBooks)
			}
		})
	}
}

func TestInMemoryBookRepository_ReserveBook(t *testing.T) {
	books := map[string]BookEntity{
		"1234": {
			Id:     "1234",
			Title:  "Book 1",
			Author: "Author 1",
			Isbn:   "1234",
			State:  Reserved,
		},
		"1235": {
			Id:     "1235",
			Title:  "Book 1",
			Author: "Author 1",
			Isbn:   "1234",
			State:  Available,
		},
		"1236": {
			Id:     "1236",
			Title:  "Book 1",
			Author: "Author 1",
			Isbn:   "1234",
			State:  NotAvailable,
		},
	}

	tests := []struct {
		name    string
		id      string
		wantErr error
	}{
		{
			name:    "reserve non existing book",
			id:      "999",
			wantErr: BookNotFound,
		},
		{
			name:    "reserve already reserved book",
			id:      "1234",
			wantErr: BookAlreadyReserved,
		},
		{
			name:    "reserve not available book",
			id:      "1236",
			wantErr: BookAlreadyReserved,
		},
		{
			name:    "reserve available book",
			id:      "1235",
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &InMemoryBookRepository{
				books:  books,
				logger: &test.MockLogger{},
			}
			err := i.ReserveBook(context.TODO(), tt.id)
			if err != tt.wantErr {
				t.Errorf("ReserveBook() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestInMemoryBookRepository_ReleaseBook(t *testing.T) {
	books := map[string]BookEntity{
		"1234": {
			Id:     "1234",
			Title:  "Book 1",
			Author: "Author 1",
			Isbn:   "1234",
			State:  Reserved,
		},
		"1235": {
			Id:     "1235",
			Title:  "Book 1",
			Author: "Author 1",
			Isbn:   "1234",
			State:  Available,
		},
		"1236": {
			Id:     "1236",
			Title:  "Book 1",
			Author: "Author 1",
			Isbn:   "1234",
			State:  NotAvailable,
		},
	}

	tests := []struct {
		name    string
		id      string
		wantErr error
	}{
		{
			name:    "release non existing book",
			id:      "999",
			wantErr: BookNotFound,
		},
		{
			name:    "release already reserved book",
			id:      "1234",
			wantErr: nil,
		},
		{
			name:    "release not available book",
			id:      "1236",
			wantErr: nil,
		},
		{
			name:    "release available book",
			id:      "1235",
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &InMemoryBookRepository{
				books:  books,
				logger: &test.MockLogger{},
			}
			err := i.ReleaseBook(context.TODO(), tt.id)
			if err != tt.wantErr {
				t.Errorf("ReleaseBook() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
