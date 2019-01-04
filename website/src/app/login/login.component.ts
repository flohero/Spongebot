import { Component, OnInit } from '@angular/core';
import {AuthService} from '../_services/auth.service';
import {User} from '../_model/user.model';
import {Router} from '@angular/router';
import {Title} from '@angular/platform-browser';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {
  user = new User();
  error = false;
  errorMsg = '';
  constructor(private authService: AuthService, private router: Router, private titleService: Title) { }

  ngOnInit() {
    this.titleService.setTitle('Login to Spongebot');
  }
  onSubmit() {
    this.authService.login(this.user).subscribe(
      data => {
        sessionStorage.setItem('account', data.token);
        this.router.navigateByUrl('/dashboard');
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
