package controller

import (
	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"study-service/common/errors"
	"study-service/pdf/service"
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
		makeHtml, _ := service.PdfService().MakeHtmlString(donationId)



	return ctx.JSON(http.StatusOK, makeHtml)

}