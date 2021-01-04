package bookcontroller

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	service "github.com/sejalnaik/student-app/book/book-service"
	"github.com/sejalnaik/student-app/model"
)

type BookController struct {
	bookService *service.BookService
}

func NewBookController(bookService *service.BookService) *BookController {
	return &BookController{
		bookService: bookService,
	}
}

func (c *BookController) CreateRoutes(r *mux.Router) {
	//create route for get books
	r.HandleFunc("/books", c.GetAllBooks).Methods("GET")

	//create route for adding one book
	r.HandleFunc("/books", c.AddBook).Methods("POST")

	//create route for getting one book
	r.HandleFunc("/books/{bookID}", c.GetBook).Methods("GET")
}

func (c *BookController) GetAllBooks(w http.ResponseWriter, r *http.Request) {
	log.Println("Get books called")

	//create bucket
	booksWithAvailable := []model.BookWithAvailable{}

	//calling service method to get all books
	if err := c.bookService.GetAllBooks(&booksWithAvailable); err != nil {
		log.Println("Get books unsuccessful")
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	//converting struct to json type and sending back json
	if booksJSON, err := json.Marshal(booksWithAvailable); err != nil {
		log.Println("Get books : JSON marshall unsuccessful")
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		log.Println("Get books successful")
		w.Write(booksJSON)
	}
}

func (c *BookController) AddBook(w http.ResponseWriter, r *http.Request) {
	log.Println("Add book called")

	//create bucket
	book := &model.Book{}

	//read book data from response body
	responseBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Add book : Could not read response body")
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	//convert json to struct type
	er := json.Unmarshal(responseBody, book)
	if er != nil {
		log.Println("Add book : Json unmarshall unsuccessful")
		log.Println(er)
		http.Error(w, er.Error(), http.StatusBadRequest)
		return
	}

	//calling service method to add book and giving back id as string
	if err := c.bookService.AddBook(book); err != nil {
		log.Println("Add book unsuccessful")
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		log.Println("Add book successful")
		w.Write([]byte(book.ID.String()))
	}
}

func (c *BookController) GetBook(w http.ResponseWriter, r *http.Request) {
	log.Println("Get book called")

	//create bucket
	book := &model.Book{}

	//getting id from query param
	params := mux.Vars(r)
	bookID := (params["bookID"])

	//calling service method to get book
	if err := c.bookService.GetBook(book, bookID); err != nil {
		log.Println("Get book unsuccessful")
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	//converting struct to json type and sending back json
	if bookJSON, err := json.Marshal(book); err != nil {
		log.Println("Get book : JSON marshall unsuccessful")
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		log.Println("Get book successful")
		w.Write(bookJSON)
	}
}
