package responses

type PaginationResponse[TResponse any] struct {
	Items            []TResponse
	TotalRecord      int
	CurrentPageIndex int
	TotalPage        int
}
