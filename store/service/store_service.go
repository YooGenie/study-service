package service

import (
	"context"
	"menu-service/common"
	requestDto "menu-service/dto/request"
	responseDto "menu-service/dto/response"
	"menu-service/store/entity"
	"menu-service/store/repository"
	"sync"
)

var (
	storeServiceOnce     sync.Once
	storeServiceInstance *storeService
)

func StoreService() *storeService {
	storeServiceOnce.Do(func() {
		storeServiceInstance = &storeService{}
	})

	return storeServiceInstance
}

type storeService struct {
}

func (storeService) Create(ctx context.Context, creation requestDto.StoreCreate) (err error) {

	password, err := common.HashAndSalt(creation.Password)
	store := entity.Store{
		Id:                         creation.Id,
		Password:                   password,
		Mobile:                     common.SetEncrypt(creation.Mobile),
		BusinessRegistrationNumber: creation.BusinessRegistrationNumber,
		Created:                    nil,
		Updated:                    nil,
	}

	if err = store.ValidatePassword(password); err != nil {
		return
	}

	if err = store.Create(ctx); err != nil {
		return
	}

	return

}

func (storeService) GetStoreById(ctx context.Context, storeNo int64) (storeSummary responseDto.StoreSummary, err error) {
	storeSummary, err = repository.StoreRepository().FindById(ctx, storeNo)
	if err != nil {
		return
	}

	storeSummary.Mobile = common.GetDecrypt(storeSummary.Mobile)
	return
}

func (storeService) Update(ctx context.Context, edition requestDto.StoreUpdate) (err error) {

	password, err := common.HashAndSalt(edition.Password)
	store := entity.Store{
		No:                         edition.No,
		Id:                         edition.Id,
		Password:                   password,
		Mobile:                     common.SetEncrypt(edition.Mobile),
		BusinessRegistrationNumber: edition.BusinessRegistrationNumber,
		Created:                    nil,
		Updated:                    nil,
	}

	if err = store.ValidatePassword(password); err != nil {
		return
	}

	if err = store.Update(ctx); err != nil {
		return
	}

	return

}

func (storeService) GetStores(ctx context.Context, searchParams requestDto.SearchStoreQueryParams, pageable requestDto.Pageable) (results responseDto.PageResult, err error) {
	menus, totalCount, err := repository.StoreRepository().FindAll(ctx, searchParams, pageable)
	if err != nil {
		return
	}
	for i := 0; i < len(menus); i++ {
		menus[i].Mobile = common.GetDecrypt(menus[i].Mobile)
	}
	results = responseDto.PageResult{
		Result:     menus,
		TotalCount: totalCount,
	}

	return
}

func (storeService) Delete(ctx context.Context, storeNo int64) (err error) {
	deleteStore := entity.Store{No: storeNo}
	if err = deleteStore.Delete(ctx); err != nil {
		return err
	}

	return
}
