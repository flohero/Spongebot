import { Injectable } from '@angular/core';
import {HttpClient, HttpHeaders, HttpResponse} from '@angular/common/http';
import { environment } from '../../environments/environment';
import {Command} from '../model/command.model';
import {AuthService} from './auth.service';
import {Observable} from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class CommandService {
  BASE_URL = `${environment.api_url}commands`;
  headers: HttpHeaders;
  constructor(private http: HttpClient, private authService: AuthService) {
    this.headers = environment.headers.append('Authorization', `Bearer ${this.authService.getToken()}`);
  }

  getAllCommands(): Observable<HttpResponse<Command[]>> {
    return this.http.get<Command[]>(this.BASE_URL, { headers: this.headers, observe: 'response' });
  }

  getCommandById(id: number): Observable<HttpResponse<Command>> {
    return this.http.get<Command>(`${this.BASE_URL}/${id}`, { headers: this.headers, observe: 'response' });
  }

  updateCommand(cmd: Command): Observable<HttpResponse<{}>> {
    return this.http.put(`${this.BASE_URL}/${cmd.id}/update`, JSON.stringify(cmd), {headers: this.headers, observe: 'response'});
  }

  deleteCommand(id: number): Observable<HttpResponse<{}>> {
    return this.http.delete(`${this.BASE_URL}/${id}/delete`, { headers: this.headers, observe: 'response'});
  }

  createCommand(command: Command): Observable<HttpResponse<Command>> {
    return this.http.post<Command>(`${this.BASE_URL}/new`, JSON.stringify(command), { headers: this.headers, observe: 'response'});
  }
}
