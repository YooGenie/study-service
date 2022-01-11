package service

import (
	"context"
	"menu-service/common/errors"
	requestDto "menu-service/dto/request"
	"menu-service/security"
	"menu-service/store/entity"
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

func (authService) AuthWithSignIdPassword(ctx context.Context, signIn requestDto.AdminSignIn) (token security.JwtToken, err error) {
	storeEntity, err := storeService.StoreService().GetStoreById(ctx, signIn.Id)
	if err != nil {
		return
	}

	//비밀번호 유효성
	err = entity.Store{}.ValidatePassword(signIn.Password)
	if err != nil {
		err = errors.ErrAuthentication
		return
	}

	token, err = security.JwtAuthentication{}.GenerateJwtToken(security.UserClaim{
		Id:    storeEntity.Id,
		Name:  "유지니",
		Roles: "store",
	})

	return
}
