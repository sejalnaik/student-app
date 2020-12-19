package service

import (
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
	uow := repository.NewUnitOfWork(s.db, true)
	if err := s.repository.Get(uow, students, nil); err != nil {
		return err
	} else {
		utility.ConvertStudentsTimeToDate(students)
		return nil
	}
}

func (s *StudentService) GetStudent(student *model.Student) error {
	uow := repository.NewUnitOfWork(s.db, true)
	queryProcessors := []repository.QueryProcessor{}
	queryProcessors = append(queryProcessors, repository.Where(student.ID))
	if err := s.repository.Get(uow, student, queryProcessors); err != nil {
		return err
	} else {
		utility.ConvertStudentTimeToDate(student)
		return nil
	}
}

func (s *StudentService) AddStudent(student *model.Student) error {
	uow := repository.NewUnitOfWork(s.db, true)
	if err := s.repository.Add(uow, student); err != nil {
		uow.Complete()
		return err
	} else {
		uow.Commit()
		return nil
	}
}

func (s *StudentService) UpdateStudent(student *model.Student) error {
	uow := repository.NewUnitOfWork(s.db, true)
	if err := s.repository.Update(uow, student); err != nil {
		uow.Complete()
		return err
	} else {
		uow.Commit()
		return nil
	}
}

func (s *StudentService) DeleteStudent(student *model.Student, studentID string) error {
	uow := repository.NewUnitOfWork(s.db, true)
	queryProcessors := []repository.QueryProcessor{}
	queryProcessors = append(queryProcessors, repository.Where(studentID))
	if err := s.repository.Delete(uow, student, queryProcessors); err != nil {
		uow.Complete()
		return err
	} else {
		uow.Commit()
		return nil
	}
}
