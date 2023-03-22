import { Component } from '@angular/core';
import { FormBuilder } from '@angular/forms';
import { UserService } from '../util/user/user.service';

@Component({
  selector: 'app-sign-in',
  templateUrl: './sign-in.component.html',
  styleUrls: ['./sign-in.component.css']
})
export class SignInComponent {

  constructor(private formBuilder:FormBuilder, private userService: UserService) { }
  
  // Form group for sign-up form
  profileForm = this.formBuilder.group({
    username:[''],
    password:[''],
  })

}
