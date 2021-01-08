import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from "@angular/common/http";
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
  searchUrl:string = "http://localhost:8080/students-search";
  sumUrl:string = "http://localhost:8080/sum";
  diffUrl:string = "http://localhost:8080/diff";
  diffAgeRecordCount:string = "http://localhost:8080/diff-age-record-count"
  totalPenlatyUrl:string = "http://localhost:8080/student-penalty"

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
    return this.httpClient.get<any>(this.sumUrl, {responseType:'text' as 'json'});
  }

  diffOfAgeAndRollNo(){
    return this.httpClient.get<any>(this.diffUrl, {responseType:'text' as 'json'});
  }

  diffOfAgeAndRecordCount(){
    return this.httpClient.get<any>(this.diffAgeRecordCount, {responseType:'text' as 'json'});
  }

  getStudentTotalPenalty(id:string){
    return this.httpClient.get<any>(this.totalPenlatyUrl + "/" + id, {responseType:'text' as 'json'});
  }

  searchStudent(studentSerach:StudentSearch):Observable<any>{
    let url:string;
    let paramsSet:string[] = []; 
    
    //create query params key value pairs
    for (let key of Object.keys(studentSerach)) {
      let value = studentSerach[key];
      if(value == ""){
        continue
      }
      paramsSet.push(key + "=" + value);
    }
    if(paramsSet.length == 0){
      url = this.searchUrl;
    }
    else{
      url = this.searchUrl + "?" + paramsSet.join("&");
    }
    console.log(url)
    return this.httpClient.get<any>(url, {observe: "response"});
  }
}

