package controller

import (
	"fmt"
	"study-service/entity"

	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3"
	"gopkg.in/testfixtures.v2"
)

type DatabaseFixture struct {
}

func (DatabaseFixture) setUpDefault(xormEngine *xorm.Engine) {
	xormEngine.Sync2(
		new(entity.Menu),
		new(entity.Store),
	)

	fixtures, err := testfixtures.NewFolder(xormEngine.DB().DB, &testfixtures.SQLite{}, "../testdata/db_fixtures")
	fmt.Println("=== RUN DatabaseFixture.setUpDefault")

	if err != nil {
		panic(err)
	}
	testfixtures.SkipDatabaseNameCheck(true)

	if err := fixtures.Load(); err != nil {
		panic(err)
	}
	fmt.Println("=== FINISH DatabaseFixture.setUpDefault")
}
