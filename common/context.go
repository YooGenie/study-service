package common

import (
	"github.com/go-xorm/xorm"
	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"net/http"
)

const (
	DateLayout6       = "060102"
	DateLayout8       = "20060102"
	DateLayout10      = "2006-01-02"
	DateLayout14      = "20060102150405"
	DateLayout19      = "2006-01-02 15:04:05"
	UserRoleMallAdmin = "mall-admin"
)

const ContextDBKey = "DB"
const ContextLoggingDBKey = "LOGGING_DB"

func InitContextDB(xormEngine *xorm.Engine) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()
			ctx := req.Context()

			session := NewSession(ctx, xormEngine)
			defer session.Close()

			ctx = context.WithValue(ctx, ContextDBKey, session)

			//loggingSession은 트렌젝션 처리 하지 않음
			loggingSession := NewSession(ctx, xormEngine)
			defer loggingSession.Close()
			ctx = context.WithValue(ctx, ContextLoggingDBKey, loggingSession)
			c.SetRequest(req.WithContext(ctx))

			switch req.Method {
			case "POST", "PUT", "DELETE", "PATCH":
				if err := session.Begin(); err != nil {
					log.Println(err)
					return err
				}
				if err := next(c); err != nil {
					session.Rollback()
					return err
				}
				if c.Response().Status >= 500 {
					session.Rollback()

					// 처리 결과 에러 발생 시 rollback만 처리하고 여기서는 에러를 반환하지 않음
					// 이미 http 결과(c.Resopnse)에 에러 관련 정보가 담겨 있음
					return nil
				}
				if err := session.Commit(); err != nil {
					return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
				}
			default:
				return next(c)
			}

			// 여기까지는 오지 않음. 따라서 그냥 nil 반환
			return nil
		}
	}
}

func NewSession(ctx context.Context, xormEngine *xorm.Engine) *xorm.Session {
	session := xormEngine.NewSession()

	func(session interface{}, ctx context.Context) {
		if s, ok := session.(interface{ SetContext(context.Context) }); ok {
			s.SetContext(ctx)
		}
	}(session, ctx)

	return session
}

func GetDB(ctx context.Context) *xorm.Session {
	v := ctx.Value(ContextDBKey)
	if v == nil {
		panic("DB is not exist")
	}
	if db, ok := v.(*xorm.Session); ok {
		return db
	}
	if db, ok := v.(*xorm.Engine); ok {
		return db.NewSession()
	}
	panic("DB is not exist")
}
