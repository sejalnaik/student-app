import { Routes } from '@angular/router';
import { HomeComponent } from '../components/home/home.component';
import { StudentCrudComponent } from '../components/student-crud/student-crud.component';
import { ErrorComponent } from "../components/error/error.component";
import { LoginComponent } from "../components/login/login.component";

export class RoutesClass {
    public static routes : Routes = [
        {path:"home", component:HomeComponent},
        {path:"list", component:StudentCrudComponent},
        {path:"login", component:LoginComponent},
        {path:"", redirectTo:"/home", pathMatch:"full"},
        { path: '**', component: ErrorComponent}
      ];
}
