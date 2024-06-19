package abstractions

type IMessageConsumer[TMessage any] interface {
	Consume(cb func(TMessage, error)) error
}
