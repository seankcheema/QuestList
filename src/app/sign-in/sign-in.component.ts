import { Component } from '@angular/core';
import { FormBuilder } from '@angular/forms';
import { User, UserService } from '../util/user/user.service';
import { Router } from '@angular/router';

@Component({
  selector: 'app-sign-in',
  templateUrl: './sign-in.component.html',
  styleUrls: ['./sign-in.component.css'],
  providers: [{provide: Window, useValue: window}]
})
export class SignInComponent {

  constructor(private formBuilder:FormBuilder, private userService: UserService, private window: Window) { }
  
  // Form group for sign-up form
  profileForm = this.formBuilder.group({
    username:[''],
    password:[''],
  })

  // Posts a new user to the back-end API
  checkUser(username: string, password: string) : void {
    this.userService.findUser({username, password} as User)
    .subscribe({ 
        next: () => {
          sessionStorage.setItem('username', username);
          this.window.location.href = '/home';
        },
        error: () => {
          alert('Invalid username or password');
        }
    });
  }

}
