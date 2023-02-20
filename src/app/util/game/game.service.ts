import { Injectable } from '@angular/core'
import { HttpClient } from '@angular/common/http'
import { HttpErrorResponse } from '@angular/common/http'

import { Observable, throwError } from 'rxjs'
import { catchError, retry } from 'rxjs/operators'

export interface Game {
    id: number;
    slug: string;
    name: string;
    released: string;
    tba: boolean;
    background_image: string;
    rating: number;
    rating_top: number;
    ratings: any[];
    ratings_count: number;
    reviews_text_count: number;
    added: number;
    added_by_status: any;
    metacritic: number;
    playtime: number;
    suggestions_count: number;
    updated: string;
    esrb_rating: any;
    platforms: any[];
}

@Injectable()
export class GameService {

    gamesUrl: string = 'http://localhost:8080/allGames/'

    constructor(private http : HttpClient) { }

    getGames(page: number) {

        this.gamesUrl = this.gamesUrl + page;

        console.log(this.gamesUrl)

        return this.http.get<Game>(this.gamesUrl)
        .pipe(
            retry(3),
            catchError(this.handleError)
        );
    }

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