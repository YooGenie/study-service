package service

import (
	"context"
	"menu-service/common"
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
