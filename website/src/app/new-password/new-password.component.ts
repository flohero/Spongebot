import { Component, OnInit } from '@angular/core';
import {FormBuilder, Validators} from '@angular/forms';
import {Title} from '@angular/platform-browser';
import {AccountService} from '../_services/account.service';
import {AuthService} from '../_services/auth.service';
import {Router} from '@angular/router';

@Component({
  selector: 'app-new-password',
  templateUrl: './new-password.component.html',
  styleUrls: ['./new-password.component.css']
})
export class NewPasswordComponent implements OnInit {
  passwordMatch = true;
  accountForm = this.fb.group({
    password:   ['', Validators.required],
    passwordRe: ['', Validators.required]
  });
  constructor(private fb: FormBuilder, private titleService: Title,
              private accountService: AccountService, private auth: AuthService,
              private router: Router) { }

  ngOnInit() {
    this.titleService.setTitle(`Spongebot: Create Account`);
    this.accountForm.get('password').valueChanges.subscribe(
      () => {
        this.passwordMatches();
      }
    );
    this.accountForm.get('passwordRe').valueChanges.subscribe(
      () => {
        this.passwordMatches();
      }
    );
  }

  passwordMatches() {
    this.passwordMatch = this.accountForm.get('password').value === this.accountForm.get('passwordRe').value;
  }

  onSubmit() {
    this.accountService.updatePassword(this.accountForm.get('password').value)
      .subscribe(
        () => {
          this.router.navigateByUrl('/dashboard');
        },
        error => {
          this.auth.errorHandler(error.status);
        }
      );
  }
}
