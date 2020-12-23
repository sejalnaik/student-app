import { Component, ElementRef, OnInit, ViewChild } from '@angular/core';
import { FormBuilder, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { NgbModal } from '@ng-bootstrap/ng-bootstrap';
import { User } from 'src/app/classes/user';
import { UserService } from 'src/app/services/user.service';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {

  @ViewChild('loginFormModalButton') loginFormModalButton:ElementRef;
  loginForm:any;
  user:User;
  modalRef: any;
  loginOrRegisterAction:string = "login"
  wrongLoginDetailsErrorShow:string = "none";

  constructor(
    private userService:UserService, 
    private router:Router, 
    private formBuilder:FormBuilder,
    private modalService: NgbModal
  ) { }

  ngOnInit(): void {
    this.formBuild()
  }

  ngAfterViewInit() {
    this.loginFormModalButton.nativeElement.click();
  }

  formBuild():void{
    this.loginForm = this.formBuilder.group({
      username: ['', [Validators.required, Validators.pattern("^[a-zA-Z_ ]+$")]],
      password: ['', Validators.required],
    });
  }

  login():void{
    this.user = { 
      username:this.loginForm.get('username').value, 
      password:this.loginForm.get('password').value
    };
    this.userService.login(this.user).subscribe(data=>{
      this.modalRef.close();
      this.router.navigate(["/home"]);
    },
      (err) => {
        console.log('HTTP Error', err);
        this.wrongLoginDetailsErrorShow = "block"
      }
    );
  }

  openLoginFormModal(loginFormModal: any):void {
    this.modalRef = this.modalService.open(loginFormModal, { ariaLabelledBy: 'modal-basic-title', backdrop: 'static', size: 'xl' });
    /*this.modalRef.result.then((result) => {
    }, (reason) => {
    });*/
  }

  setRegisterForm():void{
    this.formBuild()
    this.loginOrRegisterAction = "register"
  }

  setLoginForm():void{
    this.formBuild()
    this.loginOrRegisterAction = "login"
  }

  validate():void{
  
    if(this.loginForm.valid){
      if(this.loginOrRegisterAction == "login"){
        this.login();
      }
      else{
        this.register();
      }
    }
  }

  register():void{
    this.user = { 
      username:this.loginForm.get('username').value, 
      password:this.loginForm.get('password').value
    };
    this.userService.register(this.user).subscribe(data=>{
      alert("Successfully registered with id:" + data)
      this.setLoginForm()
      this.wrongLoginDetailsErrorShow = "none"
    },
      (err) => {
        console.log('HTTP Error', err);
      }
    );
  }
}
