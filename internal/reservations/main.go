package main

import (
	"fmt"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/kubaceg/bookstore/internal/common/genproto/book"
)

func main() {
	serverAddr := "127.0.0.1:8081"

	conn, err := grpc.Dial(serverAddr, grpc.WithInsecure())
	defer conn.Close()

	if err != nil {
		panic(err)
	}

	bookClient := book.NewBookServiceClient(conn)

	entity := book.Book{
		Id:     "1234",
		Title:  "Book 1",
		Author: "Kuba",
		Isbn:   "1234",
		State:  book.State_AVAILABLE,
	}

	id, err := bookClient.AddBook(context.TODO(), &entity)
	if err != nil {
		panic(err)
	}

	fmt.Println(id)

	returnedBook, err := bookClient.GetBook(context.TODO(), &book.BookId{Id: "1234"})
	if err != nil {
		panic(err)
	}

	fmt.Printf("Returned book: %+v\n", returnedBook)

	list, err := bookClient.GetBookList(context.TODO(), &emptypb.Empty{})
	if err != nil {
		panic(err)
	}

	fmt.Printf("Returned book list: %+v\n", list)
}
