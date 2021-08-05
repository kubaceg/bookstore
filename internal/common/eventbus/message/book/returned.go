package book

type Returned struct {
	ReservationId string `json:"reservation_id"`
	BookId        string `json:"book_id"`
}

func (c *Returned) GetTopic() string {
	return "bookstore.book.returned"
}
