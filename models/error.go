package models

type FieldError struct {
	ErrorField   string `json:"error_field"`
	ErrorMessage string `json:"error_message"`
}
