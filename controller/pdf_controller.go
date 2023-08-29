package controller

import (
	"strconv"
	"study-service/common/errors"
	service2 "study-service/service"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

type PdfController struct {
}

func (controller PdfController) Init(g *echo.Group) {
	g.GET("/:donationId", controller.GetPdf)
}

func (PdfController) GetPdf(ctx echo.Context) error {

	donationId, err := strconv.ParseInt(ctx.Param("donationId"), 10, 64)
	if err != nil {
		log.Errorf("GetDonationReceipt Error:  %s", err.Error())
		return errors.ApiParamValidError(err)
	}

	//makeHtml, _ := service.PdfService().MakeHtmlByte(donationId)
	makeHtml, _ := service2.PdfService().MakeHtmlString(donationId)

	return ctx.File(makeHtml)

}
