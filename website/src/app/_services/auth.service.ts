import { Injectable } from '@angular/core';
import { User } from '../model/user.model';
import {HttpClient, HttpHeaders} from '@angular/common/http';
import {Observable} from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class AuthService {
  BASE_URL = 'http://localhost:8080/api/user/';
  LOGIN = 'login';
  headers = new HttpHeaders({
    'Content-Type':  'application/json'
  });
  constructor(private http: HttpClient) {}
  login(user: User): Observable<User> {
    const url = `${this.BASE_URL}${this.LOGIN}`;
    return this.http.post<User>(url, JSON.stringify(user), { headers: this.headers });
  }

  logout() {
    sessionStorage.removeItem('account');
  }
}
