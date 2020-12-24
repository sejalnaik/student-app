package usermodel

import (
	"net/url"
	"regexp"

	"github.com/dgrijalva/jwt-go"
	model "github.com/sejalnaik/student-app/student/student-model"
)

const JwtKey = "mysecretkey"

type User struct {
	model.Base
	Username string `json:"username"`
	Password string `json:"password"`
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
