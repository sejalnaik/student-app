package model

import (
	"net/url"
	"regexp"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type Student struct {
	Base
	RollNo      *int    `gorm:"type:int" json:"rollNo"`
	Name        string  `gorm:"type:varchar(100)" json:"name"`
	Age         *int    `gorm:"type:int" json:"age"`
	Email       string  `gorm:"type:varchar(150)" json:"email"`
	IsMale      *bool   `gorm:"type:tinyint" json:"isMale"`
	DOB         *string `gorm:"type:date" json:"dob"`
	DOBTIME     *string `gorm:"type:datetime" json:"dobTime"`
	PhoneNumber *string `gorm:"type:varchar(12)" json:"phoneNumber"`
}

func (model *Base) BeforeCreate(scope *gorm.Scope) {
	uuid := uuid.NewV4()
	scope.SetColumn("ID", uuid)
}

func (s *Student) Validate() url.Values {
	regexpEmail := regexp.MustCompile("^[a-zA-Z0-9._%+-]+@[a-z0-9.-]+\\.[a-z]{2,4}$")
	regexpName := regexp.MustCompile("^[a-zA-Z_ ]+$")
	errs := url.Values{}

	//name is required
	if s.Name == "" {
		errs.Add("name", "Name is required")
	}

	//name must contain alphabets only
	if !regexpName.MatchString(s.Name) {
		errs.Add("name", "Name should only have alphabets")
	}

	//email is required
	if s.Email == "" {
		errs.Add("email", "Email is required")
	}

	//email must be valid
	if !regexpEmail.MatchString(s.Email) {
		errs.Add("email", "Email field should be a valid email address")
	}

	//phone number must be between 10 and 12
	if s.PhoneNumber != nil {
		if len(*s.PhoneNumber) < 10 && len(*s.PhoneNumber) > 12 {
			errs.Add("phoneNumber", "Phone number should be between 10 and 12")
		}
	}

	return errs
}
