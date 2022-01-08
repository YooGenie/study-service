package entity

import (
	"context"
	"encoding/json"
	"errors"
	log "github.com/sirupsen/logrus"
	"menu-service/common"
	"time"
)

type Member struct {
	Id        int64           `xorm:"id"`
	Email     string          `xorm:"email"`
	Password  string          `xorm:"password"`
	Mobile    string          `xorm:"mobile"`
	Name      string          `xorm:"name"`
	Nickname  string          `xorm:"nickname"`
	Role      string          `xorm:"role"`
	Created                    json.RawMessage `xorm:"json 'created'"`
	Updated                    json.RawMessage `xorm:"json 'updated'" `
	DeletedBy json.RawMessage `xorm:"json 'deleted_by'"`
	DeletedAt time.Time       `xorm:"deleted_at" `
}

func (s *Member) Create(ctx context.Context) error {

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

func (s Member) ValidatePassword(password string) (err error) {
	if common.ComparePasswords(s.Password, password) {
		return err
	}

	return err

}

func (Member) TableName() string {
	return "members"
}
