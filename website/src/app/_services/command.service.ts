import { Injectable } from '@angular/core';
import {HttpClient} from '@angular/common/http';

@Injectable({
  providedIn: 'root'
})
export class CommandService {

  constructor(http: HttpClient) { }

  getAllCommands() {

  }
}
