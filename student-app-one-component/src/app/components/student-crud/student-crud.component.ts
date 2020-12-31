import { Component, OnInit } from '@angular/core';
import { Validators, FormBuilder, FormControl } from '@angular/forms';
import { Router } from '@angular/router';
import { Student } from 'src/app/classes/student';
import { StudentService } from 'src/app/services/student.service';
import { NgbModal } from '@ng-bootstrap/ng-bootstrap';
import { NgxSpinnerService } from "ngx-spinner";
import { CookieService } from 'ngx-cookie-service';

@Component({
  selector: 'app-student-crud',
  templateUrl: './student-crud.component.html',
  styleUrls: ['./student-crud.component.css']
})
export class StudentCrudComponent implements OnInit {

  students:Student[] = [];

  id:string;
  addForm:any;
  studentAPI:Student;
  addOrUpdateAction:string;
  modalRef: any;
  loadingMessage: string = "Getting students";
  sumOfAgeAndRollNo:number;
  diffOfAgeAndRollNo:number;
  diffOfAgeAndRecordCount:number;
  
  constructor(
    private studentService:StudentService, 
    private router:Router, 
    private formBuilder:FormBuilder,
    private modalService: NgbModal,
    private spinner: NgxSpinnerService,
    private cookieService: CookieService
    ) { }

  ngOnInit(): void {
    this.spinner.show();
    this.getStudents();
    this.formBuild();
   }

  formBuild(){
    this.addForm = this.formBuilder.group({
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

  getStudents():void{
    this.studentService.getStudents().subscribe((data)=>{
      this.students = data.body;
      //this.spinner.hide();
    },
    (err) => {
      this.spinner.hide();
      console.log('HTTP Error', err);
      alert(err.error)
    });

    this.studentService.sumOfAgeAndRollNo().subscribe((data)=>{
      this.sumOfAgeAndRollNo = (JSON.parse(data))["Total"];
      //this.spinner.hide();
    },
    (err) => {
      this.spinner.hide();
      console.log('HTTP Error', err);
      alert(err.error)
    });

    this.studentService.diffOfAgeAndRollNo().subscribe((data)=>{
      this.diffOfAgeAndRollNo = (JSON.parse(data))["Total"];
      //this.spinner.hide();
    },
    (err) => {
      this.spinner.hide();
      console.log('HTTP Error', err);
      alert(err.error)
    });

    this.studentService.diffOfAgeAndRecordCount().subscribe((data)=>{
      this.diffOfAgeAndRecordCount = (JSON.parse(data))["Total"];
      //this.spinner.hide();
    },
    (err) => {
      this.spinner.hide();
      console.log('HTTP Error', err);
      alert(err.error)
    });
  }
  

  validate():void{
    if(this.addForm.valid){
      if(this.addOrUpdateAction == "add"){
        this.addStudent();
      }
      else{
        this.updateStudent();
      }
    }
  }

  onAddButtonClick(studentFormModal):void{
    /*if (this.cookieService.get("token") == ""){
      alert("Not authorized to access, please login first")
      this.router.navigate(["/login"]);
      return
    }*/
    this.setAddAction()
    this.openStudentFormModal(studentFormModal)
  }

  onUpdateButtonClick(id:string, studentFormModal):void{
    /*if (this.cookieService.get("token") == ""){
      alert("Not authorized to access, please login first")
      this.router.navigate(["/login"]);
      return
    }*/
    this.prepopulate(id)
    this.openStudentFormModal(studentFormModal)
  }

  addStudent():void{
    this.studentAPI = {id:null, 
                      rollNo:this.addForm.get('rollNo').value, 
                      name:this.addForm.get('name').value, 
                      age:this.addForm.get('age').value, 
                      email:this.addForm.get('email').value, 
                      isMale:this.addForm.get('gender').value, 
                      dob:this.addForm.get('dob').value,
                      dobTime:this.addForm.get('dobTime').value,
                      phoneNumber:this.addForm.get('phoneNumber').value};
    this.studentService.addStudent(this.studentAPI).subscribe(data=>{
      this.spinner.show()
      this.modalRef.close();
      this.getStudents();
      alert("Student added with id :" + data.body);
      if (data.headers.get("token") != null){
        this.cookieService.set("token", data.headers.get("token"));
      }
    },
      (err) => {
        this.spinner.hide();
      console.log('HTTP Error', err);
      if (err.status == 401){
        alert("Session has expired, please login first")
        this.router.navigate(["/login"]);
        return
      }
      alert(err.error)
      });
    }

    dobChange():void{
      let dobDate:Date = new Date(this.addForm.controls['dob'].value);
      let diff = (new Date().getTime() - dobDate.getTime());
      let ageTotal = Math.trunc(diff/ (1000 * 3600 * 24 *365));
      this.addForm.patchValue({
        age: ageTotal,
      });
    }

    setAddAction():void{
      this.formBuild();
      this.addOrUpdateAction = "add";
    }

    prepopulate(id:string):void{
      this.spinner.show()
      this.formBuild();
      this.addOrUpdateAction = "update";
      this.id = id;
      this.studentService.getStudent(id).subscribe((data)=>{
        this.addForm.patchValue({
          name: data.body.name,
          rollNo: data.body.rollNo,
          age: data.body.age,
          dob: data.body.dob,
          dobTime: data.body.dobTime,
          email: data.body.email,
          gender: data.body.isMale,
          phoneNumber:data.body.phoneNumber
        });
        this.spinner.hide()
      },
      (err) => {
        this.spinner.hide();
        console.log('HTTP Error', err);
        alert(err.error)
      });
    }

    updateStudent():void{
      this.spinner.show()
      this.studentAPI = {
        id:this.id, 
        rollNo:this.addForm.get('rollNo').value, 
        name:this.addForm.get('name').value, 
        age:this.addForm.get('age').value, 
        email:this.addForm.get('email').value, 
        isMale:this.addForm.get('gender').value, 
        dob:this.addForm.get('dob').value,
        dobTime:this.addForm.get('dobTime').value,
        phoneNumber:this.addForm.get('phoneNumber').value
      };

      //this.dateEmptyToNull(this.studentAPI)
      this.studentService.updateStudent(this.studentAPI).subscribe((data)=>{
        this.modalRef.close();
        this.getStudents();
        alert("Student updated"); 
      },
      (err) => {
        this.spinner.hide();
        console.log('HTTP Error', err);
        if (err.status == 401){
          alert("Session has expired, please login first")
          this.router.navigate(["/login"]);
          return
        }
        alert(err.error)
      });
    }

    deleteStudent(id:string):void{
      /*if (this.cookieService.get("token") == ""){
        alert("Not authorized to access, please login first")
        this.router.navigate(["/login"]);
        return
      }*/
      if(confirm("Are you sure you want to delete?")) {
        this.studentService.deleteStudent(id).subscribe((data)=>{
          this.spinner.show()
          this.getStudents();
          alert("Student deleted");
        },
        (err) => {
          this.spinner.hide();
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

    /*dateEmptyToNull(studentAPI:Student):void{
      if (studentAPI.dob == ""){
        studentAPI.dob = null
      }
      if (studentAPI.dobTime == ""){
        studentAPI.dobTime = null
      }
    }*/

    openStudentFormModal(studentFormModal: any):void {
      this.modalRef = this.modalService.open(studentFormModal, { ariaLabelledBy: 'modal-basic-title', backdrop: 'static', size: 'xl' });
      /*this.modalRef.result.then((result) => {
      }, (reason) => {
      });*/
    }
}
