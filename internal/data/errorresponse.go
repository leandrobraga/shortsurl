package data

type ErrorResponse struct {
	Field string      `json:"field"`
	Tag   string      `json:"tag"`
	Value interface{} `json:"value"`
}
