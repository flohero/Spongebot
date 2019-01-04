import { Component, OnInit } from '@angular/core';
import {AccountService} from '../_services/account.service';
import { Account} from '../_model/account.model';
import {AuthService} from '../_services/auth.service';

@Component({
  selector: 'app-accounts',
  templateUrl: './accounts.component.html',
  styleUrls: ['./accounts.component.css']
})
export class AccountsComponent implements OnInit {
  accounts: Account[];
  constructor(private accountService: AccountService, private auth: AuthService) { }

  ngOnInit() {
    this.loadAccounts();
  }

  loadAccounts() {
    this.accountService.getAllAccounts()
      .subscribe(
        (data) => {
          this.accounts = data.body.sort((a, b) => a.id - b.id);
        },
        error => {
          this.auth.errorHandler(error.status);
        }
      );
  }

  onDelete(id: number) {
    // TODO
    this.accountService.deleteAccountById(id)
      .subscribe(
        (data) => {
          this.loadAccounts();
        },
        error => {
          this.auth.errorHandler(error.status);
        }
      );
  }
}
