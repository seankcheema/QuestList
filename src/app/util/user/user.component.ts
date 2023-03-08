import { Component } from '@angular/core';
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
  constructor(private userService: UserService) { }

  /**
   * Posts a new user to the back-end API
   * @param username Username to add
   * @param password Password to add
   */
  addUser(username: string, password: string) : void {

  this.userService.addUser({username, password} as User);

  }
}

