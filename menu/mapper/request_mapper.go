package mapper

import (
	log "github.com/sirupsen/logrus"
	requestDto "study-service/dto/request"
	"study-service/menu/entity"
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

func UpdateMenu(edition requestDto.MenuUpdate,  menu *entity.Menu) (err error) {
	log.Infoln("")

	menu.Id =edition.Id
	menu.Name=edition.Name
	menu.Price=edition.Price
	menu.Description=edition.Description

	return
}
