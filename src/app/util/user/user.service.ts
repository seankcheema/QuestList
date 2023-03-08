import { HttpClient, HttpErrorResponse } from '@angular/common/http';
import { HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { catchError, Observable, retry, throwError } from 'rxjs';

const httpOptions = {
  headers: new HttpHeaders({
    'Content-Type': 'application/json'
  })
}

/**
 * User interface
 */
export interface User {
  username: string;
  password: string;
}

/**
 * UserService injectable
 */
@Injectable({
  providedIn: 'root'
})
export class UserService {

  usersUrl: string = 'http://localhost:8080/sign-up'; //URL to back-end API

  constructor(private http : HttpClient) { }

  /**
   * Adds a new user to the back-end API
   * @param user User to add
   * @returns Observable of User - can be subscribed to and used to update the UI
   */
  addUser(user: User): Observable<User> {
    console.log("addUser: " + user.username + " " + user.password);
    return this.http.post<User>(this.usersUrl, user, httpOptions)
    .pipe(
      retry(3),
      catchError(this.handleError)
    );
  }

  /**
   * Handles errors from the back-end API
   * @param error HttpErrorResponse
   * @returns an observable with a user-facing error message 
   */
  private handleError(error: HttpErrorResponse) {
    if (error.error instanceof ErrorEvent) {
      // A client-side or network error occurred. Handle it accordingly.
      console.error('An error occurred:', error.error.message);
    } else {
      // The backend returned an unsuccessful response code.
      // The response body may contain clues as to what went wrong,
      console.error(
        `Backend returned code ${error.status}, ` +
        `body was: ${error.error}`);
    }
    // return an observable with a user-facing error message
    return throwError(() => new Error('Something bad happened; please try again later.'));
  }
}
