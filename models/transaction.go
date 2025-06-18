package models

type Transaction struct {
	TransactionID   int    `json:"transaction_id" gorm:"primaryKey"`
	PartnerID       int    `json:"partner_id"`
	EndpointID      int    `json:"endpoint_id"`
	TransactionType string `json:"transaction_type"`
	Status          string `json:"status"`
	RequestPayload  string `json:"request_payload"`
	ResponsePayload string `json:"response_payload"`
	ErrorMessage    string `json:"error_message"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
}
