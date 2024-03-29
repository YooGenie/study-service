package dto

import (
	"strconv"

	"github.com/labstack/echo/v4"
)

type PageResult struct {
	Result     interface{} `json:"result"`
	TotalCount int64       `json:"totalCount"`
}

const PageSize = 20

type Pageable struct {
	Page     int
	PageSize int
}

func GetPageableFromRequest(ctx echo.Context) Pageable {
	page, err := strconv.Atoi(ctx.QueryParam("page"))
	if err != nil {
		page = 1
	}

	pageSize, err := strconv.Atoi(ctx.QueryParam("pageSize"))
	if err != nil {
		pageSize = PageSize
	}

	return Pageable{
		Page:     page,
		PageSize: pageSize,
	}
}
