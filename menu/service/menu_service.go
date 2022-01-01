package service

import (
	"context"
	requestDto "menu-service/dto/request"
	responseDto "menu-service/dto/response"

	"menu-service/menu/mapper"
	"menu-service/menu/repository"
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

func (menuService) CreateMenu(ctx context.Context, creation requestDto.MenuCreate) (err error) {
	newMenu, err := mapper.NewMenu(creation)
	if err != nil {
		return
	}
	if err = repository.MenuRepository().Create(ctx, &newMenu); err != nil {
		return err
	}
	return err

}

func (menuService) GetMenuById(ctx context.Context, Id int64) (menuSummary responseDto.MenuSummary, err error) {
	menu, err := repository.MenuRepository().FindById(ctx, Id)
	if err != nil {
		return menuSummary, err
	}

	menuSummary = mapper.MakeMenuSummary(menu)

	return
}

func (menuService) GetMenu(ctx context.Context, pageable requestDto.Pageable) (results responseDto.PageResult, err error) {
	menus, totalCount, err := repository.MenuRepository().FindAll(ctx, pageable)

	menuSummaries := mapper.MakeMenuSummaries(menus)

	results = responseDto.PageResult{
		Result:     menuSummaries,
		TotalCount: totalCount,
	}

	return
}

func (menuService) UpdateMenu(ctx context.Context, edition requestDto.MenuUpdate) (err error) {
	menu, err := repository.MenuRepository().FindById(ctx, edition.Id)
	if err != nil {
		return
	}

	err = mapper.UpdateMenu(edition, &menu)
	if err != nil {
		return err
	}

	if err = repository.MenuRepository().Update(ctx, &menu); err != nil {
		return  err
	}

	return
}

func (menuService) DeleteMenu(ctx context.Context, Id int64) error {
	menu, err := repository.MenuRepository().FindById(ctx, Id)
	if err != nil {
		return err
	}

	return repository.MenuRepository().Delete(ctx, &menu)
}
