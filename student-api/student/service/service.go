package service

import (
	"encoding/json"
	"errors"

	"github.com/jinzhu/gorm"
	"github.com/sejalnaik/student-app/repository"
	"github.com/sejalnaik/student-app/student/model"
	"github.com/sejalnaik/student-app/utility"
)

type StudentService struct {
	repository repository.Repository
	db         *gorm.DB
}

func NewStudentService(repository repository.Repository, db *gorm.DB) *StudentService {
	return &StudentService{
		repository: repository,
		db:         db,
	}
}

func (s *StudentService) GetAllStudents(students *[]model.Student) error {
	//create unit of work
	uow := repository.NewUnitOfWork(s.db, true)

	//call get repository method to get students
	if err := s.repository.Get(uow, students, nil); err != nil {
		return err
	} else {
		//to trim dob and dobtime
		utility.ConvertStudentsTimeToDate(students)
		return nil
	}
}

func (s *StudentService) GetStudent(student *model.Student, studentID string) error {
	//create unit of work
	uow := repository.NewUnitOfWork(s.db, true)

	//give query processor for where
	queryProcessors := []repository.QueryProcessor{}
	queryProcessors = append(queryProcessors, repository.Where(studentID))

	//call get repository method to get one student
	if err := s.repository.Get(uow, student, queryProcessors); err != nil {
		return err
	} else {
		//to trim dob and dobtime
		utility.ConvertStudentTimeToDate(student)
		return nil
	}
}

func (s *StudentService) AddStudent(student *model.Student) error {
	//create unit of work
	uow := repository.NewUnitOfWork(s.db, true)

	//check student validation
	if validErrs := student.Validate(); len(validErrs) > 0 {
		err := map[string]interface{}{"validationError": validErrs}
		errorJSONString, _ := json.Marshal(err)
		return errors.New(string(errorJSONString))
	}

	//convert empty frilds of student to null
	utility.AddStudentEmptyStringToNull(student)

	//call add repository method to add one student
	if err := s.repository.Add(uow, student); err != nil {
		uow.Complete()
		return err
	} else {
		uow.Commit()
		return nil
	}
}

func (s *StudentService) UpdateStudent(student *model.Student, studentID string) error {
	//create unit of work
	uow := repository.NewUnitOfWork(s.db, true)

	//check student validation
	if validErrs := student.Validate(); len(validErrs) > 0 {
		err := map[string]interface{}{"validationError": validErrs}
		errorJSONString, _ := json.Marshal(err)
		return errors.New(string(errorJSONString))
	}

	//give query processor for where
	queryProcessors := []repository.QueryProcessor{}
	queryProcessors = append(queryProcessors, repository.Where(studentID))

	//convert student struct to map
	studentMap := utility.ConvertStructStudentToMap(student)

	//call update repository method to update one student
	if err := s.repository.Update(uow, student, studentMap, queryProcessors); err != nil {
		uow.Complete()
		return err
	} else {
		uow.Commit()
		return nil
	}
}

func (s *StudentService) DeleteStudent(student *model.Student, studentID string) error {
	//create unit of work
	uow := repository.NewUnitOfWork(s.db, true)

	//give query processor for where
	queryProcessors := []repository.QueryProcessor{}
	queryProcessors = append(queryProcessors, repository.Where(studentID))

	//call delete repository method to delete one student
	if err := s.repository.Delete(uow, student, queryProcessors); err != nil {
		uow.Complete()
		return err
	} else {
		uow.Commit()
		return nil
	}
}
