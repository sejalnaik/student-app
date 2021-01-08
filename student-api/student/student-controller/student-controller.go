package studentcontroller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"

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
	//create route for get students
	r.HandleFunc("/students", c.GetAllStudents).Methods("GET")

	//create route for get student
	r.HandleFunc("/students/{studentID}", c.GetStudent).Methods("GET")

	//create route for searching students
	r.HandleFunc("/students-search", c.SearchStudents).Methods("GET")

	//create route for add student
	//r.HandleFunc("/students", c.AddStudent).Methods("POST")
	nAddStudent := negroni.New()
	nAddStudent.Use(negroni.HandlerFunc(tokenCheckMiddleware))
	nAddStudent.UseHandlerFunc(c.AddStudent)
	r.Handle("/students", nAddStudent).Methods("POST")

	//create route for update student
	//r.HandleFunc("/students/{studentID}", c.UpdateStudent).Methods("PUT")
	nUpdateStudent := negroni.New()
	nUpdateStudent.Use(negroni.HandlerFunc(tokenCheckMiddleware))
	nUpdateStudent.UseHandlerFunc(c.UpdateStudent)
	r.Handle("/students/{studentID}", nUpdateStudent).Methods("PUT")

	//create route for delete student
	nDeleteStudent := negroni.New()
	nDeleteStudent.Use(negroni.HandlerFunc(tokenCheckMiddleware))
	nDeleteStudent.UseHandlerFunc(c.DeleteStudent)
	r.Handle("/students/{studentID}", nDeleteStudent).Methods("DELETE")

	//create route for total penalty
	r.HandleFunc("/student-penalty/{studentID}", c.TotalPenalty).Methods("GET")

	//create route for sum of age and rollno
	r.HandleFunc("/sum", c.SumOfAgeAndRollNo).Methods("GET")

	//create route for diff of age and rollno
	r.HandleFunc("/diff", c.DiffOfAgeAndRollNo).Methods("GET")

	//create route for diff of age and record count
	r.HandleFunc("/diff-age-record-count", c.DiffOfAgeAndRecordCount).Methods("GET")
}

func tokenCheckMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
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
	//tokenString = tokenString[1 : len(tokenString)-1]

	//create empty claims
	claims := &model.Claims{}

	//parse the tokenstring to get the token
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
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
		log.Println("tokenCheckMiddleware : internal server error while parsing")
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	/*
		//add 30 minutes to expiration time
		expirationTime := time.Now().Add(30 * time.Minute)
		claims.ExpiresAt = expirationTime.Unix()

		//create new token
		newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		newTokenString, err := newToken.SignedString([]byte(model.JwtKey))
		if err != nil {
			log.Println("tokenCheckMiddleware : internal server error while creating")
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		//pass it to the next handler
		context.Set(r, "token", newTokenString)
	*/
	if !token.Valid {
		//token is invalid
		log.Println("tokenCheckMiddleware : Token is invalid")
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	next.ServeHTTP(w, r)
}

