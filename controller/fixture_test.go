package controller

import (
	"fmt"

	menu "study-service/menu/entity"
	store "study-service/store/entity"

	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3"
	"gopkg.in/testfixtures.v2"
)

type DatabaseFixture struct {
}

func (DatabaseFixture) setUpDefault(xormEngine *xorm.Engine) {
	xormEngine.Sync2(
		new(menu.Menu),
		new(store.Store),
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
