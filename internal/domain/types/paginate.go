package types

type PaginateOptions struct {
	Page    int
	Limit   int
	Keyword string
	Filters map[string]interface{}
	OrderBy string
	Order   string // "asc" or "desc"
}
