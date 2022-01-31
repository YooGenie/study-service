package repository

import (
	"context"
	"github.com/go-xorm/xorm"
	log "github.com/sirupsen/logrus"
	"study-service/common"
	"study-service/common/errors"
	requestDto "study-service/dto/request"
	responseDto "study-service/dto/response"
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

func (storeRepository) FindByNo(ctx context.Context, storeNo int64) (storeSummary responseDto.StoreSummary, err error) {

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

func (storeRepository) FindAll(ctx context.Context, searchParams requestDto.SearchStoreQueryParams, pageable requestDto.Pageable) (results []responseDto.StoreSummary, totalCount int64, err error) {
	log.Traceln("")

	queryBuilder := func() xorm.Interface {
		q := common.GetDB(ctx).Table("store")
		q.Where("1=1")
		return q
	}

	if totalCount, err = queryBuilder().Limit(pageable.PageSize, pageable.Offset).Desc("store.id").FindAndCount(&results); err != nil {
		return
	}

	if totalCount == 0 {
		return
	}

	return
}

func (storeRepository) FindById(ctx context.Context, storeIn string) (storeSummary responseDto.StoreSummary, err error) {

	queryBuilder := func() xorm.Interface {
		q := common.GetDB(ctx).Table("store")
		q.Where("1=1")
		q.And("store.id =?", storeIn)
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
