package eventbus

type Publisher interface {
	Publish(msg Message) error
}
