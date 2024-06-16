package application

import (
	"bank-reporting-service/model/request"
	"bank-reporting-service/service"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ReportApplication struct {
	counter                 int64
	reportGeneratingService service.ReportGenerationService
}

func NewApplication() ReportApplication {

	return ReportApplication{
		counter:                 0,
		reportGeneratingService: service.NewGeneratingService(),
	}
}

func (reportApplication *ReportApplication) GenerateReport(ctx echo.Context) error {
	var userInput = echo.Map{}
	error := ctx.Bind(&userInput)

	reportCode := fmt.Sprintf("%v", userInput["reportCode"])

	generateReportRequest := request.GenerateReport{
		ReportCode: reportCode,
		Parameters: userInput["parameters"],
	}

	if error != nil {
		return ctx.String(http.StatusBadRequest, error.Error())
	}

	reportApplication.counter = reportApplication.counter + 1
	acknowledgment, error := reportApplication.reportGeneratingService.GenerateReporAsync(generateReportRequest, reportApplication.counter)
	if error != nil {
		return error
	}

	bytes, error := json.Marshal(acknowledgment)
	if error != nil {
		return error
	}

	return ctx.String(http.StatusOK, string(bytes))
}
