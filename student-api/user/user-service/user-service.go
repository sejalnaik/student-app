package userservice

import (
	"encoding/json"
	"errors"

	"github.com/jinzhu/gorm"
	"github.com/sejalnaik/student-app/model"
	"github.com/sejalnaik/student-app/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repository repository.Repository
	db         *gorm.DB
}

func NewUserService(repository repository.Repository, db *gorm.DB) *UserService {
	return &UserService{
		repository: repository,
		db:         db,
	}
}

func (s *UserService) GetUser(userToBeChecked *model.User) error {
	//create unit of work
	uow := repository.NewUnitOfWork(s.db, true)

	//create a new instance to store user to be got from database
	userFromDatabase := &model.User{}

	//give query processor for where
	queryProcessors := []repository.QueryProcessor{}
	queryProcessors = append(queryProcessors, repository.Where("username=?", userToBeChecked.Username))

	//call get repository method to get one user
	if err := s.repository.Get(uow, userFromDatabase, queryProcessors); err != nil {
		return err
	}

	//decrypt and check password
	if err := bcrypt.CompareHashAndPassword([]byte(userFromDatabase.Password), []byte(userToBeChecked.Password)); err != nil {
		return errors.New("Unauthorized user")
	}
	return nil
}

func (s *UserService) AddUser(user *model.User) error {
	//create unit of work
	uow := repository.NewUnitOfWork(s.db, true)

	//check user validation
	if validErrs := user.Validate(); len(validErrs) > 0 {
		err := map[string]interface{}{"validationError": validErrs}
		errorJSONString, _ := json.Marshal(err)
		return errors.New(string(errorJSONString))
	}

	//encrypt password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 8)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	//call add repository method to add one user
	if err := s.repository.Add(uow, user); err != nil {
		uow.Complete()
		return err
	}
	uow.Commit()
	return nil
}

func (s *UserService) CheckIfUsernameAvailable(userToBeChecked *model.User) error {
	//create unit of work
	uow := repository.NewUnitOfWork(s.db, true)

	//create a new instance to store user to be got from database
	userFromDatabase := &model.User{}

	//give query processor for where
	queryProcessors := []repository.QueryProcessor{}
	queryProcessors = append(queryProcessors, repository.Where("username=?", userToBeChecked.Username))

	//call get repository method to get one user
	if err := s.repository.Get(uow, userFromDatabase, queryProcessors); err != nil {
		return err
	}
	return nil
}
