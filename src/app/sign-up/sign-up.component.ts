import { Component, OnInit } from '@angular/core';
import { FormBuilder } from '@angular/forms'
import { User, UserService } from '../util/user/user.service';
import { UserAuthService } from '../util/user-auth/user-auth.service';
import { Router } from '@angular/router';

@Component({
  selector: 'app-sign-up',
  templateUrl: './sign-up.component.html',
  styleUrls: ['./sign-up.component.css'],
  providers: [{provide: Window, useValue: window}]
})
export class SignUpComponent{

  /**
   * Constructor for SignUpComponent class
   * @param formBuilder FormBuilder used to create the sign-up form
   */
  constructor(private formBuilder:FormBuilder, private userService: UserService, private userAuthService: UserAuthService, private window: Window) { }
  
  // Form group for sign-up form
  profileForm = this.formBuilder.group({
    username:[''],
    email:[''],
    password:[''],
  })

  /**
   * Posts a new user to the back-end API
   * @param username Username of the new user
   * @param password Password of the new user
   * @param email Email of the new user
   */
  addUser(username: string, email:string, password: string) : void {
    this.userService.addUser({username, email, password} as User)
    .subscribe({ 
      next: () => {
        this.userAuthService.login(username);
        this.window.location.href = '/home';
      },
      error: () => {
        alert('This user already exists');
      }
    });
  }
}
