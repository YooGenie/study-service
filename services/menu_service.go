package services

import (
	"context"
	"menu-service/dtos"
	"menu-service/entities"
	"menu-service/factories"
	"menu-service/repositories"
	"sync"
)

var (
	menuServiceOnce     sync.Once
	menuServiceInstance *menuService
)

func MenuService() *menuService {
	menuServiceOnce.Do(func() {
		menuServiceInstance = &menuService{}
	})

	return menuServiceInstance
}

type menuService struct {
}

func (menuService) CreateMenu(ctx context.Context, createResult dtos.MenuMake) (int64, error) {
	menu, err := factories.NewMenu(ctx, createResult)
	if err != nil {
		return 0, err
	}
	if menu.Name == "" || menu.Price == 0 || len(menu.Name) > 100 {
		return 0, err
	}
	if err := repositories.MenuRepository().Create(ctx, &menu); err != nil {
		return 0, err
	}

	return menu.Id, nil
}

func (menuService) GetMenuById(ctx context.Context, Id int64) (entities.Menu, error) {
	return repositories.MenuRepository().FindById(ctx, Id)
}

func (menuService) GetMenu(ctx context.Context, pageable dtos.Pageable) ([]entities.Menu, int64, error) {
	return repositories.MenuRepository().FindAll(ctx, pageable)
}

func (menuService) UpdateMenu(ctx context.Context, menuMake dtos.MenuMake) (int64, error) {
	menu, err := repositories.MenuRepository().FindById(ctx, menuMake.Id)
	if err != nil {
		return 0, err
	}

	menu.UpdateMenu(ctx, menuMake)

	if err := repositories.MenuRepository().Update(ctx, &menu); err != nil {
		return 0, err
	}

	return menu.Id, nil
}

func (menuService) DeleteMenu(ctx context.Context, Id int64) error {
	menu, err := repositories.MenuRepository().FindById(ctx, Id)
	if err != nil {
		return err
	}

	return repositories.MenuRepository().Delete(ctx, &menu)
}
