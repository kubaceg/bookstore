package service

import (
	context "context"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/kubaceg/bookstore/internal/books/repository"
	"github.com/kubaceg/bookstore/internal/common/genproto/book"
)

type BookGrpcService struct {
	book.UnimplementedBookServiceServer

	repo repository.BookRepository
}

func NewBookGrpcService(r repository.BookRepository) book.BookServiceServer {
	return &BookGrpcService{repo: r}
}

func (s *BookGrpcService) AddBook(ctx context.Context, b *book.Book) (*book.BookId, error) {
	entity := repository.BookEntity{
		Id:     b.Id,
		Title:  b.Title,
		Author: b.Author,
		Isbn:   b.Isbn,
		State:  repository.BookState(b.State),
	}

	id, err := s.repo.AddBook(ctx, entity)

	if err != nil {
		return nil, err
	}

	return &book.BookId{Id: *id}, nil
}

func (s *BookGrpcService) GetBook(ctx context.Context, id *book.BookId) (entity *book.Book, err error) {
	bookEntity, err := s.repo.GetBook(ctx, id.Id)

	if err != nil {
		return nil, err
	}

	entity = &book.Book{
		Id:     bookEntity.Id,
		Title:  bookEntity.Title,
		Author: bookEntity.Author,
		Isbn:   bookEntity.Isbn,
		State:  book.State(bookEntity.State),
	}

	return
}

func (s *BookGrpcService) GetBookList(ctx context.Context, empty *emptypb.Empty) (*book.BookList, error) {
	var books []*book.Book

	list, err := s.repo.GetBookList(ctx, repository.BookListParams{})
	if err != nil {
		return nil, err
	}

	for _, b := range list {
		entity := &book.Book{
			Id:     b.Id,
			Title:  b.Title,
			Author: b.Author,
			Isbn:   b.Isbn,
			State:  book.State(b.State),
		}
		books = append(books, entity)
	}

	return &book.BookList{Books: books}, nil
}

func (s *BookGrpcService) ReserveBook(ctx context.Context, bookId *book.BookId) (*book.ReservationStatus, error) {
	status := &book.ReservationStatus{}

	entity, err := s.repo.GetBook(ctx, bookId.Id)
	if err != nil {
		return nil, err
	}

	if book.State(entity.State) != book.State_AVAILABLE {
		return nil, repository.BookAlreadyReserved
	}

	entity.State = repository.Reserved

	err = s.repo.UpdateBook(ctx, *entity)
	if err != nil {
		return nil, err
	}

	status.State = true

	return status, nil
}

func (s *BookGrpcService) ReleaseBook(ctx context.Context, bookId *book.BookId) (*book.ReservationStatus, error) {
	status := &book.ReservationStatus{}

	entity, err := s.repo.GetBook(ctx, bookId.Id)
	if err != nil {
		return nil, err
	}

	entity.State = repository.Available

	err = s.repo.UpdateBook(ctx, *entity)
	if err != nil {
		return nil, err
	}

	status.State = true

	return status, nil
}
