package models

import (
	"time"
)

// BaseModel TODO: for ID - primary_key;type:uuid;default:uuid_generate_v4()
type BaseModel struct {
	ID        string    `json:"-"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	DeletedAt time.Time `json:"-"`
}
