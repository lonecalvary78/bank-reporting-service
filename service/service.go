package service

import (
	"bank-reporting-service/model/request"
	"bank-reporting-service/model/response"
)

type ReportGenerationService struct {
}

func NewGeneratingService() ReportGenerationService {
	return ReportGenerationService{}
}

func (service *ReportGenerationService) GenerateReporAsync(generateReportRequest request.GenerateReport, counter int64) (response.Ackowledgment, error) {

	return response.Ackowledgment{
		RequestID: counter,
		Status:    "SUCCESS",
	}, nil
}
