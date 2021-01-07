import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { BookIssues } from '../classes/student';

@Injectable({
  providedIn: 'root'
})
export class BookIssueService {

  bookIssues:BookIssues[] = [];
  baseUrl:string = "http://localhost:8080/book-issues";

  constructor(private httpClient:HttpClient) { }

  getBookIssues(studentId:string):Observable<BookIssues[]>{
    return this.httpClient.get<BookIssues[]>(this.baseUrl + "/" + studentId);
  }

  addBookIssue(bookIssue:BookIssues):Observable<BookIssues>{
    return this.httpClient.post<BookIssues>(this.baseUrl, bookIssue, {responseType:'text' as 'json'});
  }

  upadteBookIssue(bookIssue:BookIssues, studentId:string):Observable<BookIssues>{
    return this.httpClient.put<BookIssues>(this.baseUrl + "/" + studentId, bookIssue, {responseType:'text' as 'json'});
  }
}
