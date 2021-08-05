package eventbus

type Message interface {
	GetTopic() string
}
