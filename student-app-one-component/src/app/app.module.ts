import { BrowserModule } from '@angular/platform-browser';
import { NgModule, CUSTOM_ELEMENTS_SCHEMA } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { HttpClientModule } from "@angular/common/http";
import {  RouterModule } from '@angular/router';
import {  ReactiveFormsModule } from "@angular/forms";
import { NgbModule } from '@ng-bootstrap/ng-bootstrap';
import { NgxSpinnerModule } from "ngx-spinner";

import { AppComponent } from './app.component';
import { StudentCrudComponent } from './components/student-crud/student-crud.component';
import { StudentService } from "./services/student.service";
import { HomeComponent } from './components/home/home.component';
import { ErrorComponent } from './components/error/error.component';
import { RoutesClass } from "./classes/route-class";
import { EmptyToNullDirectveDirective } from './directives/empty-to-null-directve.directive';

@NgModule({
  declarations: [
    AppComponent,
    StudentCrudComponent,
    HomeComponent,
    ErrorComponent,
    EmptyToNullDirectveDirective
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
    StudentService
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
