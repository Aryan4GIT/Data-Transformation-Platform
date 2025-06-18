package models

type TransactionLog struct {
	ClientIP    string `json:"client_ip"`
	QueryParams string `json:"query_params"`
	PartnerID   int    `json:"partner_id"`
	Error       string `json:"error"`
	Status      string `json:"status"`
}
