import {Injectable} from '@angular/core';
import {User} from '../_model/user.model';
import {environment} from '../../environments/environment';
import {HttpClient} from '@angular/common/http';
import {Observable} from 'rxjs';
import {JwtHelperService} from '@auth0/angular-jwt';
import {Router} from '@angular/router';

@Injectable({
  providedIn: 'root'
})
export class AuthService {
  BASE_URL = `${environment.api_url}user/`;
  LOGIN = 'login';
  headers = environment.headers;
  helper = new JwtHelperService();

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

  tokenIsValid(): boolean {
    const token = this.getToken();
    if (!this.tokenAvailable(token)) {
      return false;
    }
    try {
      const expired = this.helper.decodeToken(token);
      const now = Math.floor(Date.now() / 1000);
      if (expired.Claims.exp <= now) {
        return false;
      }
    } catch (e) {
      return false;
    }
    return true;
  }

  isAdmin(): boolean {
    const token = this.getToken();
    if (!this.tokenAvailable(token)) {
      return false;
    }
    try {
      const tk = this.helper.decodeToken(token);
      return tk.Admin;
    } catch (e) {
      return false;
    }
  }

  errorHandler(statusCode: number) {
    if (statusCode === 403) {
      this.logout();
      this.returnToLogin();
    }
  }

  returnToLogin() {
    this.router.navigateByUrl('/login');
  }

  tokenAvailable(token: any) {
    return token;
  }

  getUserIdFromToken(): number {
    const token = this.getToken();
    if (!this.tokenAvailable(token)) {
      return -1;
    }
    try {
      const tk = this.helper.decodeToken(token);
      return tk.UserId;
    } catch (e) {
      return -1;
    }
  }
}
