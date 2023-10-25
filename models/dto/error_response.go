package dto

type ErrorResponse struct {
	Code    int         `json:"code,omitempty"`
	Status  string      `json:"status,omitempty"`
	Message interface{} `json:"message,omitempty"`
}
