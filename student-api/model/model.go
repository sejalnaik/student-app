package model

import (
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type Student struct {
	Base
	RollNo  int    `json:"rollNo"`
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Email   string `json:"email"`
	IsMale  bool   `json:"isMale"`
	DOB     string `gorm:"type:date" json:"dob"`
	DOBTIME string `gorm:"type:datetime" json:"dobTime"`
}

type Base struct {
	ID        uuid.UUID  `gorm:"type:varchar(36);primary_key" json:"id"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `sql:"index" json:"-"`
}

func (model *Base) BeforeCreate(scope *gorm.Scope) {
	uuid := uuid.NewV4()
	scope.SetColumn("ID", uuid)
}
