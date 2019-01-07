import {HttpHeaders} from '@angular/common/http';

export const environment = {
  production: true,
  api_url: 'http://localhost:8080/api/',
  headers: new HttpHeaders({
    'Content-Type':  'application/json'
  })
};
