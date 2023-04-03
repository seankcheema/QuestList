import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class UserAuthService {

  constructor() { }

  public isUserLoggedIn(username: string): boolean {
    const user = sessionStorage.getItem('username');
    return (user === username);
  }

  public login(username: string): void {
    sessionStorage.setItem('username', username);
  }

}
