import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from "@angular/common/http";
import { Observable } from 'rxjs';
import { Student } from "../classes/student";
import { CookieService } from 'ngx-cookie-service';

@Injectable({
  providedIn: 'root'
})
export class StudentService {

  students:Student[] = [];
  constructor(
    private httpClient:HttpClient,
    private cookieService: CookieService
    ){}
  baseUrl:string = "http://localhost:8080/api/students";

  getStudents():Observable<Student[]>{
    //create header instance
    let httpHeaders:HttpHeaders = new HttpHeaders()
    
    //add token to header from cookie
    httpHeaders =  httpHeaders.append("token", this.cookieService.get("token"));

    return this.httpClient.get<Student[]>(this.baseUrl, {'headers' : httpHeaders});
  }
  
  addStudent(student:Student):Observable<Student>{
    //create header instance
    let httpHeaders:HttpHeaders = new HttpHeaders()
    
    //add token to header from cookie
    httpHeaders =  httpHeaders.append("token", this.cookieService.get("token"));

    return this.httpClient.post<Student>(this.baseUrl, student, {'headers':httpHeaders, responseType:'text' as 'json'});
  }

  getStudent(id:string):Observable<Student>{
    //create header instance
    let httpHeaders:HttpHeaders = new HttpHeaders()
    
    //add token to header from cookie
    httpHeaders =  httpHeaders.append("token", this.cookieService.get("token"));

    return this.httpClient.get<Student>(this.baseUrl + "/" + id, {'headers' : httpHeaders});
  }

  updateStudent(student:Student):Observable<Student>{
    //create header instance
    let httpHeaders:HttpHeaders = new HttpHeaders()
    
    //add token to header from cookie
    httpHeaders =  httpHeaders.append("token", this.cookieService.get("token"));

    return this.httpClient.put<Student>(this.baseUrl + "/" + student.id, student, {'headers':httpHeaders, responseType:'text' as 'json'});
  }
  deleteStudent(id:string){
    //create header instance
    let httpHeaders:HttpHeaders = new HttpHeaders()
    
    //add token to header from cookie
    httpHeaders =  httpHeaders.append("token", this.cookieService.get("token"));

    return this.httpClient.delete<Student>(this.baseUrl + "/" + id, {'headers':httpHeaders, responseType:'text' as 'json'});
  }
}

