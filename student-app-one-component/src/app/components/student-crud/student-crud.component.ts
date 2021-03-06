import { Component, OnInit } from '@angular/core';
import { Validators, FormBuilder, FormControl } from '@angular/forms';
import { Router } from '@angular/router';
import { BookIssues, Student, Book, BookWithAvailable, StudentSearch, StudentId } from 'src/app/classes/student';
import { StudentService } from 'src/app/services/student.service';
import { NgbModal } from '@ng-bootstrap/ng-bootstrap';
import { CookieService } from 'ngx-cookie-service';
import { BookService } from 'src/app/services/book.service';
import { BookIssueService } from 'src/app/services/book-issue.service';
import { Observable } from 'rxjs';

@Component({
  selector: 'app-student-crud',
  templateUrl: './student-crud.component.html',
  styleUrls: ['./student-crud.component.css']
})
export class StudentCrudComponent implements OnInit {

  students:Student[] = [];
  booksWithAvailable:BookWithAvailable[] = []
  bookIssues:BookIssues[] = []

  studentId:StudentId = {id:""};
  studentForm:any;
  studentSearchForm:any;
  studentAPI:Student;
  searchedStudent:StudentSearch
  addOrUpdateAction:string;
  modalRef: any;
  sumOfAgeAndRollNo:number;
  diffOfAgeAndRollNo:number;
  diffOfAgeAndRecordCount:number;
  totalPenalty:number;
  booksDropdown:string[] = [];
  
  constructor(
    private studentService:StudentService,
    private bookService:BookService, 
    private bookIssueService:BookIssueService,  
    private router:Router, 
    private formBuilder:FormBuilder,
    private modalService: NgbModal,
    private cookieService: CookieService
    ) { }

  ngOnInit(): void {
    this.getStudents();
    this.getBooks();
    this.createStudentForm();
    this.createStudentSearchForm();
    this.populateBooksDropdown();
  }

  populateBooksDropdown():void{
    this.bookService.getBooks().subscribe((data)=>{
      for(let i = 0; i < data.body.length; i++){
       this.booksDropdown = [...this.booksDropdown, data.body[i].name];
      }
    },
    (err) => {
      console.log('HTTP Error', err);
      alert(err.error)
    });
  }

   //create student add/update form
  createStudentForm(){
    this.studentForm = this.formBuilder.group({
      rollNo: [null, Validators.min(0)],
      name: ['', [Validators.required,  Validators.pattern("^[a-zA-Z_ ]+$")]],
      age: [null, Validators.min(0)],
      dob: [null],
      dobTime: [null],
      gender: [null],
      email: ['', [Validators.required, Validators.pattern("^[a-zA-Z0-9._%+-]+@[a-z0-9.-]+\\.[a-z]{2,4}$")]],
      phoneNumber:[null, [Validators.minLength(10), Validators.maxLength(12)]]
    });
  }

  //create student search form
  createStudentSearchForm(){
    this.studentSearchForm = this.formBuilder.group({
      name: ['', [Validators.pattern("^[a-zA-Z_ ]+$")]],
      from: [''],
      to: [''],
      email: [''],
      age:[null],
      books:['']
    });
  }

  //get all students
  getStudents():void{
    this.studentService.getStudents().subscribe((data)=>{
      this.students = data.body;
    },
    (err) => {
      console.log('HTTP Error', err);
      alert(err.error)
    });

    /*this.studentService.sumOfAgeAndRollNo().subscribe((data)=>{
      this.sumOfAgeAndRollNo = (JSON.parse(data))["Total"];
    },
    (err) => {
      console.log('HTTP Error', err);
      alert(err.error)
    });

    this.studentService.diffOfAgeAndRollNo().subscribe((data)=>{
      this.diffOfAgeAndRollNo = (JSON.parse(data))["Total"];
    },
    (err) => {
      console.log('HTTP Error', err);
      alert(err.error)
    });

    this.studentService.diffOfAgeAndRecordCount().subscribe((data)=>{
      this.diffOfAgeAndRecordCount = (JSON.parse(data))["Total"];
    },
    (err) => {
      console.log('HTTP Error', err);
      alert(err.error)
    });*/
  }
  
