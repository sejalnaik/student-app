package bookissuescontroller

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	service "github.com/sejalnaik/student-app/book-issues/book-issues-service"
	"github.com/sejalnaik/student-app/model"
)

type BookIssuesController struct {
	bookIssuesService *service.BookIssuesService
}

func NewBookIssuesController(bookIssuesService *service.BookIssuesService) *BookIssuesController {
	return &BookIssuesController{
		bookIssuesService: bookIssuesService,
	}
}

func (c *BookIssuesController) CreateRoutes(r *mux.Router) {
	//create route for get books
	r.HandleFunc("/book-issues", c.GetAllBooksIssues).Methods("GET")

	//create route for adding one book
	r.HandleFunc("/book-issues", c.AddBookIssue).Methods("POST")
}

func (c *BookIssuesController) GetAllBooksIssues(w http.ResponseWriter, r *http.Request) {
	log.Println("Get book issues called")

	//create bucket
	bookIssues := []model.BookIssue{}

	//calling service method to get all books issues
	if err := c.bookIssuesService.GetAllBookIssues(&bookIssues); err != nil {
		log.Println("Get books issues unsuccessful")
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	//converting struct to json type and sending back json
	if booksIssuesJSON, err := json.Marshal(bookIssues); err != nil {
		log.Println("Get books issues : JSON marshall unsuccessful")
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		log.Println("Get books issues successful")
		w.Write(booksIssuesJSON)
	}
}

func (c *BookIssuesController) AddBookIssue(w http.ResponseWriter, r *http.Request) {
	log.Println("Add book issue called")

	//create bucket
	bookIssue := &model.BookIssue{}

	//read book data from response body
	responseBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Add book issue : Could not read response body")
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	//convert json to struct type
	er := json.Unmarshal(responseBody, bookIssue)
	if er != nil {
		log.Println("Add book issue : Json unmarshall unsuccessful")
		log.Println(er)
		http.Error(w, er.Error(), http.StatusBadRequest)
		return
	}

	//calling service method to add book issue and giving back id as string
	if err := c.bookIssuesService.AddBookIssue(bookIssue); err != nil {
		log.Println("Add book issue unsuccessful")
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		log.Println("Add book issue successful")
		w.Write([]byte(bookIssue.ID.String()))
	}
}
