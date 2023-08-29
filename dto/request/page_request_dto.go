package dto

import (
	"strconv"

	"github.com/labstack/echo/v4"
)

const Page = 1
const PageSize = 10

type Pageable struct {
	Offset   int
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

	offset := (page - 1) * pageSize

	return Pageable{
		Offset:   offset,
		Page:     page,
		PageSize: pageSize,
	}
}
