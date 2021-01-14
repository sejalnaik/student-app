import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders, HttpParams } from "@angular/common/http";
import { Observable } from 'rxjs';
import { Student, StudentSearch } from "../classes/student";
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

  getStudents():Observable<any>{

    return this.httpClient.get<any>(this.baseUrl, {observe: "response"});
  }
  
  addStudent(student:Student):Observable<any>{
    //create header instance
    let httpHeaders:HttpHeaders = new HttpHeaders()
    
    //add token to header from cookie
    httpHeaders =  httpHeaders.append("token", this.cookieService.get("token"));

    return this.httpClient.post<any>(this.baseUrl, student, {'headers':httpHeaders, responseType:'text' as 'json', observe: "response"});
  }

  getStudent(id:string):Observable<any>{

    return this.httpClient.get<any>(this.baseUrl + "/" + id, {observe: "response"});
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
    return this.httpClient.get<any>(this.baseUrl + "/sum", {responseType:'text' as 'json'});
  }

  diffOfAgeAndRollNo(){
    return this.httpClient.get<any>(this.baseUrl + "/diff", {responseType:'text' as 'json'});
  }

  diffOfAgeAndRecordCount(){
    return this.httpClient.get<any>(this.baseUrl + "/diff-age-record-count", {responseType:'text' as 'json'});
  }

  getStudentTotalPenalty(id:string){
    return this.httpClient.get<any>(this.baseUrl + "/penalty/" + id, {responseType:'text' as 'json'});
  }

  searchStudent(studentSerach:StudentSearch):Observable<any>{
    let url:string = this.baseUrl + "/search";
    let params:HttpParams = new HttpParams();

    for (let key of Object.keys(studentSerach)) {
      let value = studentSerach[key];
      if((value != "") && (value != null)){
        params = params.append(key, value)
      }
    }    
    return this.httpClient.get<any>(url, {params:params, observe: "response"});
  }
}

