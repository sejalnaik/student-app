package bookissuesservice

import (
	"errors"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/sejalnaik/student-app/model"
	"github.com/sejalnaik/student-app/repository"
	"github.com/sejalnaik/student-app/utility"
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

func (s *BookIssuesService) GetAllBookIssues(bookIssues *[]model.BookIssue, studentID string) error {
	//create unit of work
	uow := repository.NewUnitOfWork(s.db, true)

	//update book issues(for updating penalty)
	updateBookIssue := map[string]interface{}{"penalty": gorm.Expr("if(abs(DATEDIFF(issue_date, CURDATE())) > 10 and returned = 0, (abs(DATEDIFF(issue_date, CURDATE()))-10)*2.5, 0)")}
	queryProcessors := []repository.QueryProcessor{}
	queryProcessors = append(queryProcessors, repository.Where("student_id=?", studentID))
	if err := s.repository.Update(uow, &model.BookIssue{}, updateBookIssue, queryProcessors); err != nil {
		return err
	}

	//give query processor for preload where
	queryProcessors = []repository.QueryProcessor{}
	queryProcessors = append(queryProcessors, repository.Where("student_id=?", studentID))
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

	//create bucket for booKs with available column
	bookWithAvailable := &model.BookWithAvailable{}

	//get book with available column
	queryProcessors := []repository.QueryProcessor{}
	queryProcessors = append(queryProcessors, repository.Model(&model.Book{}))
	queryProcessors = append(queryProcessors, repository.Where("books.id=?", bookIssue.Book.ID))
	queryProcessors = append(queryProcessors, repository.Select("(MIN(total_stock) - COUNT(IF(book_issues.returned = 0, 1, NULL))) AS available"))
	queryProcessors = append(queryProcessors, repository.Joins("left join book_issues on book_issues.book_id = books.id"))
	queryProcessors = append(queryProcessors, repository.Group("books.id"))
	queryProcessors = append(queryProcessors, repository.Group("books.name"))
	if err := s.repository.Scan(uow, &model.Book{}, bookWithAvailable, queryProcessors); err != nil {
		return err
	}

	//check if available is 0
	if bookWithAvailable.Available == 0 {
		return errors.New("Book is not available")
	}

	//check if book is already issued
	queryProcessors = []repository.QueryProcessor{}
	queryProcessors = append(queryProcessors, repository.Where("student_id=?", bookIssue.StudentID))
	queryProcessors = append(queryProcessors, repository.Where("book_id=?", bookIssue.Book.ID))
	queryProcessors = append(queryProcessors, repository.Where("returned=?", "0"))
	if err := s.repository.Get(uow, &model.BookIssue{}, queryProcessors); err == nil {
		return errors.New("Book is already issued")
	} else if err.Error() == "record not found" {
		goto canBeAdded
	} else {
		return err
	}
canBeAdded:
	//create issue date and time and give it to book issue
	currentTime := time.Now()
	currentTime.Format("2006.01.02 15:04:05")
	currentTimeInString := currentTime.String()
	currentTimeInString = currentTimeInString[:19]
	bookIssue.IssueDate = currentTimeInString

	//give returned as false
	bookIssue.Returned = false

	//give penalty as 0.0
	bookIssue.Penalty = 0.0

	//check if book exists
	queryProcessors = []repository.QueryProcessor{}
	queryProcessors = append(queryProcessors, repository.Where("id=?", bookIssue.BookID))
	if err := s.repository.Get(uow, &model.Book{}, queryProcessors); err != nil {
		return err
	}

	//check if student exists
	queryProcessors = []repository.QueryProcessor{}
	queryProcessors = append(queryProcessors, repository.Where("id=?", bookIssue.StudentID))
	if err := s.repository.Get(uow, &model.Student{}, queryProcessors); err != nil {
		return err
	}

	//call add repository method to add one book issue

	if err := s.repository.Add(uow, bookIssue); err != nil {
		uow.Complete()
		return err
	}
	uow.Commit()
	return nil
}

/*func (s *BookIssuesService) GetBookIssue(bookIssue *model.BookIssue, bookIssueID string) error {
	//create unit of work
	uow := repository.NewUnitOfWork(s.db, true)

	//give query processor for where
	queryProcessors := []repository.QueryProcessor{}
	queryProcessors = append(queryProcessors, repository.Where("id=?", bookIssue.BookID))
	if err := s.repository.Get(uow, &model.Book{}, queryProcessors); err != nil {
		return err
	}

	//call get repository method to get one bookIsuue
	if err := s.repository.Get(uow, bookIssue, queryProcessors); err != nil {
		return err
	} else {
		return nil
	}
}*/

func (s *BookIssuesService) UpdateBookIssue(bookIssue *model.BookIssue, bookIssueID string) error {
	//create unit of work
	uow := repository.NewUnitOfWork(s.db, true)

	//check if book exists
	queryProcessors := []repository.QueryProcessor{}
	queryProcessors = append(queryProcessors, repository.Where("id=?", bookIssue.BookID))
	if err := s.repository.Get(uow, &model.Book{}, queryProcessors); err != nil {
		return err
	}

	//check if student exists
	queryProcessors = []repository.QueryProcessor{}
	queryProcessors = append(queryProcessors, repository.Where("id=?", bookIssue.StudentID))
	if err := s.repository.Get(uow, &model.Student{}, queryProcessors); err != nil {
		return err
	}

	//give retuened as true to book issue
	bookIssue.Returned = true

	//give penalty as 0.0 to book issue
	bookIssue.Penalty = 0.0

	//convert book issue struct to map
	bookIssueMap := utility.ConvertStructBookIssueToMap(bookIssue)

	//give query processor for where
	queryProcessors = []repository.QueryProcessor{}
	queryProcessors = append(queryProcessors, repository.Where("id=?", bookIssueID))

	//call update repository method to update one book issue
	if err := s.repository.Update(uow, bookIssue, bookIssueMap, queryProcessors); err != nil {
		uow.Complete()
		return err
	}

	/*if err := s.repository.Save(uow, bookIssue, queryProcessors); err != nil {
		uow.Complete()
		return err
	}*/
	uow.Commit()
	return nil
}
