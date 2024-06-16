package main

import (
	"bank-reporting-service/application"

	"github.com/labstack/echo/v4"
)

func main() {
	server := echo.New()
	reportApplication := application.NewApplication()
	server.POST("/reports/async",reportApplication.GenerateReport)
	server.Logger.Fatal(server.Start(":8080"))
}
