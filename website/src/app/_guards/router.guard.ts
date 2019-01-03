import { Injectable } from '@angular/core';
import {CanActivate, Router} from '@angular/router';
import {AuthService} from '../_services/auth.service';

@Injectable({
  providedIn: 'root'
})
export class RouterGuard implements CanActivate {
  constructor(private router: Router, private auth: AuthService) {}

  canActivate(): boolean {
    if (!this.auth.isValid()) {
      this.router.navigateByUrl('/login');
      return false;
    }
    return true;
  }
}
