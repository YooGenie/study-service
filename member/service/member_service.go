package service

import (
	"context"
	"study-service/common"
	requestDto "study-service/dto/request"
	responseDto "study-service/dto/response"

	//responseDto "study-service/dto/response"
	"study-service/member/entity"
	"study-service/member/repository"
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

func (memberService) GetMemberById(ctx context.Context, email string) (memberSummary responseDto.MemberSummary, err error) {
	memberSummary, err = repository.MemberRepository().FindById(ctx, email)
if err != nil {
return
}
	memberSummary.Mobile = common.GetDecrypt(memberSummary.Mobile)
return
}
