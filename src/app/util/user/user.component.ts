import { Component } from '@angular/core';
import { FormBuilder, FormControl, FormGroup } from '@angular/forms';
import { User, UserService } from './user.service';

@Component({
  selector: 'app-user',
  templateUrl: './user.component.html',
  styleUrls: ['./user.component.css']
})

/**
 * UserComponent
 */
export class UserComponent {

  /**
   * Constuctor for UserComponent class
   * @param userService Injectable UserService that provides game data Observable
   */
  constructor(private userService: UserService, private formBuilder:FormBuilder) { }

  profileForm: FormGroup = this.formBuilder.group({
    username: [''],
    password: ['']
  });

  /**
   * Posts a new user to the back-end API
   * @param username Username of the new user
   * @param password Password of the new user
   */
  addUser(username: string, password: string) : void {

    this.userService.addUser({username, password} as User);

  }
}

