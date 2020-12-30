// Student API
//
// Example Swagger spec.
//
//	Schemes: http
//	BasePath: /
//	Version: 1.0.0
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//     Security:
//     - api_key:
//
//     SecurityDefinitions:
//     api_key:
//          type: apiKey
//          name: token
//          in: header
// swagger:meta
package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/sejalnaik/student-app/model"
	"github.com/sejalnaik/student-app/repository"
	studentcontroller "github.com/sejalnaik/student-app/student/student-controller"

	studentservice "github.com/sejalnaik/student-app/student/student-service"

	usercontroller "github.com/sejalnaik/student-app/user/user-controller"

	userservice "github.com/sejalnaik/student-app/user/user-service"
)

func main() {

	//create db
	db, err := gorm.Open("mysql", "root:root@tcp(localhost:4040)/student_app?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()
	if err != nil {
		fmt.Print(err)
	}
	log.Println("DB connection established")
	db.AutoMigrate(&model.Student{}, &model.User{})

	//create router
	r := mux.NewRouter()
	headers := handlers.AllowedHeaders([]string{"Content-Type", "token"})
	methods := handlers.AllowedMethods([]string{"POST", "GET", "PUT", "DELETE"})
	origin := handlers.AllowedOrigins([]string{"*"})

	server := &http.Server{
		Handler:      handlers.CORS(headers, methods, origin)(r),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		Addr:         ":8080",
	}

	log.Println("Server connected")

	//create repository
	repository := repository.NewRepository()

	//create student service
	studentService := studentservice.NewStudentService(repository, db)

	//create student controller
	studentController := studentcontroller.NewStudentController(studentService)

	//create student routes
	studentController.CreateRoutes(r)

	//create user service
	userService := userservice.NewUserService(repository, db)

	//create user controller
	userController := usercontroller.NewUserController(userService)

	//create user routes
	userController.CreateRoutes(r)

	//listen to port 8080
	log.Fatal(server.ListenAndServe())
}
