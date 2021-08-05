package reservations

type Created struct {
	ReservationId string `json:"reservation_id"`
	BookId        string `json:"book_id"`
}

func (c Created) GetTopic() string {
	return "bookstore.reservation.created"
}
