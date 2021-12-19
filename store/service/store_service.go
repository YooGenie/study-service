package service

import (
	"context"
	requestDto "menu-service/dto/request"
	"menu-service/store/entity"
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

	store := entity.Store{
		Id: creation.Id,
		Password: creation.Password,
		Mobile: creation.Mobile,
		BusinessRegistrationNumber: creation.BusinessRegistrationNumber,
		Created: nil,
		Updated: nil,
	}

	if err = store.Create(ctx); err != nil {
		return
	}
	return

}