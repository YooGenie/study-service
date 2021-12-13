package mapper

import (
	requestDto "menu-service/dto/request"
	"menu-service/menu/entity"
	"time"
)


func NewMenu(creation requestDto.MenuCreate) (menu entity.Menu, err error) {
	menu = entity.Menu{
		Name:        creation.Name,
		Price:       creation.Price,
		CreatedAt:   time.Now(),
		CreatedBy:   "1@naver.com",
		UpdatedAt:   time.Now(),
		UpdatedBy:   "1@naver.com",
		Description: creation.Description,
	}

	return
}

