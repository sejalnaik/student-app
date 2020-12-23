package userModel

import "github.com/sejalnaik/student-app/student/model"

type User struct {
	model.Base
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`          
}
