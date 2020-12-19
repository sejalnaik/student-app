import { Component, OnInit } from '@angular/core';
import { Validators, FormBuilder, FormControl } from '@angular/forms';
import { Router } from '@angular/router';
import { Student } from 'src/app/classes/student';
import { StudentService } from 'src/app/services/student.service';

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
  
  constructor(
    private studentService:StudentService, 
    private router:Router, 
    private formBuilder:FormBuilder
    ) { 
      this.formBuild();
  }

  formBuild(){
    this.addForm = this.formBuilder.group({
      rollNo: ['', Validators.required],
      name: ['', [Validators.required,  Validators.pattern("^[a-zA-Z_ ]+$")]],
      age: ['', Validators.required],
      dob: ['', Validators.required],
      gender: ['', Validators.required],
      email: ['', [Validators.required, Validators.email]]
    });
  }
  
  ngOnInit(): void {
    this.getStudents();
   }

  getStudents():void{
    this.studentService.getStudents().subscribe((data)=>{
      this.students = data;
    },
    (err) => console.log('HTTP Error', err)
    );
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

  addStudent():void{
    this.studentAPI = {id:null, 
                      rollNo:this.addForm.get('rollNo').value, 
                      name:this.addForm.get('name').value, 
                      age:this.addForm.get('age').value, 
                      email:this.addForm.get('email').value, 
                      isMale:this.addForm.get('gender').value, 
                      dob:this.addForm.get('dob').value};
    this.studentService.addStudent(this.studentAPI).subscribe(data=>{
      this.getStudents();
      alert("Student added");
    },
    (err) => console.log('HTTP Error', err)
    );
    }

    dobChange():void{
      let dobDate:Date = new Date(this.addForm.controls['date'].value);
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
      this.formBuild();
      this.addOrUpdateAction = "update";
      this.id = id;
      this.studentService.getStudent(id).subscribe((data)=>{
        this.addForm.patchValue({
          name: data[0].name,
          rollNo: data[0].rollNo,
          age: data[0].age,
          dob: data[0].dob,
          email: data[0].email,
          gender: data[0].isMale
        });
      },
      (err) => console.log('HTTP Error', err)
      );
    }

    updateStudent():void{
      this.studentAPI = {
        id:this.id, 
        rollNo:this.addForm.get('rollNo').value, 
        name:this.addForm.get('name').value, 
        age:this.addForm.get('age').value, 
        email:this.addForm.get('email').value, 
        isMale:this.addForm.get('gender').value, 
        dob:this.addForm.get('dob').value
      };

      this.studentService.updateStudent(this.studentAPI).subscribe((data)=>{
        this.getStudents();
        alert("Student updated"); 
      },
      (err) => console.log('HTTP Error', err)
      );
    }

    deleteStudent(id:string):void{
      if(confirm("Are you sure to delete?")) {
        this.studentService.deleteStudent(id).subscribe((data)=>{
          this.getStudents();
          alert("Student deleted");
        },
        (err) => console.log('HTTP Error', err)
        );
      }
    }
}
