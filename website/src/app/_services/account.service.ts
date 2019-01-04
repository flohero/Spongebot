import { Injectable } from '@angular/core';
import { environment } from '../../environments/environment';
import {HttpClient, HttpHeaders, HttpResponse} from '@angular/common/http';
import {AuthService} from './auth.service';
import {Observable} from 'rxjs';
import { Account } from '../_model/account.model';

@Injectable({
  providedIn: 'root'
})
export class AccountService {
  BASE_URL = `${environment.api_url}users`;
  headers: HttpHeaders;
  constructor(private http: HttpClient, private auth: AuthService) {
    this.headers = environment.headers.append('Authorization', `Bearer ${this.auth.getToken()}`);
  }

  getAllAccounts(): Observable<HttpResponse<Account[]>> {
    return this.http.get<Account[]>(this.BASE_URL, {headers: this.headers, observe: 'response' });
  }

  deleteAccountById(id: number): Observable<HttpResponse<{}>> {
    if (this.auth.getUserIdFromToken() === id) {
      this.auth.returnToLogin();
    }
    return this.http.delete(`${this.BASE_URL}/${id}/delete`, { headers: this.headers, observe: 'response'});
  }
}
