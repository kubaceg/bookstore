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

func (b *BookGrpcService) AddBook(ctx context.Context, b2 *book.Book) (*book.BookId, error) {
	entity := repository.BookEntity{
		Id:     b2.Id,
		Title:  b2.Title,
		Author: b2.Author,
		Isbn:   b2.Isbn,
	}

	id, err := b.repo.AddBook(ctx, entity)

	if err != nil {
		return nil, err
	}

	return &book.BookId{Id: *id}, nil
}

func (b *BookGrpcService) GetBook(ctx context.Context, id *book.BookId) (entity *book.Book, err error) {
	bookEntity, err := b.repo.GetBook(ctx, id.Id)

	if err != nil {
		return nil, err
	}

	entity = &book.Book{
		Id:     bookEntity.Id,
		Title:  bookEntity.Title,
		Author: bookEntity.Author,
		Isbn:   bookEntity.Isbn,
	}

	return
}

func (b *BookGrpcService) GetBookList(ctx context.Context, empty *emptypb.Empty) (*book.BookList, error) {
	panic("implement me")
}
