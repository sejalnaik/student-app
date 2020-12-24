import { BrowserModule } from '@angular/platform-browser';
import { NgModule, CUSTOM_ELEMENTS_SCHEMA } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { HttpClientModule } from "@angular/common/http";
import {  RouterModule } from '@angular/router';
import {  ReactiveFormsModule } from "@angular/forms";
import { NgbModule } from '@ng-bootstrap/ng-bootstrap';
import { NgxSpinnerModule } from "ngx-spinner";
import { CookieService } from 'ngx-cookie-service';

import { AppComponent } from './app.component';
import { StudentCrudComponent } from './components/student-crud/student-crud.component';
import { StudentService } from "./services/student.service";
import { UserService } from "./services/user.service";
import { HomeComponent } from './components/home/home.component';
import { ErrorComponent } from './components/error/error.component';
import { RoutesClass } from "./classes/route-class";
import { EmptyToNullDirectveDirective } from './directives/empty-to-null-directve.directive';
import { LoginComponent } from './components/login/login.component';

@NgModule({
  declarations: [
    AppComponent,
    StudentCrudComponent,
    HomeComponent,
    ErrorComponent,
    EmptyToNullDirectveDirective,
    LoginComponent
  ],
  imports: [
    BrowserModule,
    FormsModule,
    HttpClientModule,
    RouterModule.forRoot(RoutesClass.routes),
    ReactiveFormsModule,
    NgbModule,
    NgxSpinnerModule
  ],
  schemas: [CUSTOM_ELEMENTS_SCHEMA],
  providers: [
    StudentService,
    UserService,
    CookieService
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
