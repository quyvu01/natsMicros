package abstractions

type IMessagePublisher[TMessage any] interface {
	Publish(msg TMessage) error
}
