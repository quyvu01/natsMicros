package responses

type PaginationResponse[TResponse any] struct {
	Items            []TResponse
	TotalRecord      int64
	CurrentPageIndex int64
	TotalPage        int64
}
