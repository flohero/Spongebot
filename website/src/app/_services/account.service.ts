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
  SINGLE_USER_URL = `${environment.api_url}user`;
  headers: HttpHeaders;
  constructor(private http: HttpClient, private auth: AuthService) {
    this.headers = environment.headers.append('Authorization', `Bearer ${this.auth.getToken()}`);
  }

  getAllAccounts(): Observable<HttpResponse<Account[]>> {
    return this.http.get<Account[]>(this.BASE_URL, {headers: this.headers, observe: 'response' });
  }

  getAccountByUsername(username: string): Observable<HttpResponse<Account[]>> {
    return this.http.get<Account[]>(`${this.BASE_URL}/username/${username}`, {headers: this.headers, observe: 'response' });
  }

  deleteAccountById(id: number): Observable<HttpResponse<{}>> {
    if (this.auth.getUserIdFromToken() === id) {
      this.auth.returnToLogin();
    }
    return this.http.delete(`${this.BASE_URL}/${id}/delete`, { headers: this.headers, observe: 'response'});
  }

  createAccount(account: Account): Observable<HttpResponse<Account>> {
    return this.http.post<Account>(`${this.SINGLE_USER_URL}/new`, account, { headers: this.headers, observe: 'response'});
  }

  updatePassword(password: string): Observable<HttpResponse<{}>> {
    return this.http.put<Account>(`${this.SINGLE_USER_URL}/update`,
      JSON.stringify({
      password: password
    }),
      { headers: this.headers, observe: 'response'});
  }
}
