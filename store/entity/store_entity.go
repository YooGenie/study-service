package entity

import (
	"context"
	"encoding/json"
	"errors"
	log "github.com/sirupsen/logrus"
	"menu-service/common"
	"time"
)

type Store struct {
	No                         int64           `xorm:"no"`
	Id                         string          `xorm:"id"`
	Password                   string          `xorm:"password"`
	Mobile                     string          `xorm:"mobile"`
	BusinessRegistrationNumber string          `xorm:"business_registration_number"`
	Created                    json.RawMessage `xorm:"json 'created'"`
	Updated                    json.RawMessage `xorm:"json 'updated'" `
	DeletedAt                  time.Time       `xorm:"deleted" `
}

func (s *Store) Create(ctx context.Context) error {

	if rowsAffected, err := common.GetDB(ctx).Insert(s); err != nil {
		log.Errorln(err)
		return err
	} else if rowsAffected == 0 {
		msg := "가입이 되지 않았습니다. 시스템관리자에게 문의하여 주세요"
		log.Errorln(msg)
		return errors.New(msg)
	}
	return nil
}

func (s Store) ValidatePassword(password string) (err error) {
	if common.ComparePasswords(s.Password, password) {
		return err
	}

	return err

}

func (Store) TableName() string {
	return "store"
}

func (s *Store) Update(ctx context.Context) error {

	if rowsAffected, err := common.GetDB(ctx).Update(s); err != nil {
		log.Errorln(err)
		return err
	} else if rowsAffected == 0 {
		msg := "수정 되지 않았습니다. 시스템관리자에게 문의하여 주세요"
		log.Errorln(msg)
		return errors.New(msg)
	}
	return nil
}

func (s *Store) Delete(ctx context.Context) error {

	if rowsAffected, err := common.GetDB(ctx).Where("no=?", s.No).Delete(s); err != nil {
		log.Errorln(err)
		return err
	} else if rowsAffected == 0 {
		msg := "삭제가 반영되지 않았습니다. 시스템관리자에게 문의하여 주세요"
		log.Errorln(msg)
		return errors.New(msg)
	}

	return nil
}
