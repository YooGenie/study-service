package dto

import "time"

//응답DTO
type MenuSummary struct {
	Id    int64  `json:"id" `
	Name  string `json:"name"`
	Price int64  `json:"price" `
	Description string    `json:"description" `
	CreatedBy   string     `json:"createdBy"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedBy   string     `json:"updatedBy"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
