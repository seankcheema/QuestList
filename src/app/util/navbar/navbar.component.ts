import { Component } from '@angular/core';
import { UserAuthService } from '../user-auth/user-auth.service';

@Component({
  selector: 'app-navbar',
  templateUrl: './navbar.component.html',
  styleUrls: ['./navbar.component.css']
})
export class NavbarComponent {

  isLoggedIn: boolean = false;

  constructor(private userAuthService: UserAuthService) {}

  ngOnInit(): void {

    this.isLoggedIn = this.userAuthService.isUserLoggedIn(sessionStorage.getItem('username') || '');
  }

}
