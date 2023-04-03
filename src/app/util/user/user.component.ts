import { Component } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { UserAuthService } from '../user-auth/user-auth.service';
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

  isLoggedIn: boolean = false;

  /**
   * Constuctor for UserComponent class
   * @param userService Injectable UserService that provides game data Observable
   */
  constructor(private userService: UserService, private userAuthService: UserAuthService, private route: ActivatedRoute) { }

  ngOnInit(): void {

    const username = this.route.snapshot.paramMap.get('username');

    if(username != null)
    this.isLoggedIn = this.userAuthService.isUserLoggedIn(username);
  }

}

