package models

type Endpoint struct {
	EndpointID  int    `json:"endpoint_id" gorm:"primaryKey"`
	PartnerID   int    `json:"partner_id"`
	EndpointURL string `json:"endpoint_url"`
	HTTPMethod  string `json:"http_method"`
	Headers     string `json:"headers"`
	IsActive    bool   `json:"is_active"`
	CreatedAt   string `json:"created_at"`
}
