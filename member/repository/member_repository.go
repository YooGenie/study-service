package repository

import (
	"context"
	"study-service/common"
	"study-service/common/errors"
	responseDto "study-service/dto/response"
	"sync"
)

var (
	memberRepositoryOnce     sync.Once
	memberRepositoryInstance *memberRepository
)

func MemberRepository() *memberRepository {
	memberRepositoryOnce.Do(func() {
		memberRepositoryInstance = &memberRepository{}
	})

	return memberRepositoryInstance
}

type memberRepository struct {
}

func (memberRepository) FindById(ctx context.Context, email string) (memberSummary responseDto.MemberSummary, err error) {

	q := common.GetDB(ctx).Table("members").Where("email=?", email)

	has, err := q.Get(&memberSummary)
	if err != nil {
		return
	}

	if has == false {
		err = errors.ErrNoResult
		return
	}

	return
}


