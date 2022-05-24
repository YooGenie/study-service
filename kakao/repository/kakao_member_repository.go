package repository

import (
	"context"
	"study-service/common"
	"study-service/common/errors"
	responseDto "study-service/dto/response"

	log "github.com/sirupsen/logrus"
	//responseDto "study-service/dto/response"
	"study-service/kakao/entity"
	"sync"

	"github.com/go-xorm/xorm"
)

var (
	kaKaoMemberRepositoryOnce     sync.Once
	kaKaoMemberRepositoryInstance *kaKaoMemberRepository
)

func KaKaoMemberRepository() *kaKaoMemberRepository {
	kaKaoMemberRepositoryOnce.Do(func() {
		kaKaoMemberRepositoryInstance = &kaKaoMemberRepository{}
	})

	return kaKaoMemberRepositoryInstance
}

type kaKaoMemberRepository struct {
}

func (kaKaoMemberRepository) FindById(ctx context.Context, id int64) (member entity.KaKaoMember, err error) {
	log.Traceln("")

	session := common.GetDB(ctx)
	member.Id = id

	has, err := session.Get(&member)
	if err != nil {
		err = errors.ErrAuthentication
		return
	}

	if has == false {
		msg := "등록된 회원 정보가 없습니다."
		log.Errorln(msg)
		err = errors.ErrNoResult
		return
	}

	return
}

func (kaKaoMemberRepository) FindByIdMaskMobile(ctx context.Context, id int64) (member responseDto.MaskedMemberSummary, err error) {
	log.Traceln("")

	member.Id = id
	queryBuilder := func() xorm.Interface {
		q := common.GetDB(ctx).Table("members")
		q.Select("id, nickname, mobile")
		q.Where("1=1")
		q.And("members.id = ?", id)

		return q
	}

	has, err := queryBuilder().Get(&member)

	if err != nil {
		err = errors.ErrProcessingImpossible
		log.Errorln(err)
		return
	}

	if has == false {
		msg := "등록된 회원 정보가 없습니다."
		log.Errorln(msg)
		err = errors.ErrNoResult
		return
	}

	return member, nil
}

func (repository kaKaoMemberRepository) FindByKakaoId(ctx context.Context, kakaoId int64) (member entity.KaKaoMember, err error) {
	log.Traceln("")
	session := common.GetDB(ctx).Where("withdraw_at is NULL")
	if member, err = repository.GetMemberFromDB(session, kakaoId); err != nil {
		return
	}
	return
}

func (repository kaKaoMemberRepository) FindByKakaoIdWithoutWithdraw(ctx context.Context, kakaoId int64) (member entity.KaKaoMember, err error) {
	log.Traceln("")
	session := common.GetDB(ctx).Desc("id")
	if member, err = repository.GetMemberFromDB(session, kakaoId); err != nil {
		return
	}
	return
}

func (kaKaoMemberRepository) GetMemberFromDB(session *xorm.Session, kakaoId int64) (member entity.KaKaoMember, err error) {
	log.Traceln("")
	member.KakaoId = kakaoId
	_, err = session.Get(&member)
	if err != nil {
		err = errors.ApiInternalError
		return
	}
	return
}

func (kaKaoMemberRepository) Create(ctx context.Context, member *entity.KaKaoMember) (memberId int64, err error) {
	log.Traceln("")

	session := common.GetDB(ctx)
	memberId, err = session.Insert(member)
	if err != nil {
		err = errors.ApiInternalError
		log.Errorln(err)
		return
	}
	memberId = member.Id

	return
}

func (kaKaoMemberRepository) Update(ctx context.Context, member *entity.KaKaoMember) (err error) {
	log.Traceln("")

	if _, err = common.GetDB(ctx).ID(member.Id).UseBool("term_text_message_agreed", "term_email_agreed").Update(member); err != nil {
		err = errors.ApiInternalError
		log.Errorln(err)
		return
	}

	return
}
