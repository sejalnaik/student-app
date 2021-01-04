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

func (s *BookService) GetAllBooks(booksWithAvailable *[]model.BookWithAvailable) error {
	//create unit of work
	uow := repository.NewUnitOfWork(s.db, true)

	//give query processor for where
	queryProcessors := []repository.QueryProcessor{}
	queryProcessors = append(queryProcessors, repository.Model())
	queryProcessors = append(queryProcessors, repository.Select("books.ID as id, books.name as name, MIN(total_stock) as total_stock, (MIN(total_stock) - COUNT(book_issues.returned)) AS available"))
	queryProcessors = append(queryProcessors, repository.Joins("inner join book_issues on book_issues.book_id = books.id"))
	queryProcessors = append(queryProcessors, repository.Where("book_issues.returned = 0"))
	queryProcessors = append(queryProcessors, repository.Group("books.id"))
	queryProcessors = append(queryProcessors, repository.Group("books.name"))

	//call get repository method to get books
	if err := s.repository.Scan(uow, &model.Book{}, booksWithAvailable, queryProcessors); err != nil {
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

func (s *BookService) GetBook(book *model.Book, bookID string) error {
	//create unit of work
	uow := repository.NewUnitOfWork(s.db, true)

	//give query processor for where
	queryProcessors := []repository.QueryProcessor{}
	queryProcessors = append(queryProcessors, repository.Where("id=?", bookID))

	//call get repository method to get one book
	if err := s.repository.Get(uow, book, queryProcessors); err != nil {
		return err
	} else {
		//to trim dob and dobtime
		return nil
	}
}
