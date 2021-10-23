package dtos

import (
	"time"
)

//요청DTO
type MenuMake struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Description string `json:"description"`
	//TestNumeric string `json:"testNumeric" validate:"numeric"`
}

//응답DTO
type ResponseId struct {
	Id int64
}

type MenuSummary struct {
	Id          int64     `json:"id"`
	Name        string    `json:"name"`
	Price       int       `json:"price"`
	Description string    `json:"description"`
	CreatedBy   string    `json:"createdBy"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedBy   string    `json:"updatedBy"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
