package studentcontroller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"

	"github.com/sejalnaik/student-app/model"
	service "github.com/sejalnaik/student-app/student/student-service"
)

type studentController struct {
	studentService *service.StudentService
}

func NewStudentController(studentService *service.StudentService) *studentController {
	return &studentController{
		studentService: studentService,
	}
}

func (c *studentController) CreateRoutes(r *mux.Router) {
	//create subrouter
	apiRoutes := r.PathPrefix("/api").Subrouter()
	//token check middleware
	apiRoutes.Use(tokenCheckMiddleware)
	//get students
	apiRoutes.HandleFunc("/students", c.GetAllStudents).Methods("GET")
	//get one student
	apiRoutes.HandleFunc("/students/{studentID}", c.GetStudent).Methods("GET")
	//add student
	apiRoutes.HandleFunc("/students", c.AddStudent).Methods("POST")
	//update student
	apiRoutes.HandleFunc("/students/{studentID}", c.UpdateStudent).Methods("PUT")
	//delete student
	apiRoutes.HandleFunc("/students/{studentID}", c.DeleteStudent).Methods("DELETE")
}

func tokenCheckMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("tokenCheckMiddleware called")

		//get token string from header
		tokenString := r.Header.Get("token")

		//if token not present send "not authorized to access"
		if tokenString == "" {
			log.Println("Token is not present, access not allowed")
			http.Error(w, "Token is not present, access not allowed", http.StatusUnauthorized)
			return
		}

		//trim inverted commas from the token
		tokenString = tokenString[1 : len(tokenString)-1]

		//parse the tokenstring to get the token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Error while parsing token")
			}
			return []byte(model.JwtKey), nil
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				//token is invalid
				log.Println("tokenCheckMiddleware : Token is invalid")
				log.Println(err)
				http.Error(w, err.Error(), http.StatusUnauthorized)
				return
			}
			//internal server error
			log.Println("tokenCheckMiddleware : internal server error")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if !token.Valid {
			//token is invalid
			log.Println("tokenCheckMiddleware : Token is invalid")
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func (c *studentController) GetAllStudents(w http.ResponseWriter, r *http.Request) {
	log.Println("Get students called")
	//create bucket
	students := []model.Student{}

	//calling service method to get all students
	if err := c.studentService.GetAllStudents(&students); err != nil {
		log.Println("Get students unsuccessful")
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	//converting struct to json type and sending back json
	if studentsJSON, err := json.Marshal(students); err != nil {
		log.Println("Get students : JSON marshall unsuccessful")
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		log.Println("Get students successful")
		w.Write(studentsJSON)
	}
}

func (c *studentController) GetStudent(w http.ResponseWriter, r *http.Request) {
	log.Println("Get student called")
	//create bucket
	student := model.Student{}

	//getting id from query param
	params := mux.Vars(r)
	studentID := (params["studentID"])

	//calling service method to get student
	if err := c.studentService.GetStudent(&student, studentID); err != nil {
		log.Println("Get student unsuccessful")
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	//converting struct to json type and sending back json
	if studentsJSON, err := json.Marshal(student); err != nil {
		log.Println("Get student : JSON marshall unsuccessful")
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		log.Println("Get student successful")
		w.Write(studentsJSON)
	}
}

func (c *studentController) AddStudent(w http.ResponseWriter, r *http.Request) {
	log.Println("Add student called")
	//create bucket
	student := &model.Student{}

	//read student data from response body
	responseBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Add student : Could not read response body")
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	//convert json to struct type
	er := json.Unmarshal(responseBody, student)
	if er != nil {
		log.Println("Add student : Json unmarshall unsuccessful")
		log.Println(er)
		http.Error(w, er.Error(), http.StatusBadRequest)
		return
	}

	//calling service method to add student and giving back id as string
	if err := c.studentService.AddStudent(student); err != nil {
		log.Println("Add student unsuccessful")
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		log.Println("Add student successful")
		w.Write([]byte(student.ID.String()))
	}
}

func (c *studentController) UpdateStudent(w http.ResponseWriter, r *http.Request) {
	log.Println("Update student called")
	//create bucket
	student := &model.Student{}

	//getting id from query param
	params := mux.Vars(r)
	studentID := (params["studentID"])

	//read student data from response body
	responseBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Update student : Could not read response body")
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	//convert json to struct type
	er := json.Unmarshal(responseBody, student)
	if er != nil {
		log.Println("Update student : Json unmarshall unsuccessful")
		log.Println(er)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//calling service method to update student and giving back id as string
	if err := c.studentService.UpdateStudent(student, studentID); err != nil {
		log.Println("Update student unsuccessful")
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		log.Println("Update student successful")
		//w.Write([]byte(student.ID.String()))
	}
}

func (c *studentController) DeleteStudent(w http.ResponseWriter, r *http.Request) {
	log.Println("Delete student called")
	//create bucket
	student := &model.Student{}

	//getting id from query param
	params := mux.Vars(r)
	studentID := (params["studentID"])

	//calling service method to delete student and giving back id as string
	if err := c.studentService.DeleteStudent(student, studentID); err != nil {
		log.Println("Delete student unsuccessful")
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		log.Println("Delete student successful")
		//w.Write([]byte(studentID))
	}
}