  //get all books
  getBooks(){
    this.bookService.getBooks().subscribe((data)=>{
      this.booksWithAvailable = data.body;
    },
    (err) => {
      console.log('HTTP Error', err);
      alert(err.error)
    });
  }

  //get all penalty for a student
  getTotalPenalty(id:string){
    this.studentService.getStudentTotalPenalty(id).subscribe((data)=>{
      this.totalPenalty = (JSON.parse(data))["Total"];
      console.log("total penlaty" + this.totalPenalty);
    },
    (err) => {
      console.log('HTTP Error', err);
      alert(err.error)
    });
  }

  //to check if add or update operation
  validate():void{
    if(this.studentForm.valid){
      if(this.addOrUpdateAction == "add"){
        this.addStudent();
      }
      else{
        this.updateStudent();
      }
    }
  }

  //add student button clicked
  onAddButtonClick(studentFormModal):void{
    /*if (this.cookieService.get("token") == ""){
      alert("Not authorized to access, please login first")
      this.router.navigate(["/login"]);
      return
    }*/
    this.setAddAction()
    this.openStudentFormModal(studentFormModal)
  }

  //update student buttom clicked
  onUpdateButtonClick(id:string, studentFormModal:any):void{
    /*if (this.cookieService.get("token") == ""){
      alert("Not authorized to access, please login first")
      this.router.navigate(["/login"]);
      return
    }*/
    this.prepopulate(id)
    this.openStudentFormModal(studentFormModal)
  }

  //book issues button clicked
  onBookIssuesButtonClick(id:string, bookIssuesModal:any):void{
    this.studentId.id = id;
    this.bookIssueService.getBookIssues(id).subscribe((data)=>{
      this.bookIssues = data;

      //open modal
      this.openBookIssuesModal(bookIssuesModal);

      //get total penalty for the student
      this.getTotalPenalty(id);

      //get total books
      this.getBooks();
    },
    (err) => {
      console.log('HTTP Error', err);
      alert(err.error)
    });
  }

  getBookIssues(){
    this.bookIssueService.getBookIssues(this.studentId.id).subscribe((data)=>{
      this.bookIssues = data;

      //get total penalty for the student
      this.getTotalPenalty(this.studentId.id);

      //get total books
      this.getBooks();
    },
    (err) => {
      console.log('HTTP Error', err);
      alert(err.error)
    });
  }

  //on issue book button click
  onIssueBookButtonClick(book:Book){
    let bookIssue:BookIssues = {id:null, bookId:book.id, studentId:this.studentId.id, book:book, issueDate:null, returned:null, penalty:null}
    this.bookIssueService.addBookIssue(bookIssue).subscribe((data)=>{  
      this.getBookIssues();
        alert("Book issue added with :" + data ); 
      },
      (err) => {
        console.log('HTTP Error', err);
        if (err.status == 401){
          alert("Session has expired, please login first")
          this.router.navigate(["/login"]);
          return
        }
        alert(err.error)
      });
  }

  //on returned button click
  onReturnedButtonClick(bookIssue:BookIssues){
    console.log(bookIssue);
    this.bookIssueService.upadteBookIssue(bookIssue, bookIssue.id).subscribe((data)=>{
      this.getBookIssues();
      this.getBooks();
      alert("Book returned"); 
    },
    (err) => {
      console.log('HTTP Error', err);
      if (err.status == 401){
        alert("Session has expired, please login first")
        this.router.navigate(["/login"]);
        return
      }
      alert(err.error)
    });
  }

  onSearchButtonClick(){
    //create search student
    this.searchedStudent = {
      name: this.studentSearchForm.get('name').value,
      email: this.studentSearchForm.get('email').value,
      from:this.studentSearchForm.get('from').value,
      to:this.studentSearchForm.get('to').value,
      age:this.studentSearchForm.get('age').value,
      books:this.studentSearchForm.get('books').value
    }
    
    //call search studnets service
    //this.studentService.searchStudent(this.searchedStudent)
    this.studentService.searchStudent(this.searchedStudent).subscribe((data)=>{
      this.students = data.body
    },
    (err) => {
      console.log('HTTP Error', err);
      alert(err.error)
    });
  }

  resetSearchForm():void{
    this.createStudentSearchForm();
    this.getStudents();
  }

