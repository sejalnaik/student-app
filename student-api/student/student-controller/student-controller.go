package studentcontroller

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	model "github.com/sejalnaik/student-app/student/student-model"
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
	//get students
	r.HandleFunc("/students", c.GetAllStudents).Methods("GET")
	//get one student
	r.HandleFunc("/students/{studentID}", c.GetStudent).Methods("GET")
	//add student
	r.HandleFunc("/students", c.AddStudent).Methods("POST")
	//update student
	r.HandleFunc("/students/{studentID}", c.UpdateStudent).Methods("PUT")
	//delete student
	r.HandleFunc("/students/{studentID}", c.DeleteStudent).Methods("DELETE")
}

func (c *studentController) GetAllStudents(w http.ResponseWriter, r *http.Request) {
	//create bucket
	students := []model.Student{}

	//calling service method to get all students
	if err := c.studentService.GetAllStudents(&students); err != nil {
		log.Println("Get students unsuccessful")
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	//converting struct to json type and sending back json
	if studentsJSON, err := json.Marshal(students); err != nil {
		log.Println("Get students : JSON marshall unsuccessful")
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
	} else {
		log.Println("Get students successful")
		w.Write(studentsJSON)
	}
}

func (c *studentController) GetStudent(w http.ResponseWriter, r *http.Request) {
	//create bucket
	student := model.Student{}

	//getting id from query param
	params := mux.Vars(r)
	studentID := (params["studentID"])

	//calling service method to get student
	if err := c.studentService.GetStudent(&student, studentID); err != nil {
		log.Println("Get student unsuccessful")
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	//converting struct to json type and sending back json
	if studentsJSON, err := json.Marshal(student); err != nil {
		log.Println("Get student : JSON marshall unsuccessful")
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
	} else {
		log.Println("Get student successful")
		w.Write(studentsJSON)
	}
}

func (c *studentController) AddStudent(w http.ResponseWriter, r *http.Request) {
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
		http.Error(w, err.Error(), http.StatusBadRequest)
	} else {
		log.Println("Add student successful")
		w.Write([]byte(student.ID.String()))
	}
}

func (c *studentController) UpdateStudent(w http.ResponseWriter, r *http.Request) {
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
		http.Error(w, err.Error(), http.StatusBadRequest)
	} else {
		log.Println("Update student successful")
		w.Write([]byte(student.ID.String()))
	}
}

func (c *studentController) DeleteStudent(w http.ResponseWriter, r *http.Request) {
	//create bucket
	student := &model.Student{}

	//getting id from query param
	params := mux.Vars(r)
	studentID := (params["studentID"])

	//calling service method to delete student and giving back id as string
	if err := c.studentService.DeleteStudent(student, studentID); err != nil {
		log.Println("Delete student unsuccessful")
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
	} else {
		log.Println("Delete student successful")
		w.Write([]byte(studentID))
	}
}
