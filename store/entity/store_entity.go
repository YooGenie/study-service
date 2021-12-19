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
	Id                         string          `xorm:"id"`
	Password                   string          `xorm:"password"`
	Mobile                     string          `xorm:"mobile"`
	BusinessRegistrationNumber string          `xorm:"business_registration_number"`
	Created                    json.RawMessage `xorm:"json 'created'"`
	Updated                    json.RawMessage `xorm:"json 'updated'" `
	DeletedAt                  time.Time       `xorm:"deleted_at"`
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

func (Store) TableName() string {
	return "store"
}
