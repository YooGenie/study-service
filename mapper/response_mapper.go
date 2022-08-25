package mapper

import (
	response "study-service/dto/response"
	entity2 "study-service/entity"
)

func MakeMenuSummary(menu entity2.Menu) (menuSummary response.MenuSummary) {
	menuSummary = response.MenuSummary{
		Id:          menu.Id,
		Name:        menu.Name,
		Price:       menu.Price,
		CreatedBy:   menu.CreatedBy,
		CreatedAt:   menu.CreatedAt,
		UpdatedBy:   menu.UpdatedBy,
		UpdatedAt:   menu.UpdatedAt,
		Description: menu.Description,
	}

	return
}

func MakeMenuSummaries(menus []entity2.Menu) (menuSummary []response.MenuSummary) {
	for _, menu := range menus {
		menuSummary = append(menuSummary, MakeMenuSummary(menu))
	}

	return
}
