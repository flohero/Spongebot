import {HttpHeaders} from '@angular/common/http';

export const environment = {
  production: true,
  api_url: 'http://10.0.0.101:8080/api/',
  headers: new HttpHeaders({
    'Content-Type':  'application/json'
  })
};
