package service

import (
	"context"
	"study-service/common/errors"
	requestDto "study-service/dto/request"
	"study-service/member/entity"
	memberService "study-service/member/service"
	"study-service/security"
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

func (authService) AuthWithSignIdPassword(ctx context.Context, signIn requestDto.AdminSignIn) (token security.JwtToken, err error)  {
	memberEntity, err := memberService.MemberService().GetMemberById(ctx, signIn.Email)
	if err != nil {
		return
	}

	//비밀번호 유효성
	err = entity.Member{}.ValidatePassword(signIn.Password)
	if err != nil {
		err = errors.ErrAuthentication
		return
	}

	token, err = security.JwtAuthentication{}.GenerateJwtToken(security.UserClaim{
		Id:    memberEntity.Email,
		Name:  "유지니",
		Roles: memberEntity.Role,
	})

	return
}
