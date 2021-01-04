package bookissuesservice

import (
	"github.com/jinzhu/gorm"
	"github.com/sejalnaik/student-app/model"
	"github.com/sejalnaik/student-app/repository"
)

type BookIssuesService struct {
	repository repository.Repository
	db         *gorm.DB
}

func NewBookIssuesService(repository repository.Repository, db *gorm.DB) *BookIssuesService {
	return &BookIssuesService{
		repository: repository,
		db:         db,
	}
}

func (s *BookIssuesService) GetAllBookIssues(bookIssues *[]model.BookIssue) error {
	//create unit of work
	uow := repository.NewUnitOfWork(s.db, true)

	//give query processor for where
	queryProcessors := []repository.QueryProcessor{}
	queryProcessors = append(queryProcessors, repository.PreloadAssociations([]string{"Book"}))

	//call get repository method to get book issues
	if err := s.repository.Get(uow, bookIssues, queryProcessors); err != nil {
		return err
	} else {
		return nil
	}
}

func (s *BookIssuesService) AddBookIssue(bookIssue *model.BookIssue) error {
	//create unit of work
	uow := repository.NewUnitOfWork(s.db, true)

	//create bucket for book
	book := model.Book{}

	//call get repo to get book
	//give query processor for where
	queryProcessors := []repository.QueryProcessor{}
	queryProcessors = append(queryProcessors, repository.Where("id=?", bookIssue.BookID))
	if err := s.repository.Get(uow, &book, queryProcessors); err != nil {
		return err
	}

	//add book to book issue
	bookIssue.Book = book

	//call add repository method to add one book issue
	if err := s.repository.Add(uow, bookIssue); err != nil {
		uow.Complete()
		return err
	}
	uow.Commit()
	return nil
}

func (s *BookIssuesService) GetBookIssue(bookIssue *model.BookIssue, bookIssueID string) error {
	//create unit of work
	uow := repository.NewUnitOfWork(s.db, true)

	//give query processor for where
	queryProcessors := []repository.QueryProcessor{}
	queryProcessors = append(queryProcessors, repository.Where("id=?", bookIssueID))
	queryProcessors = append(queryProcessors, repository.PreloadAssociations([]string{"Book"}))

	//call get repository method to get one bookIsuue
	if err := s.repository.Get(uow, bookIssue, queryProcessors); err != nil {
		return err
	} else {
		return nil
	}
}
