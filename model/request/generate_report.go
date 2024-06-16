package request

type GenerateReport struct {
	ReportCode string      `json:"reportCode"`
	Parameters interface{} `json:"parameters"`
}
