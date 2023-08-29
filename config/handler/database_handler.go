package handler

import (
	"log"
	"net/http"
	"study-service/config"

	"github.com/labstack/echo/v4"
)

func CreateDatabaseContext(xormEngine config.DatabaseWrapper) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			req := c.Request()
			session, ctx := xormEngine.CreateSession(req.Context())

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
				if c.Response().Status >= 400 {
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
