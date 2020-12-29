import { Component, OnInit } from '@angular/core';
import { CookieService } from 'ngx-cookie-service';

@Component({
  selector: 'app-navbar',
  templateUrl: './navbar.component.html',
  styleUrls: ['./navbar.component.css']
})
export class NavbarComponent implements OnInit {

  logoutLinkShow:string;
  logInLinkShow:string;
  listLinkShow:string;

  constructor(private cookieService: CookieService) { 
    if (cookieService.get("token") == ""){
      this.logInLinkShow = "inline-block";
      this.logoutLinkShow = "none";
      this.listLinkShow = "none"
    }
    else{
      this.logInLinkShow = "none";
      this.logoutLinkShow = "inline-block";
      this.listLinkShow = "inline-block"
    }
  }

  ngOnInit(): void {
  }

}
