package model

import (
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Base model
type Base struct {
	ID        *uuid.UUID       `json:"id,omitempty" gorm:"primaryKey;unique;type:varchar(36);not null" format:"uuid"`        // model ID
	CreatedAt *strfmt.DateTime `json:"created_at,omitempty" gorm:"type:timestamptz" format:"date-time" swaggertype:"string"` // created at automatically inserted on post
	UpdatedAt *strfmt.DateTime `json:"updated_at,omitempty" gorm:"type:timestamptz" format:"date-time" swaggertype:"string"` // updated at automatically changed on put or add on post
	DeletedAt gorm.DeletedAt   `json:"-" gorm:"index" swaggerignore:"true"`
}

