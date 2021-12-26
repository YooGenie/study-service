package service

import (
	"context"
	"menu-service/common/errors"
	requestDto "menu-service/dto/request"
	storeService "menu-service/store/service"
	"sync"
)

var (
	authServiceOnce     sync.Once
	authServiceInstance *authService
)

func AuthService() *authService {
	authServiceOnce.Do(func() {
		authServiceInstance = &authService{}
	})
	return authServiceInstance
}

type authService struct {
}

func (authService) AuthWithSignIdPassword(ctx context.Context, signIn requestDto.AdminSignIn) (err error) {
	storeEntity, err := storeService.StoreService().GetStoreById(ctx, signIn.Id)
	if err != nil {
		return
	}

	check := signIn.Id == storeEntity.Id
	if !check {
		err = errors.ErrAuthorization
		return
	}

	return
}
