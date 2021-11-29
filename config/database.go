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

func ConfigureDatabase() DatabaseWrapper {
	dbConnection := Config.Database.ConnectionString

	xormDb, err := xorm.NewEngine(Config.Database.Driver, dbConnection)
	if err != nil {
		panic(fmt.Errorf("Database open error: error: %s \n", err))
	} else {
		fmt.Println("DB connected: ", Config.Database.Connection)
	}

	xormDb.SetMaxOpenConns(10)
	xormDb.SetMaxIdleConns(5)
	xormDb.SetConnMaxLifetime(10 * time.Minute)

	//xormDb.ShowSQL(Config.Log.ShowSql)
	xormDb.Logger().SetLevel(core.LOG_INFO)

	return DatabaseWrapper{xormDb}
}

type DatabaseWrapper struct {
	*xorm.Engine
}

func (d DatabaseWrapper) CreateSession(ctx context.Context) (*xorm.Session, context.Context) {
	session := d.NewSession()

	func(session interface{}, ctx context.Context) {
		if s, ok := session.(interface{ SetContext(context.Context) }); ok {
			s.SetContext(ctx)
		}
	}(session, ctx)
	defer session.Close()

	return session, context.WithValue(ctx, ContextDBKey, session)
}

func CleanUp() {
	xormDb.Close()
}