  //add student
  addStudent():void{
    this.studentAPI = {id:null, 
                      rollNo:this.studentForm.get('rollNo').value, 
                      name:this.studentForm.get('name').value, 
                      age:this.studentForm.get('age').value, 
                      email:this.studentForm.get('email').value, 
                      isMale:this.studentForm.get('gender').value, 
                      dob:this.studentForm.get('dob').value,
                      dobTime:this.studentForm.get('dobTime').value,
                      phoneNumber:this.studentForm.get('phoneNumber').value};
    this.studentService.addStudent(this.studentAPI).subscribe(data=>{
      this.modalRef.close();
      this.getStudents();
      alert("Student added with id :" + data.body);
      if (data.headers.get("token") != null){
        this.cookieService.set("token", data.headers.get("token"));
      }
    },
      (err) => {
      console.log('HTTP Error', err);
      if (err.status == 401){
        alert("Session has expired, please login first")
        this.router.navigate(["/login"]);
        return
      }
      alert(err.error)
      });
    }

    //to calculate age from date of birth
    dobChange():void{
      let dobDate:Date = new Date(this.studentForm.controls['dob'].value);
      let diff = (new Date().getTime() - dobDate.getTime());
      let ageTotal = Math.trunc(diff/ (1000 * 3600 * 24 *365));
      this.studentForm.patchValue({
        age: ageTotal,
      });
    }

    //set add or update student form
    setAddAction():void{
      this.createStudentForm();
      this.addOrUpdateAction = "add";
    }

    //prepopulate update student form
    prepopulate(id:string):void{
      this.createStudentForm();
      this.addOrUpdateAction = "update";
      this.studentId.id = id;
      this.studentService.getStudent(id).subscribe((data)=>{
        this.studentForm.patchValue({
          name: data.body.name,
          rollNo: data.body.rollNo,
          age: data.body.age,
          dob: data.body.dob,
          dobTime: data.body.dobTime,
          email: data.body.email,
          gender: data.body.isMale,
          phoneNumber:data.body.phoneNumber
        });
      },
      (err) => {
        console.log('HTTP Error', err);
        alert(err.error)
      });
    }

    //update student
    updateStudent():void{
      this.studentAPI = {
        id:this.studentId.id, 
        rollNo:this.studentForm.get('rollNo').value, 
        name:this.studentForm.get('name').value, 
        age:this.studentForm.get('age').value, 
        email:this.studentForm.get('email').value, 
        isMale:this.studentForm.get('gender').value, 
        dob:this.studentForm.get('dob').value,
        dobTime:this.studentForm.get('dobTime').value,
        phoneNumber:this.studentForm.get('phoneNumber').value};

      //this.dateEmptyToNull(this.studentAPI)
      this.studentService.updateStudent(this.studentAPI).subscribe((data)=>{
        this.modalRef.close();
        this.getStudents();
        alert("Student updated"); 
      },
      (err) => {
        console.log('HTTP Error', err);
        if (err.status == 401){
          alert("Session has expired, please login first")
          this.router.navigate(["/login"]);
          return
        }
        alert(err.error)
      });
    }

    //delete student
    deleteStudent(id:string):void{
      /*if (this.cookieService.get("token") == ""){
        alert("Not authorized to access, please login first")
        this.router.navigate(["/login"]);
        return
      }*/
      if(confirm("Are you sure you want to delete?")) {
        this.studentService.deleteStudent(id).subscribe((data)=>{
          this.getStudents();
          alert("Student deleted");
        },
        (err) => {
          console.log('HTTP Error', err);
          if (err.status == 401){
            alert("Session has expired, please login first")
            this.router.navigate(["/login"]);
            return
          }
          alert(err.error)
        });
      }
    }

    //open student add/update form modal
    openStudentFormModal(studentFormModal: any):void {
      this.modalRef = this.modalService.open(studentFormModal, { ariaLabelledBy: 'modal-basic-title', backdrop: 'static', size: 'xl' });
      /*this.modalRef.result.then((result) => {
      }, (reason) => {
      });*/
    }

    //open book issue modal
    openBookIssuesModal(bookIssuesModal: any):void {
      this.modalRef = this.modalService.open(bookIssuesModal, { ariaLabelledBy: 'modal-basic-title', backdrop: 'static', size: 'xl' });
      /*this.modalRef.result.then((result) => {
      }, (reason) => {
      });*/
    }
}
