package repository

import (
	"context"
	"github.com/go-xorm/xorm"
	"study-service/common"
	"study-service/common/errors"
	requestDto "study-service/dto/request"
	"study-service/menu/entity"
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
func (menuRepository) Create(ctx context.Context, menu *entity.Menu) error {
	db := common.GetDB(ctx)
	_, err := db.Insert(menu)
	if err != nil {
		return err
	}

	return nil
}

func (menuRepository) FindById(ctx context.Context, Id int64) (entity.Menu, error) {
	var menu = entity.Menu{Id: Id}
	has, err := common.GetDB(ctx).Get(&menu)
	 if !has {
		 err = errors.ErrNoResult
		 return menu, err
	 }

	if err != nil {
		return menu, err
	}

	return menu, err
}

func (menuRepository) FindAll(ctx context.Context, pageable requestDto.Pageable) (menus []entity.Menu, totalCount int64, err error) {

	queryBuilder := func() xorm.Interface {
		q := common.GetDB(ctx).Table("menu").Select("menu.*").Where("1=1")

		return q
	}

	var results []struct {
		entity.Menu `xorm:"extends"`
	}

	if totalCount, err = queryBuilder().Limit(pageable.PageSize).Desc("menu.id").FindAndCount(&results); err != nil {
		return
	}
	if totalCount == 0 {
		return nil, 0, err
	}

	for _, result := range results {
		var menu = entity.Menu{}
		menu = result.Menu
		menus = append(menus, menu)
	}

	return menus, totalCount, err
}

func (menuRepository) Update(ctx context.Context, menu *entity.Menu) error {
	session := common.GetDB(ctx).ID(menu.Id)

	if _, err := session.Update(menu); err != nil {
		return err
	}

	return nil
}

func (menuRepository) Delete(ctx context.Context, menu *entity.Menu) error {
	session := common.GetDB(ctx).ID(menu.Id)

	if _, err := session.Delete(menu); err != nil {
		return err
	}

	return nil
}
