package entity

import (
	"context"
	"time"
)

type Menu struct {
	Id          int64     `xorm:"id pk autoincr"`
	Name        string    `xorm:"name"`
	Price       int64      `xorm:"price"`
	CreatedAt   time.Time `xorm:"created"`
	CreatedBy   string    `xorm:"created_by"`
	UpdatedAt   time.Time `xorm:"updated"`
	UpdatedBy   string    `xorm:"updated_by"`
	Description string    `xorm:"description"`
}

//func (m *Menu) UpdateMenu(ctx context.Context, updateResult dto2.MenuMake) {
//
//	m.Id = updateResult.Id
//	m.Name = updateResult.Name
//	m.Price = updateResult.Price
//	m.UpdatedAt = time.Now()
//	m.UpdatedBy = "1@naver.com"
//	m.Description = updateResult.Description
//
//}

func (m *Menu) ChangeUpdateBy(ctx context.Context) {
	m.UpdatedBy = "1"
}

func (Menu) TableName() string {
	return "menu"
}
