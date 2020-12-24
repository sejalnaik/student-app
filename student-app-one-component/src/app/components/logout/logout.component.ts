import { Component, OnInit } from '@angular/core';
import { CookieService } from 'ngx-cookie-service';
import { Router } from '@angular/router';

@Component({
  selector: 'app-logout',
  templateUrl: './logout.component.html',
  styleUrls: ['./logout.component.css']
})
export class LogoutComponent implements OnInit {

  constructor(
    private cookieService: CookieService,
    private router:Router
    ) {
        //delete token cookie
        cookieService.delete("token")

        //redirect to list
        this.router.navigate(["/home"]);
      }

  ngOnInit(): void {

  }

}
