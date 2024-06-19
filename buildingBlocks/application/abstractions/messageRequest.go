package abstractions

type IMessageRequest[TRequest any, TResponse any] interface {
	Request(msg TRequest) (TResponse, error)
}
