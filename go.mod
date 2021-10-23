module "menu-service"

go 1.15

require (
	//github.com/360EntSecGroup-Skylar/excelize/v2 v2.3.1 // indirect
	github.com/aws/aws-sdk-go v1.35.35
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/go-playground/validator/v10 v10.4.1
	github.com/go-sql-driver/mysql v1.5.0
	github.com/go-xorm/xorm v0.7.9
	github.com/jinzhu/configor v1.2.0
	github.com/labstack/echo v3.3.10+incompatible
	github.com/labstack/gommon v0.3.0
	github.com/mattn/go-sqlite3 v1.14.6
	github.com/satori/go.uuid v1.2.0
	github.com/sirupsen/logrus v1.7.0
	github.com/stretchr/testify v1.6.1
	github.com/valyala/fasttemplate v1.2.1 // indirect
	golang.org/x/crypto v0.0.0-20200820211705-5c72a883971a
	golang.org/x/net v0.0.0-20201110031124-69a78807bb2b
	gopkg.in/testfixtures.v2 v2.6.0
	xorm.io/core v0.7.3
)
