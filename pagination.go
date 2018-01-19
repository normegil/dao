package dao

import "math"

type Pagination struct {
	offset int64
	limit  int64
}

func (p Pagination) Limit() int64 {
	limit := p.limit
	if limit <= 0 {
		limit = math.MaxInt64
	}
	return limit
}

func (p *Pagination) SetLimit(limit int64) {
	p.limit = limit
}

func (p Pagination) Offset() int64 {
	return p.offset
}

func (p *Pagination) SetOffset(offset int64) {
	p.offset = offset
}