func (c *studentController) GetAllStudents(w http.ResponseWriter, r *http.Request) {
	log.Println("Get students called")
	/*
		//get token from context set by middleware
		tokenString := context.Get(r, "token")

		//set header only if token is set in middleware
		if tokenString != nil {
			//set header of the request to take the token
			w.Header().Add("Access-Control-Expose-Headers", "token")
			w.Header().Set("token", tokenString.(string))
		}
	*/
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
	/*
		//get token from context set by middleware
		tokenString := context.Get(r, "token")

		//set header only if token is set in middleware
		if tokenString != nil {
			//set header of the request to take the token
			w.Header().Add("Access-Control-Expose-Headers", "token")
			w.Header().Set("token", tokenString.(string))
		}
	*/
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
	/*
		//get token from context set by middleware
		tokenString := context.Get(r, "token")

		//set header only if token is set in middleware
		if tokenString != nil {
			//set header of the request to take the token
			w.Header().Add("Access-Control-Expose-Headers", "token")
			w.Header().Set("token", tokenString.(string))
		}
	*/
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
	/*
		//get token from context set by middleware
		tokenString := context.Get(r, "token")

		//set header only if token is set in middleware
		if tokenString != nil {
			//set header of the request to take the token
			w.Header().Add("Access-Control-Expose-Headers", "token")
			w.Header().Set("token", tokenString.(string))
		}
	*/
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
	/*
		//get token from context set by middleware
		tokenString := context.Get(r, "token")

		//set header only if token is set in middleware
		if tokenString != nil {
			//set header of the request to take the token
			w.Header().Add("Access-Control-Expose-Headers", "token")
			w.Header().Set("token", tokenString.(string))
		}
	*/
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

func (c *studentController) SumOfAgeAndRollNo(w http.ResponseWriter, r *http.Request) {
	log.Println("SumOfAgeAndRollNo called")

	//decalre a sum variable for sum
	var sum model.Result

	if err := c.studentService.SumOfAgeAndRollNo(&sum); err != nil {
		log.Println("Sum unsuccessful")
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	//converting struct to json type and sending back json
	if sumJSON, err := json.Marshal(&sum); err != nil {
		log.Println("Sum : JSON marshall unsuccessful")
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		log.Println("sum successful")
		w.Write(sumJSON)
	}
}

func (c *studentController) DiffOfAgeAndRollNo(w http.ResponseWriter, r *http.Request) {
	log.Println("DiffOfAgeAndRollNo called")

	//decalre a sum variable for sum
	var diff model.Result

	if err := c.studentService.DiffOfAgeAndRollNo(&diff); err != nil {
		log.Println("Diff unsuccessful")
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	//converting struct to json type and sending back json
	if sumJSON, err := json.Marshal(&diff); err != nil {
		log.Println("Difff : JSON marshall unsuccessful")
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		log.Println("Diff successful")
		w.Write(sumJSON)
	}
}

func (c *studentController) DiffOfAgeAndRecordCount(w http.ResponseWriter, r *http.Request) {
	log.Println("DiffOfAgeAndRecordCount called")

	//decalre a sum variable for sum
	var diff model.Result

	if err := c.studentService.DiffOfAgeAndRecordCount(&diff); err != nil {
		log.Println("Diff unsuccessful")
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	//converting struct to json type and sending back json
	if sumJSON, err := json.Marshal(&diff); err != nil {
		log.Println("Difff : JSON marshall unsuccessful")
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		log.Println("Diff successful")
		w.Write(sumJSON)
	}
}

func (c *studentController) TotalPenalty(w http.ResponseWriter, r *http.Request) {
	log.Println("TotalPenalty called")

	//decalre a sum variable for total penalty
	var sum model.TotalPenalty

	//getting id from query param
	params := mux.Vars(r)
	studentID := (params["studentID"])

	//calling service method for total penalty
	if err := c.studentService.TotalPenalty(&sum, studentID); err != nil {
		log.Println("Total penalty unsuccessful")
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	//converting struct to json type and sending back json
	if sumJSON, err := json.Marshal(&sum); err != nil {
		log.Println("Total penalty : JSON marshall unsuccessful")
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		log.Println("Total penalty unsuccessful")
		w.Write(sumJSON)
	}
}

func (c *studentController) SearchStudents(w http.ResponseWriter, r *http.Request) {
	log.Println("Search students called")

	//create map for query params
	paramsMap := r.URL.Query()

	//if no query params present then call get all students controller method
	if len(paramsMap) == 0 {
		c.GetAllStudents(w, r)
		return
	}

	//create bucket
	students := []model.Student{}

	//call service method for search students
	if err := c.studentService.SearchStudents(paramsMap, &students); err != nil {
		log.Println("Search students unsuccessful")
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	//converting struct to json type and sending back json
	if studentsJSON, err := json.Marshal(students); err != nil {
		log.Println("Search students : JSON marshall unsuccessful")
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		log.Println("Search students successful")
		w.Write(studentsJSON)
	}
}
