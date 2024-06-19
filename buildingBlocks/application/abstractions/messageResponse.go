package abstractions

type IMessageResponse[TRequest any, TResponse any] interface {
	Response(cb func(TRequest) (TResponse, error)) error
}
