package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type Base struct {
	ID        uuid.UUID  `gorm:"type:varchar(36);primary_key" json:"id"`
	CreatedAt time.Time  `gorm:"type:datetime" json:"-"`
	UpdatedAt time.Time  `gorm:"type:datetime" json:"-"`
	DeletedAt *time.Time `gorm:"type:datetime" sql:"index" json:"-"`
}
