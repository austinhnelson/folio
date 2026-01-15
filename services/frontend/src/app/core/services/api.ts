import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { User } from '../../shared/models/user';

@Injectable({
  providedIn: 'root',
})
export class Api {
  private baseUrl: string = "http://localhost:8080";

  constructor(private http: HttpClient) { }

  register(user: User): Observable<any> {
    return this.http.post(`${this.baseUrl}/auth/register`, user, { withCredentials: true });
  }

  login (user: Pick<User, 'email' | 'password'>): Observable<any> {
    return this.http.post(`${this.baseUrl}/auth/login`, user, { withCredentials: true });
  }
}
