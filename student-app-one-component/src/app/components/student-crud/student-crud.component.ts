import { Component, OnInit } from '@angular/core';
import { Validators, FormBuilder, FormControl } from '@angular/forms';
import { Router } from '@angular/router';
import { Student } from 'src/app/classes/student';
import { StudentService } from 'src/app/services/student.service';
import { NgbModal } from '@ng-bootstrap/ng-bootstrap';
import { NgxSpinnerService } from "ngx-spinner";

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
  
  constructor(
    private studentService:StudentService, 
    private router:Router, 
    private formBuilder:FormBuilder,
    private modalService: NgbModal,
    private spinner: NgxSpinnerService,
    ) { 
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
      phoneNumber:[null]
    });
  }
  
  ngOnInit(): void {
    this.spinner.show();
    this.getStudents();
   }

  getStudents():void{
    this.studentService.getStudents().subscribe((data)=>{
      this.students = data;
      //this.spinner.hide();
    },
    (err) => {
      console.log('HTTP Error', err);
      this.spinner.hide();
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
    this.setAddAction()
    this.openStudentFormModal(studentFormModal)
  }

  onUpdateButtonClick(id:string, studentFormModal):void{
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
      alert("Student added with id :" + data);
    },
      (err) => {
        console.log('HTTP Error', err);
        this.spinner.hide() 
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
          name: data.name,
          rollNo: data.rollNo,
          age: data.age,
          dob: data.dob,
          dobTime: data.dobTime,
          email: data.email,
          gender: data.isMale,
          phoneNumber:data.phoneNumber
        });
        this.spinner.hide()
      },
      (err) => {
        console.log('HTTP Error', err);
        this.spinner.hide()
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
        alert("Student updated with id :" + data); 
      },
      (err) => {
        console.log('HTTP Error', err);
        this.spinner.hide()
      });
    }

    deleteStudent(id:string):void{
      if(confirm("Are you sure to delete?")) {
        this.studentService.deleteStudent(id).subscribe((data)=>{
          this.spinner.show()
          this.getStudents();
          alert("Student deleted with id :" + data);
        },
        (err) => console.log('HTTP Error', err)
        );
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
