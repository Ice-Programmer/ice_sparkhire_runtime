package utils

const (
	DefaultPageNum  = 1
	DefaultPageSize = 10
	MaxPageSize     = 300
)

func SetPageDefault(pageSize, pageNum int32) (size int32, num int32) {
	if pageNum <= 0 {
		pageNum = DefaultPageNum
	}
	if pageSize <= 0 {
		pageSize = DefaultPageSize
	}
	if pageSize > MaxPageSize {
		pageSize = MaxPageSize
	}
	return pageSize, pageNum
}
