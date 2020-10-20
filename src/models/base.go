package models

import "time"

// Base ... Represents the basic fields for an entity
type Base struct {
	ID         string    `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	CreatedAt  time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
}
