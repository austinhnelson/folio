import { Component, OnInit } from '@angular/core';
import { Api } from '../../core/services/api';
import { User } from '../../shared/models/user';

@Component({
  templateUrl: './register.html',
})
export class Register implements OnInit {
  constructor(private api: Api) {}

  ngOnInit() {
    this.registerUser();
  }

  private registerUser() {
    const newUser: User = {
      username: "newuser",
      password: "password123",
      email: "test@test.com"
    };

    this.api.register(newUser).subscribe({
      error: (err) => console.error("Error registering user:", err)
    });
  }
}
