import { Injectable } from '@angular/core';
import { User } from '../model/user.model';
import { environment } from '../../environments/environment';
import { HttpClient } from '@angular/common/http';
import {Observable} from 'rxjs';
import { JwtHelperService } from '@auth0/angular-jwt';
import {Router} from '@angular/router';

@Injectable({
  providedIn: 'root'
})
export class AuthService {
  BASE_URL = `${environment.api_url}user/`;
  LOGIN = 'login';
  headers = environment.headers;
  constructor(private http: HttpClient, private router: Router) {}
  login(user: User): Observable<User> {
    const url = `${this.BASE_URL}${this.LOGIN}`;
    return this.http.post<User>(url, JSON.stringify(user), { headers: this.headers });
  }

  logout() {
    sessionStorage.removeItem('account');
  }
  getToken(): string {
    return sessionStorage.getItem('account');
  }

  isValid(): boolean {
    const helper = new JwtHelperService();
    const token = this.getToken();
    if (!token) {
      return false;
    }
    try {
      const expired = helper.decodeToken(token);
      const now = Math.floor(Date.now() / 1000);
      if (expired.Claims.exp <= now) {
        return false;
      }
    } catch (e) {
      return false;
    }
    return true;
  }
  errorHandler(statusCode: number) {
    if (statusCode === 403) {
      this.logout();
      this.router.navigateByUrl('/login');
    }
  }
}
