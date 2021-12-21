package repository

import (
	"context"
	"github.com/go-xorm/xorm"
	"menu-service/common"
	"menu-service/common/errors"
	responseDto "menu-service/dto/response"
	"sync"
)

var (
	storeRepositoryOnce     sync.Once
	storeRepositoryInstance *storeRepository
)

func StoreRepository() *storeRepository {
	storeRepositoryOnce.Do(func() {
		storeRepositoryInstance = &storeRepository{}
	})

	return storeRepositoryInstance
}

type storeRepository struct {
}

func (storeRepository) FindById(ctx context.Context, storeNo int64) (storeSummary responseDto.StoreSummary, err error) {

	queryBuilder := func() xorm.Interface {
		q := common.GetDB(ctx).Table("store")
		q.Where("1=1")
		q.And("store.no =?", storeNo)
		return q
	}

	has, err := queryBuilder().Get(&storeSummary)
	if err != nil {
		return
	}

	if has == false {
		err = errors.ErrNoResult
		return
	}

	return
}
