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
  baseUrl:string = "http://localhost:8080/students";
  sumUrl:string = "http://localhost:8080/sum";
  diffUrl:string = "http://localhost:8080/diff";

  getStudents():Observable<any>{
    //create header instance
    let httpHeaders:HttpHeaders = new HttpHeaders()
    
    //add token to header from cookie
    httpHeaders =  httpHeaders.append("token", this.cookieService.get("token"));

    return this.httpClient.get<any>(this.baseUrl, {'headers' : httpHeaders, observe: "response"});
  }
  
  addStudent(student:Student):Observable<any>{
    //create header instance
    let httpHeaders:HttpHeaders = new HttpHeaders()
    
    //add token to header from cookie
    httpHeaders =  httpHeaders.append("token", this.cookieService.get("token"));

    return this.httpClient.post<any>(this.baseUrl, student, {'headers':httpHeaders, responseType:'text' as 'json', observe: "response"});
  }

  getStudent(id:string):Observable<any>{
    //create header instance
    let httpHeaders:HttpHeaders = new HttpHeaders()
    
    //add token to header from cookie
    httpHeaders =  httpHeaders.append("token", this.cookieService.get("token"));

    return this.httpClient.get<any>(this.baseUrl + "/" + id, {'headers' : httpHeaders, observe: "response"});
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

  sumOfAgeAndRollNo(){
    return this.httpClient.get<any>(this.sumUrl, {responseType:'text' as 'json'});
  }

  diffOfAgeAndRollNo(){
    return this.httpClient.get<any>(this.diffUrl, {responseType:'text' as 'json'});
  }
}

