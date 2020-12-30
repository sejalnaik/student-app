package usercontroller

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/sejalnaik/student-app/model"
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
	r.HandleFunc("/login", c.Login).Methods("POST")
	//register user
	r.HandleFunc("/register", c.Register).Methods("POST")
}

func (c *userController) Login(w http.ResponseWriter, r *http.Request) {
	log.Println("Login user called")
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
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		log.Println("Get user successful")

		//create the token
		if tokenString, err := createToken(user); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		} else {
			log.Println("Token created successfully")
			log.Println("Token sent successfully")
			w.Write([]byte(tokenString))
		}
	}
}

func (c *userController) Register(w http.ResponseWriter, r *http.Request) {
	log.Println("Register user called")
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

	//check if the username already exists
	if err := c.userService.CheckIfUsernameAvailable(user); err != nil {
		log.Println("Username is unique")
		//calling service method to add user and giving back id as string
		if err := c.userService.AddUser(user); err != nil {
			log.Println("Add user unsuccessful")
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		} else {
			log.Println("Add user successful")
			w.Write([]byte(user.ID.String()))
			return
		}
	}

	//give back error if username exists
	log.Println("Username exists")
	http.Error(w, "Username exists", http.StatusBadRequest)
}

func createToken(user *model.User) (string, error) {
	//set expiration time for cookie
	expirationTime := time.Now().Add(30 * time.Minute)

	//create a claim for the token
	claims := &model.Claims{
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	//create the token using hash algo and claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//create the JWT string
	tokenString, err := token.SignedString([]byte(model.JwtKey))
	if err != nil {
		log.Println("Create token: could not create token")
		log.Println(err)
		return "", err
	}
	return tokenString, nil
}
