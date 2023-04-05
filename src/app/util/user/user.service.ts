import { HttpClient, HttpErrorResponse } from '@angular/common/http';
import { HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable, throwError } from 'rxjs';

const httpOptions = {
  headers: new HttpHeaders({
    'Content-Type': 'application/json',
  })
}

/**
 * User interface
 */
export interface User {
  username: string;
  email: string;
  password: string;
}

/**
 * UserService injectable
 */
@Injectable({
  providedIn: 'root'
})
export class UserService {

  baseUrl: string = 'http://localhost:8080'; //URL to back-end API

  constructor(private http : HttpClient) { }

  /**
   * Adds a new user to the back-end API
   * @param user User to add
   * @returns Observable of User - can be subscribed to and used to update the UI
   */
  addUser(user: User): Observable<User> {

    const signUpURL:string = this.baseUrl + '/sign-up';

    console.log("addUser: " + user.username + ' ' + user.email + ' ' + user.password + ' '+ signUpURL);

    return this.http.post<User>(signUpURL, user, httpOptions);
  }

  /**
   * Finds if the user exists in the back-end API
   * @param user User to find
   * @returns Observable of User - can be subscribed to and used to update the UI
   */
  findUser(user:User): Observable<User> {
      
      const signInURL:string = this.baseUrl + '/sign-in';
  
      console.log("findUser: " + user.username + ' ' + user.password + ' '+ signInURL);
  
      return this.http.post<User>(signInURL, user, httpOptions);
  }

  /**
   * Handles errors from the back-end API
   * @param error HttpErrorResponse
   * @returns an observable with a user-facing error message 
   */
  private handleError(error: HttpErrorResponse) {
    if (error.status === 0) {
      // A client-side or network error occurred. Handle it accordingly.
      console.error('An error occurred:', error.error);
    } else {
      // The backend returned an unsuccessful response code.
      // The response body may contain clues as to what went wrong.
      console.error(
        `Backend returned code ${error.status}, body was: `, error.error);
    }
    // Return an observable with a user-facing error message.
    return throwError(() => new Error('Something bad happened; please try again later.'));
  }
}
