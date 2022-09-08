package httpenc

type Pagination struct {
	Index int   `json:"index"`
	Limit int   `json:"limit"`
	Total int64 `json:"total"`
	Prev  bool  `json:"prev"`
	Next  bool  `json:"next"`
}

func (pagination *Pagination) LimitOrOffset() (limit, offset int) {
	index := pagination.Index
	if index == 0 {
		pagination.Index = 1
	} else if index == 1 {
		index--
	}

	limit = pagination.Limit
	if limit == 0 {
		pagination.Limit = 12
		limit = pagination.Limit
	}

	offset = index * limit
	return
}

func (pagination *Pagination) PrevOrNext(total int64) {
	pagination.Total = total

	if pagination.Index > 1 {
		pagination.Prev = true
	}

	if pagination.Total > int64(pagination.Index)*int64(pagination.Limit) {
		pagination.Next = true
	}
}
