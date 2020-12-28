package model

import (
	"net/url"
	"regexp"
	"time"

	"github.com/dgrijalva/jwt-go"
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

type Base struct {
	ID        uuid.UUID  `gorm:"type:varchar(36);primary_key" json:"id"`
	CreatedAt time.Time  `gorm:"type:datetime" json:"-"`
	UpdatedAt time.Time  `gorm:"type:datetime" json:"-"`
	DeletedAt *time.Time `gorm:"type:datetime" sql:"index" json:"-"`
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

const JwtKey = "mysecretkey"

type User struct {
	Base
	Username string `gorm:"type:varchar(100)" json:"username"`
	Password string `gorm:"type:varchar(100)" json:"password"`
	//Email    string `json:"email"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func (s *User) Validate() url.Values {
	//regexpEmail := regexp.MustCompile("^[a-zA-Z0-9._%+-]+@[a-z0-9.-]+\\.[a-z]{2,4}$")
	regexpName := regexp.MustCompile("^[A-Za-z]+$")
	errs := url.Values{}

	//username is required
	if s.Username == "" {
		errs.Add("username", "Username is required")
	}

	//username must contain alphabets only
	if !regexpName.MatchString(s.Username) {
		errs.Add("username", "Username should only have alphabets")
	}

	//Password is required
	if s.Password == "" {
		errs.Add("password", "Password is required")
	}

	//email is required
	/*if s.Email == "" {
		errs.Add("email", "Email is required")
	}

	//email must be valid
	if !regexpEmail.MatchString(s.Email) {
		errs.Add("email", "Email field should be a valid email address")
	}*/
	return errs
}
