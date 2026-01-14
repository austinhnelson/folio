import { Component, OnInit } from '@angular/core';
import { Api } from '../../core/services/api';
import { User } from '../../shared/models/user';

@Component({
  selector: 'app-register',
  imports: [],
  templateUrl: './register.html',
  styleUrl: './register.css',
})
export class Register implements OnInit {
  constructor(private api: Api) {}

  ngOnInit() {
    this.registerUser();
  }

  registerUser() {
    const newUser: User = {
      username: 'newuser',
      password: 'password123',
      email: 'test@test.com'
    };

    this.api.register(newUser).subscribe({
      next: (data) => console.log(data),
      error: (err) => console.error('Error registering user:', err)
    });
  }
}
