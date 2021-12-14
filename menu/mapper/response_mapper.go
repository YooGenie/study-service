package mapper

import (
	"menu-service/dto/response"
	"menu-service/menu/entity"
)



func MakeMenuSummary(menu entity.Menu) (menuSummary response.MenuSummary) {
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