package response

type Ackowledgment struct {
	RequestID int64  `json:"requestId"`
	Status    string `json:"status"`
}
