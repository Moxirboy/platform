package utils

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

const (
	DefaultSize = 10
)

type PaginationQuery struct {
	Size    int    `json:"size"`
	Page    int    `json:"page"`
	OrderBy string `json:"orderBy"`
}

func (q *PaginationQuery) SetSize(Size string) error {
	if Size == "" {
		q.Size = DefaultSize
		return nil
	}
	n, err := strconv.Atoi(Size)
	if err != nil {
		return err
	}
	q.Size = n
	return nil
}

func (q *PaginationQuery) SetPage(Page string) error {
	if Page == "" {
		q.Size = 0
		return nil
	}
	n, err := strconv.Atoi(Page)
	if err != nil {
		return err
	}
	q.Page = n

	return nil
}

func (q *PaginationQuery) SetOrderBy(OrderBy string) error {
	if OrderBy == "" {
		q.OrderBy = "id"
	}
	q.OrderBy = OrderBy
	return nil
}
func (q *PaginationQuery) GetOrderBy() string {
	return q.OrderBy
}

func (q *PaginationQuery) GetOffSet() int {
	if q.Page == 0 {
		return 0
	}
	return (q.Page - 1) * q.Size
}

func (q *PaginationQuery) GetLimit() int {
	return q.Size
}

func (q *PaginationQuery) GetPage() int {
	return q.Page
}

func (q *PaginationQuery) GetSize() int {
	return q.Size
}

func (q *PaginationQuery) GetQueryString() string {
	return fmt.Sprintf("page=%v&size=%v&orderBy=%v", q.GetPage(), q.GetSize(), q.GetOrderBy())
}

func GetQueryFromCtx(ctx *gin.Context) (*PaginationQuery, error) {
	q := &PaginationQuery{}
	if err := q.SetPage(ctx.Query("page")); err != nil {
		return nil, err
	}
	if err := q.SetSize(ctx.Query("size")); err != nil {
		return nil, err
	}
	if err := q.SetOrderBy(ctx.Query("orderBy")); err != nil {
		return nil, err
	}

	return q, nil
}
