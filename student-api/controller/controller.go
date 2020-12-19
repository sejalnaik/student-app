package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sejalnaik/student-app/model"
	"github.com/sejalnaik/student-app/service"
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
	r.HandleFunc("/students", c.GetAllStudents).Methods("GET")
	r.HandleFunc("/students/{studentID}", c.GetStudent).Methods("GET")
	r.HandleFunc("/students", c.AddStudent).Methods("POST")
	r.HandleFunc("/students/{studentID}", c.UpdateStudent).Methods("PUT")
	r.HandleFunc("/students/{studentID}", c.DeleteStudent).Methods("DELETE")
}

func (c *studentController) GetAllStudents(w http.ResponseWriter, r *http.Request) {
	//create bucket
	students := []model.Student{}

	//calling service method to get all students
	if err := c.studentService.GetAllStudents(&students); err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	//converting struct to json type and sending back json
	if studentsJSON, err := json.Marshal(students); err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
	} else {
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
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	//converting struct to json type and sending back json
	if studentsJSON, err := json.Marshal(student); err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
	} else {
		w.Write(studentsJSON)
	}
}

func (c *studentController) AddStudent(w http.ResponseWriter, r *http.Request) {
	//create bucket
	student := &model.Student{}

	//read student data from response body
	responseBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	//convert json to struct type
	er := json.Unmarshal(responseBody, student)
	if er != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//calling service method to add student and giving back id as string
	if err := c.studentService.AddStudent(student); err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
	} else {
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
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	er := json.Unmarshal(responseBody, student)
	if er != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//calling service method to update student and giving back id as string
	if err := c.studentService.UpdateStudent(student, studentID); err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
	} else {
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
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
	} else {
		w.Write([]byte(studentID))
	}
}
