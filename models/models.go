package models

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"
)

type Client struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"unique;not null" json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type MappingRule struct {
	ID              uint           `gorm:"primaryKey" json:"id"`
	ClientID        uint           `gorm:"not null" json:"client_id"`
	Client          Client         `gorm:"foreignKey:ClientID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	SourcePath      JSONStringList `gorm:"type:jsonb;not null" json:"source_path"`
	DestinationPath JSONStringList `gorm:"type:jsonb;not null" json:"destination_path"`
	TransformType   string         `gorm:"not null" json:"transform_type"`
	TransformLogic  string         `gorm:"type:text" json:"transform_logic"` // New: dynamic logic
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
}

type JSONStringList []string

func (j *JSONStringList) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to unmarshal JSONB value")
	}
	return json.Unmarshal(bytes, j)
}

func (j JSONStringList) Value() (driver.Value, error) {
	return json.Marshal(j)
}
