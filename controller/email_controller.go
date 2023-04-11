package controller

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"study-service/service"
)

type EmailController struct {
}

func (controller EmailController) Init(g *echo.Group) {
	g.GET("/send", controller.SendMessage)
}

func (EmailController) SendMessage(ctx echo.Context) error {

	err := service.EmailService().SendMessage(ctx.Request().Context())
	if err != nil {
		fmt.Println(err)
		return err
	}

	return ctx.NoContent(http.StatusOK)

}
