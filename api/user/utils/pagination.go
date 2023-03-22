package utils

import "math"

type Pagination struct {
	PerPage     int   `json:"per_page"`
	CurrentPage int   `json:"current_page"`
	Total       int64 `json:"total"`
	TotalPages  int   `json:"total_pages"`
}

func (p *Pagination) GetOffset() int {
	return (p.CurrentPage - 1) * p.PerPage
}

func (p *Pagination) GetLimit() int {
	return p.PerPage
}

func (p *Pagination) SetTotalPages() {
	p.TotalPages = int(math.Ceil(float64(p.Total) / float64(p.PerPage)))
}
