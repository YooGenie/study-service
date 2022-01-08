package repository

import (
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

