package models

type TransformationRequest struct {
	InputData map[string]interface{} `json:"input_data" binding:"required"`
}
