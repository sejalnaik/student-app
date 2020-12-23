package usercontroller

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	model "github.com/sejalnaik/student-app/user/user-model"
	service "github.com/sejalnaik/student-app/user/user-service"
)

type userController struct {
	userService *service.UserService
}

func NewUserController(userService *service.UserService) *userController {
	return &userController{
		userService: userService,
	}
}

func (c *userController) CreateRoutes(r *mux.Router) {
	//login user
	r.HandleFunc("/login", c.GetUser).Methods("POST")
	//add student
	r.HandleFunc("/register", c.AddUser).Methods("POST")
}

func (c *userController) GetUser(w http.ResponseWriter, r *http.Request) {
	//create bucket
	user := &model.User{}

	//read user data from response body
	responseBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Get user : Could not read response body")
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	//convert json to struct type
	er := json.Unmarshal(responseBody, user)
	if er != nil {
		log.Println("Get user : Json unmarshall unsuccessful")
		log.Println(er)
		http.Error(w, er.Error(), http.StatusBadRequest)
		return
	}

	//calling service method to get user
	if err := c.userService.GetUser(user); err != nil {
		log.Println("Get user unsuccessful")
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
	} else {
		log.Println("Get user successful")
		w.Write([]byte(user.ID.String()))
	}
}

func (c *userController) AddUser(w http.ResponseWriter, r *http.Request) {
	//create bucket
	user := &model.User{}

	//read user data from response body
	responseBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Add user : Could not read response body")
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	//convert json to struct type
	er := json.Unmarshal(responseBody, user)
	if er != nil {
		log.Println("Add user : Json unmarshall unsuccessful")
		log.Println(er)
		http.Error(w, er.Error(), http.StatusBadRequest)
		return
	}

	//calling service method to add user and giving back id as string
	if err := c.userService.AddUser(user); err != nil {
		log.Println("Add user unsuccessful")
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
	} else {
		log.Println("Add user successful")
		w.Write([]byte(user.ID.String()))
	}
}
