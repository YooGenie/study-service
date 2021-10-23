package factories

import (
	"context"
	"menu-service/dtos"
	"menu-service/entities"
	"time"
)

//Dto와 엔티티 맵핑하는 부분
func NewMenu(ctx context.Context, createResult dtos.MenuMake) (entities.Menu, error) {
	//userClaim := common.GetUserClaim(ctx)
	menu := entities.Menu{
		Name:        createResult.Name,
		Price:       createResult.Price,
		CreatedAt:   time.Now(),
		CreatedBy:   "1@naver.com",
		UpdatedAt:   time.Now(),
		UpdatedBy:   "1@naver.com",
		Description: createResult.Description,
	}

	return menu, nil
}
