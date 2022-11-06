package model

const (
	statusSuccess = 1000
)

// DefaultResponse used in every response
type DefaultResponse struct {
	StatusCode int         `json:"statusCode"`
	Data       interface{} `json:"data"`
}

// SuccessResponse returns success
func SuccessResponse() *DefaultResponse {
	dr := &DefaultResponse{StatusCode: statusSuccess}
	return dr
}
