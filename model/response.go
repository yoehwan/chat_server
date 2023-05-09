package model

import "encoding/json"

type Response struct {
	Status  int
	Message string         `json:"Message,omitempty"`
	Data    map[string]any `json:"Data,omitempty"`
}

func SuccessResponse(data any) *Response {
	var jsonData map[string]interface{}
	bytes, _ := json.Marshal(data)
	json.Unmarshal(bytes, &jsonData)
	return &Response{
		Status: 0,
		Data:   jsonData,
	}
}

func FailResponse(message string) *Response {
	return &Response{
		Status:  1,
		Message: message,
	}
}
