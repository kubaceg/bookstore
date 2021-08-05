package reservations

type Error struct {
	BookId string `json:"book_id"`
}

func (c *Error) GetTopic() string {
	return "bookstore.reservation.error"
}
