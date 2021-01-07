import { Injectable } from '@angular/core';
import { Book } from "../classes/student";
import { HttpClient, HttpHeaders } from "@angular/common/http";
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class BookService {

  books:Book[] = [];
  baseUrl:string = "http://localhost:8080/books";

  constructor(private httpClient:HttpClient) { }

  getBooks():Observable<any>{
    return this.httpClient.get<any>(this.baseUrl, {observe: "response"});
  }
}
