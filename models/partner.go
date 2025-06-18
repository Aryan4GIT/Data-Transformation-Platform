package models

type Partner struct {
	PartnerID   int    `json:"partner_id" gorm:"primaryKey"`
	PartnerName string `json:"partner_name"`
	PartnerType string `json:"partner_type"`
	IsActive    bool   `json:"is_active"`
}
