package service

import (
	"context"
	requestDto "study-service/dto/request"
	responseDto "study-service/dto/response"
	repository2 "study-service/repository"

	"study-service/mapper"
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
	if err = repository2.MenuRepository().Create(ctx, &newMenu); err != nil {
		return err
	}
	return err

}

func (menuService) GetMenuById(ctx context.Context, Id int64) (menuSummary responseDto.MenuSummary, err error) {
	menu, err := repository2.MenuRepository().FindById(ctx, Id)
	if err != nil {
		return menuSummary, err
	}

	menuSummary = mapper.MakeMenuSummary(menu)

	return
}

func (menuService) GetMenu(ctx context.Context, pageable requestDto.Pageable) (results responseDto.PageResult, err error) {
	menus, totalCount, err := repository2.MenuRepository().FindAll(ctx, pageable)

	menuSummaries := mapper.MakeMenuSummaries(menus)

	results = responseDto.PageResult{
		Result:     menuSummaries,
		TotalCount: totalCount,
	}

	return
}

func (menuService) UpdateMenu(ctx context.Context, edition requestDto.MenuUpdate) (err error) {
	menu, err := repository2.MenuRepository().FindById(ctx, edition.Id)
	if err != nil {
		return
	}

	err = mapper.UpdateMenu(edition, &menu)
	if err != nil {
		return err
	}

	if err = repository2.MenuRepository().Update(ctx, &menu); err != nil {
		return  err
	}

	return
}

func (menuService) DeleteMenu(ctx context.Context, Id int64) error {
	menu, err := repository2.MenuRepository().FindById(ctx, Id)
	if err != nil {
		return err
	}

	return repository2.MenuRepository().Delete(ctx, &menu)
}
