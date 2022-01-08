package service

import (
	"context"
	"menu-service/common"
	requestDto "menu-service/dto/request"
	"menu-service/member/entity"
	"sync"
)

var (
	memberServiceOnce     sync.Once
	memberServiceInstance *memberService
)

func MemberService() *memberService {
	memberServiceOnce.Do(func() {
		memberServiceInstance = &memberService{}
	})

	return memberServiceInstance
}

type memberService struct {
}

func (memberService) Create(ctx context.Context, creation requestDto.MemberCreate) (err error) {

	password, err := common.HashAndSalt(creation.Password)
	member := entity.Member{
		Email:    creation.Email,
		Password: password,
		Mobile:   common.SetEncrypt(creation.Mobile),
		Name:     creation.Name,
		Nickname: creation.Nickname,
		Role: "MEMBER",
		Created:  nil,
		Updated:  nil,
	}

	if err = member.ValidatePassword(password); err != nil {
		return
	}

	if err = member.Create(ctx); err != nil {
		return
	}

	return

}
