package config

import (
	"context"
	"fmt"
	"time"

	"github.com/go-xorm/xorm"
	"xorm.io/core"
)

var (
	xormDb *xorm.Engine
)

//DB 생성
func ConfigureDatabase() DatabaseWrapper {
	dbConnection := Config.Database.ConnectionString

	xormDb, err := xorm.NewEngine(Config.Database.Driver, dbConnection) //Create Engine
	if err != nil {
		panic(fmt.Errorf("Database open error: error: %s \n", err))
	} else {
		fmt.Println("DB connected: ", Config.Database.Connection) //DB 연결 성공
	}

	//https://pkg.go.dev/database/sql#DB.SetMaxOpenConns 참고
	xormDb.SetMaxOpenConns(10)  //데이터베이스에 열린 연결의 최대 수가 설정
	xormDb.SetMaxIdleConns(5) //idle connection pool 의 최대 연결 수를 설정
	xormDb.SetConnMaxLifetime(10 * time.Minute) //연결을 재사용할 수 있는 최대 시간을 설정

	xormDb.ShowSQL(Config.Log.ShowSql) //로그 수준이 INFO보다 클 경우 로거에서 SQL 문인지 여부
	xormDb.Logger().SetLevel(core.LOG_INFO) //로거 레벨 설정

	return DatabaseWrapper{xormDb}
}

type DatabaseWrapper struct {
	*xorm.Engine
}

//세션 만드는 함수
func (d DatabaseWrapper) CreateSession(ctx context.Context) (*xorm.Session, context.Context) {
	session := d.NewSession()

	func(session interface{}, ctx context.Context) {
		if s, ok := session.(interface{ SetContext(context.Context) }); ok {
			s.SetContext(ctx)
		}
	}(session, ctx)
	defer session.Close() //풀로부터 연결 해지

	return session, context.WithValue(ctx, ContextDBKey, session)
}

func CleanUp() {
	xormDb.Close()
}
