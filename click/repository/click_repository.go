package repository

import (
	"context"
	"github.com/go-xorm/xorm"
	log "github.com/sirupsen/logrus"
	"study-service/common"
	requestDto "study-service/dto/request"
	responseDto "study-service/dto/response"
	"sync"
)

var (
	clickRepositoryOnce     sync.Once
	clickRepositoryInstance *clickRepository
)

func ClickRepository() *clickRepository {
	clickRepositoryOnce.Do(func() {
		clickRepositoryInstance = &clickRepository{}
	})

	return clickRepositoryInstance
}

type clickRepository struct {
}

func (clickRepository) FindAll(ctx context.Context, searchParams requestDto.SearchClickQueryParams, pageable requestDto.Pageable) (results []responseDto.ClickSummary, totalCount int64, err error) {
	log.Traceln("")

	queryBuilder := func() xorm.Interface {
		q := common.GetDB(ctx).Table("click")
		q.Where("1=1")
		return q
	}

	if totalCount, err = queryBuilder().Limit(pageable.PageSize, pageable.Offset).Desc("click.id").FindAndCount(&results); err != nil {
		return
	}

	if totalCount == 0 {
		return
	}

	return
}
