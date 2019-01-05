import { Injectable } from '@angular/core';
import { CanActivate } from '@angular/router';
import {AuthService} from '../_services/auth.service';
import { Location} from '@angular/common';

@Injectable({
  providedIn: 'root'
})
export class AdminGuard implements CanActivate {
  constructor(private auth: AuthService, private _location: Location) {}
  canActivate(): boolean {
    return this.auth.isAdmin();
  }
}
