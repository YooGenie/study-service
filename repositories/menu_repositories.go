package repositories

import (
	"context"
	"menu-service/common"
	"menu-service/dtos"
	"menu-service/entities"
	"sync"
)

var (
	menuRepositoryOnce     sync.Once
	menuRepositoryInstance *menuRepository
)

//싱글톤 쓰는 이유 : 최초한번만 메모리에 할당하기 위해서
func MenuRepository() *menuRepository {
	menuRepositoryOnce.Do(func() {
		menuRepositoryInstance = &menuRepository{}
	})

	return menuRepositoryInstance
}

type menuRepository struct {
}

//데이터삽입부분
func (menuRepository) Create(ctx context.Context, menu *entities.Menu) error {
	db := common.GetDB(ctx)
	_, err := db.Insert(menu)
	if err != nil {
		return err
	}

	return nil
}

func (menuRepository) FindById(ctx context.Context, Id int64) (entities.Menu, error) {
	var menu = entities.Menu{Id: Id}
	_, err := common.GetDB(ctx).Get(&menu)
	if err != nil {
		return menu, err
	}

	return menu, nil
}

func (menuRepository) FindAll(ctx context.Context, pageable dtos.Pageable) (menu []entities.Menu, totalCount int64, err error) {
	db := common.GetDB(ctx)
	offset := (pageable.Page - 1) * pageable.PageSize

	if totalCount, err = db.Desc("id").Limit(pageable.PageSize, offset).FindAndCount(&menu); err != nil {
		return menu, 0, err
	}
	if totalCount == 0 {
		return menu, 0, nil
	}

	return menu, totalCount, nil
}

func (menuRepository) Update(ctx context.Context, menu *entities.Menu) error {
	session := common.GetDB(ctx).ID(menu.Id)

	if _, err := session.Update(menu); err != nil {
		return err
	}

	return nil
}

func (menuRepository) Delete(ctx context.Context, menu *entities.Menu) error {
	session := common.GetDB(ctx).ID(menu.Id)

	if _, err := session.Delete(menu); err != nil {
		return err
	}

	return nil
}
