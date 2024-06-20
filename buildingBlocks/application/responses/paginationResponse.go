package responses

type PaginationResponse[TResponse any] struct {
	Items            []TResponse `json:"items"`
	TotalRecord      int64       `json:"totalRecord"`
	CurrentPageIndex int64       `json:"currentPageIndex"`
	TotalPage        int64       `json:"totalPage"`
}
