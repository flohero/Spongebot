import { Component, OnInit } from '@angular/core';
import {Account} from '../_model/account.model';
import {FormBuilder, Validators} from '@angular/forms';
import {Title} from '@angular/platform-browser';
import {AccountService} from '../_services/account.service';
import {AuthService} from '../_services/auth.service';
import {Router} from '@angular/router';

@Component({
  selector: 'app-account-create',
  templateUrl: './account-create.component.html',
  styleUrls: ['./account-create.component.css']
})
export class AccountCreateComponent implements OnInit {
  errorMsg = '';
  passwordMatch = false;
  accountForm = this.fb.group({
    username:   ['', Validators.required],
    password:   ['', Validators.required],
    passwordRe: ['', Validators.required],
    admin:      [false]
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

  onSubmit() {
    this.errorMsg = '';
    const account = new Account();
    account.username = this.accountForm.get('username').value;
    account.password = this.accountForm.get('password').value;
    account.admin = this.accountForm.get('admin').value;
    this.accountService.getAccountByUsername(account.username)
      .subscribe(
        () => {
          this.errorMsg = 'Username already in use';
        },
        error => {
          switch (error.status) {
            case 403: this.auth.errorHandler(403); break;
            case 404: this.createUser(account); break;
          }
        }
      );
  }

  createUser(account: Account) {
    this.accountService.createAccount(account)
      .subscribe(
        data => {
          console.log(data);
          if (data.status === 201) {
            this.router.navigateByUrl('/dashboard/accounts');
          }
        },
        error => {
          this.auth.errorHandler(error.status);
        }
      );
  }

  passwordMatches() {
    this.passwordMatch = this.accountForm.get('password').value === this.accountForm.get('passwordRe').value;
  }
}
