package models

type Authentication struct {
	AuthID    int    `json:"auth_id" gorm:"primaryKey"`
	PartnerID int    `json:"partner_id"`
	ExpiresAt string `json:"expires_at"`
	AuthType  string `json:"auth_type"`
}
