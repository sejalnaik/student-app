package studentservice

import (
	"encoding/json"
	"errors"
	"log"

	"github.com/jinzhu/gorm"
	"github.com/sejalnaik/student-app/model"
	"github.com/sejalnaik/student-app/repository"
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
	queryProcessors = append(queryProcessors, repository.Where("id=?", studentID))

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

	//convert empty fields of student to null
	utility.AddStudentEmptyStringToNull(student)

	log.Println("**********************************************************")
	log.Println(student)

	//call add repository method to add one student
	if err := s.repository.Add(uow, student); err != nil {
		uow.Complete()
		return err
	}
	uow.Commit()
	return nil
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
	queryProcessors = append(queryProcessors, repository.Where("id=?", studentID))

	//convert student struct to map
	studentMap := utility.ConvertStructStudentToMap(student)

	//call update repository method to update one student
	if err := s.repository.Update(uow, student, studentMap, queryProcessors); err != nil {
		uow.Complete()
		return err
	}
	uow.Commit()
	return nil
}

func (s *StudentService) DeleteStudent(student *model.Student, studentID string) error {
	//create unit of work
	uow := repository.NewUnitOfWork(s.db, true)

	//check if student has any book issues
	queryProcessors := []repository.QueryProcessor{}
	queryProcessors = append(queryProcessors, repository.Where("student_id=?", studentID))
	if err := s.repository.Get(uow, &model.BookIssue{}, queryProcessors); err == nil {
		return errors.New("Student cannot be deleted because it has book issues")
	} else if err.Error() == "record not found" {
		goto canBeDeleted
	} else {
		return err
	}

	//give query processor for where
canBeDeleted:
	queryProcessors = []repository.QueryProcessor{}
	queryProcessors = append(queryProcessors, repository.Where("id=?", studentID))

	//call delete repository method to delete one student
	if err := s.repository.Delete(uow, student, queryProcessors); err != nil {
		uow.Complete()
		return err
	}
	uow.Commit()
	return nil
}

func (s *StudentService) SumOfAgeAndRollNo(sum *model.Result) error {
	//create unit of work
	uow := repository.NewUnitOfWork(s.db, true)

	//set condition for select query
	condition := "sum(roll_no+age) as total"

	//give query processor for where
	queryProcessors := []repository.QueryProcessor{}
	queryProcessors = append(queryProcessors, repository.Select(condition))

	//call select repository method to calculate sum of age and rollno of all students
	if err := s.repository.Scan(uow, &model.Student{}, sum, queryProcessors); err != nil {
		return err
	} else {
		return nil
	}
}

func (s *StudentService) DiffOfAgeAndRollNo(diff *model.Result) error {
	//create unit of work
	uow := repository.NewUnitOfWork(s.db, true)

	//set condition for select query
	condition := "sum(roll_no-age) as total"

	//give query processor for where
	queryProcessors := []repository.QueryProcessor{}
	queryProcessors = append(queryProcessors, repository.Select(condition))

	//call select repository method to calculate diff of age and rollno of all students
	if err := s.repository.Scan(uow, &model.Student{}, diff, queryProcessors); err != nil {
		return err
	} else {
		return nil
	}
}

func (s *StudentService) DiffOfAgeAndRecordCount(diff *model.Result) error {
	//create unit of work
	uow := repository.NewUnitOfWork(s.db, true)

	//set condition for select query
	condition := "sum(age) - count(*) as total"

	//give query processor for where
	queryProcessors := []repository.QueryProcessor{}
	queryProcessors = append(queryProcessors, repository.Select(condition))

	//call select repository method to calculate diff of age and rollno of all students
	if err := s.repository.Scan(uow, &model.Student{}, diff, queryProcessors); err != nil {
		return err
	} else {
		return nil
	}
}

func (s *StudentService) TotalPenalty(sum *model.TotalPenalty, studentID string) error {
	//create unit of work
	uow := repository.NewUnitOfWork(s.db, true)

	//set condition for select query
	condition := "sum(penalty) as total"

	//give query processor for where
	queryProcessors := []repository.QueryProcessor{}
	queryProcessors = append(queryProcessors, repository.Where("student_id=?", studentID))
	queryProcessors = append(queryProcessors, repository.Select(condition))

	//call select repository method to calculate total penalty
	if err := s.repository.Scan(uow, &model.BookIssue{}, sum, queryProcessors); err != nil {
		return err
	} else {
		return nil
	}
}

func (s *StudentService) SearchStudents(paramsMap map[string][]string, students *[]model.Student) error {
	//create unit of work
	uow := repository.NewUnitOfWork(s.db, true)

	//create where condition
	condition := ""
	paramsMapLength := len(paramsMap)
	for key, value := range paramsMap {
		if key == "from" {
			condition = condition + "dob >= " + "'" + value[0] + "'" + " "

		} else if key == "to" {
			condition = condition + "dob <= " + "'" + value[0] + "'" + " "
		} else {
			condition = condition + key + " like " + "'%" + value[0] + "%' "
		}
		paramsMapLength = paramsMapLength - 1
		if paramsMapLength == 0 {
			break
		}
		condition = condition + " and "
	}

	//give query processor for where
	queryProcessors := []repository.QueryProcessor{}
	queryProcessors = append(queryProcessors, repository.Where(condition))

	//call get repository method to get students
	if err := s.repository.Get(uow, students, queryProcessors); err != nil {
		return err
	} else {
		//to trim dob and dobtime
		utility.ConvertStudentsTimeToDate(students)
		return nil
	}
}
