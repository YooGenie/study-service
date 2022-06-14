package dto

type Example struct {
	StartDate string `json:"startDate" validate:"date12,required"`
	EndDate   string `json:"endDate" validate:"date12,required"`
}
