package model

import (
	"net/url"
	"regexp"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

//const regexpEmail = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
//const regexpName = regexp.MustCompile("^[A-Za-z]+$")

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

type SpecialDate struct {
	DOBTIME time.Time `gorm:"type:datetime" json:"dobTime"`
}

func (sd *SpecialDate) UnmarshalJSON(input []byte) error {

	strInput := string(input)
	strInput = strings.Trim(strInput, `"`)
	newTime, err := time.Parse("2006-01-02T15:04:05", strInput)
	if err != nil {
		return err
	}

	sd.DOBTIME = newTime
	return nil
}

func (model *Base) BeforeCreate(scope *gorm.Scope) {
	uuid := uuid.NewV4()
	scope.SetColumn("ID", uuid)
}

func (s *Student) Validate() url.Values {
	regexpEmail := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	regexpName := regexp.MustCompile("^[A-Za-z]+$")
	errs := url.Values{}

	if s.RollNo == 0 {

		errs.Add("rollNo", "Roll No is required")
	}
	if s.Name == "" {
		errs.Add("name", "Name is required")
	}
	if !regexpName.MatchString(s.Name) {
		errs.Add("name", "Name should only have alphabets")
	}
	if s.Email == "" {
		errs.Add("email", "Email is required")
	}
	if !regexpEmail.MatchString(s.Email) {
		errs.Add("email", "Email field should be a valid email address")
	}
	if s.Age == 0 {
		errs.Add("age", "Age is required")
	}
	if s.DOB == "" {
		errs.Add("dob", "DOB No is required")
	}
	if s.DOBTIME == "" {
		errs.Add("dobTime", "DOB/Time No is required")
	}
	return errs
}
