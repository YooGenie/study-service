package service

import (
	"context"
	"study-service/entity"
	 "study-service/repository"



	"sync"

	log "github.com/sirupsen/logrus"
)

var (
	kakaoMemberServiceOnce     sync.Once
	kakaoMemberServiceInstance *kakaoMemberService
)

func KaKaoMemberService() *kakaoMemberService {
	kakaoMemberServiceOnce.Do(func() {
		kakaoMemberServiceInstance = &kakaoMemberService{}
	})

	return kakaoMemberServiceInstance
}

type kakaoMemberService struct {
}

//func (kakaoMemberService) GetMember(ctx context.Context, memberId int64) (memberSummary responseDto.MemberSummary, err error) {
//	member, err := repository.KaKaoMemberRepository().FindById(ctx, memberId)
//	if err != nil {
//		return
//	}
//
//
//
//	return
//}

func (kakaoMemberService) GetMemberByKakaoId(ctx context.Context, kakaoId int64) (member entity.KaKaoMember, err error) {
	member, err = repository.KaKaoMemberRepository().FindByKakaoId(ctx, kakaoId)
	if err != nil {
		return
	}

	return
}

func (kakaoMemberService) Create(ctx context.Context, member *entity.KaKaoMember) (int64, error) {
	return repository.KaKaoMemberRepository().Create(ctx, member)
}

func (kakaoMemberService) Update(ctx context.Context, member *entity.KaKaoMember) (err error) {
	return repository.KaKaoMemberRepository().Update(ctx, member)
}

//
//func (kakaoMemberService) UpdateMobile(ctx context.Context, memberId int64, updateParam requestDto.UpdateMemberParam) (err error) {
//	log.Traceln("")
//	// 1. Member Entity 조회
//	member, err := repository.KaKaoMemberRepository().FindById(ctx, memberId)
//	if err != nil {
//		return err
//	}
//	// 2. Member Entity 수정
//	member.UpdateMobile(updateParam)
//	// 3. Member Entity 저장
//	return repository.KaKaoMemberRepository().Update(ctx, &member)
//}

//func (kakaoMemberService) GetMembers(ctx context.Context, searchMemberQueryParams requestDto.SearchMemberQueryParams, pageable requestDto.Pageable) ([]entity.Member, int64, error) {
//	return repository.KaKaoMemberRepository().FindAll(ctx, searchMemberQueryParams, pageable)
//}

//func (kakaoMemberService) GetMaskedMember(ctx context.Context, memberId int64) (member responseDto.MaskedMemberSummary, err error) {
//	log.Traceln("")
//
//	member, err = repository.KaKaoMemberRepository().FindByIdMaskMobile(ctx, memberId)
//	if err != nil {
//		return
//	}
//
//	if member.Mobile != "" {
//		mobile := common.GetDecrypt(member.Mobile)
//		member.MemberMobile = common.MaskMobile(mobile)
//	}
//
//	return
//}

func (kakaoMemberService) Withdraw(ctx context.Context, id int64) error {
	log.Traceln("")
	// 1. member Id로 엔티티를 조회한다.
	member, err := repository.KaKaoMemberRepository().FindById(ctx, id)
	if err != nil {
		return err
	}
	// 2. Delete 처리 한다.
	member.Withdraw()
	// 3. 저장한다.
	return repository.KaKaoMemberRepository().Update(ctx, &member)
}
