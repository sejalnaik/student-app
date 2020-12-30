import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { User } from "../classes/user";
import { CookieService } from 'ngx-cookie-service';

@Injectable({
  providedIn: 'root'
})
export class UserService {
  
  constructor(private httpClient:HttpClient, private cookieService: CookieService){}
  baseUrl:string = "http://localhost:8080";

  login(user:User):Observable<any>{
    return this.httpClient.post<any>(this.baseUrl + "/login", user, {responseType:'text' as 'json'});
   }

   register(user:User):Observable<User>{
    return this.httpClient.post<User>(this.baseUrl + "/register", user, {responseType:'text' as 'json'});
  }
}
