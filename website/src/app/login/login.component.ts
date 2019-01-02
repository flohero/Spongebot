import { Component, OnInit } from '@angular/core';
import {AuthService} from '../_services/auth.service';
import {User} from '../model/user.model';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {
  user = new User();
  error = false;
  errorMsg = '';
  constructor(private authService: AuthService) { }

  ngOnInit() {
  }
  onSubmit() {
    this.authService.login(this.user).subscribe(
      data => {
        sessionStorage.setItem('account', data.token);
      },
      error => {
        console.log(error);
        if (error.error.message) {
          this.errorMsg = error.error.message;
        } else {
          this.errorMsg = error.error;
        }
      }
    );
  }
}
