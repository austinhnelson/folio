import { Component, OnInit } from '@angular/core';
import { Api } from '../../core/services/api';

@Component({
  templateUrl: './login.html',
})
export class Login implements OnInit {
  constructor(private api: Api) {}

  ngOnInit() {
    this.loginUser();
  }

  private loginUser() {
    const credentials = {
      email: "test@test.com",
      password: "password123"
    };

    this.api.login(credentials).subscribe({
      error: (err) => console.error("Error logging in user:", err)
    });
  }
}
