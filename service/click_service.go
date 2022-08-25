package service

import (
	"context"
	requestDto "study-service/dto/request"
	responseDto "study-service/dto/response"
	entity2 "study-service/entity"

	"sync"
	"time"
)

var (
	clickServiceOnce     sync.Once
	clickServiceInstance *clickService
)

func ClickService() *clickService {
	clickServiceOnce.Do(func() {
		clickServiceInstance = &clickService{}
	})

	return clickServiceInstance
}

type clickService struct {
}

func (clickService) Create(ctx context.Context, creation requestDto.ClickCreate) (err error) {
	if creation.Click == "" {
		return err
	}

	click := entity2.Click{
		CreatedAt: time.Now(),
	}


	if err = click.Create(ctx); err != nil {
		return
	}

	return

}

func (clickService) GetClicks(ctx context.Context, searchParams requestDto.SearchClickQueryParams, pageable requestDto.Pageable) (results responseDto.PageResult, err error) {
	clicks, totalCount, err := entity2.ClickRepository().FindAll(ctx, searchParams, pageable)
	if err != nil {
		return
	}
	results = responseDto.PageResult{
		Result:     clicks,
		TotalCount: totalCount,
	}

	return
}