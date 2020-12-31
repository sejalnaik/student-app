package bookservice

import (
	"encoding/json"
	"errors"

	"github.com/jinzhu/gorm"
	"github.com/sejalnaik/student-app/model"
	"github.com/sejalnaik/student-app/repository"
)

type BookService struct {
	repository repository.Repository
	db         *gorm.DB
}

func NewBookService(repository repository.Repository, db *gorm.DB) *BookService {
	return &BookService{
		repository: repository,
		db:         db,
	}
}

func (s *BookService) GetAllBooks(books *[]model.Book) error {
	//create unit of work
	uow := repository.NewUnitOfWork(s.db, true)

	//call get repository method to get books
	if err := s.repository.Get(uow, books, nil); err != nil {
		return err
	} else {
		return nil
	}
}

func (s *BookService) AddBook(book *model.Book) error {
	//create unit of work
	uow := repository.NewUnitOfWork(s.db, true)

	//check book validation
	if validErrs := book.Validate(); len(validErrs) > 0 {
		err := map[string]interface{}{"validationError": validErrs}
		errorJSONString, _ := json.Marshal(err)
		return errors.New(string(errorJSONString))
	}

	//call add repository method to add one book
	if err := s.repository.Add(uow, book); err != nil {
		uow.Complete()
		return err
	}
	uow.Commit()
	return nil
}
