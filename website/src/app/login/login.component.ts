import { Component, OnInit } from '@angular/core';
import {AuthService} from '../services/auth.service';
import {User} from '../model/user.model';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {
  user = new User();
  constructor(private authService: AuthService) { }

  ngOnInit() {
  }
  onSubmit() {
    this.authService.login(this.user).subscribe(
      data => {
        sessionStorage.setItem('account', data.token);
        console.log('logged in!');
        console.log(data);
      },
      error => {
        console.log(error);
      }
    );
  }
}
