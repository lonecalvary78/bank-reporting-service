package service

import (
	"bank-reporting-service/model/request"
	"testing"
)

func TestGenerateReporAsync(t *testing.T) {
	reportGenerationService := NewGeneratingService()
	_, error := reportGenerationService.GenerateReporAsync(request.GenerateReport{
		ReportCode: "RPT-001",
		Parameters: nil,
	}, 1)

	if error != nil {
		t.Errorf("There is some error occured during testing [error-message: %v]", error)
	}
}

func BenchmarkGenerateReporAsync(b *testing.B) {
	reportGenerationService := NewGeneratingService()
	for i := 0; i < b.N; i++ {
		counter := i + 1
		reportGenerationService.GenerateReporAsync(request.GenerateReport{
			ReportCode: "RPT-001",
			Parameters: nil,
		}, int64(counter))
	}
}
