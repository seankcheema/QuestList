import { Injectable } from '@angular/core'
import { HttpClient } from '@angular/common/http'
import { HttpErrorResponse } from '@angular/common/http'

import { Observable, throwError } from 'rxjs'
import { catchError, retry } from 'rxjs/operators'

/**
 * Game interface
 */
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

/**
 * GameService injectable
 */
@Injectable(
    {providedIn: 'root'}
)
export class GameService {

    gamesUrl: string = 'http://localhost:8080'; //URL to back-end API

    constructor(private http : HttpClient) { }

    /**
     * Gets a list of games from the back-end API
     * @param page specifies the page number to query in the back-end API
     * @returns an observable of type Game[]
     */
    getGames(page: string | null): Observable<Game[]>{
        return this.http.get<Game[]>(this.gamesUrl + "/games?page=" + (page || '1'))
        .pipe(
            retry(3),
            catchError(this.handleError)
        );
    }

    /**
     * Gets a list of recent games from the back-end API
     * @returns an observable of type Game[]
     */
    getRecentGames(): Observable<Game[]> {
        return this.http.get<Game[]>(this.gamesUrl + "/recent")
        .pipe(
            retry(3),
            catchError(this.handleError)
        )
    }

    /**
     * Gets an array of games closest matching slug from the back-end API
     * @param slug Specifies the slug of the game to get from the back-end API
     * @returns an observable of type Game[]
     */
    getGame(slug: string | null): Observable<Game[]> {
        return this.http.get<Game[]>(this.gamesUrl + "/specific-game?slug=" + (slug || ''))
        .pipe(
            retry(3),
            catchError(this.handleError)
        )
    }

    /**
     * Gets a list of upcoming games from the back-end API
     * @returns an observable of type Game[]
     */
    getUpcomingGames(): Observable<Game[]> {
        return this.http.get<Game[]>(this.gamesUrl + "/upcominggames")
        .pipe(
            retry(3),
            catchError(this.handleError)
        )
    }

    /**
     * Gets a list of top-rated games from the back-end API
     * @returns an obserable of type Game[]
     */
    getTopRatedGames(): Observable<Game[]> {
        return this.http.get<Game[]>(this.gamesUrl + "/topgames")
        .pipe(
            retry(3),
            catchError(this.handleError)
        )
    }

    /**
     * Gets a list of featured games from the back-end
     * @returns an observable of type Game[]
     */
    getFeaturedGame(): Observable<Game> {
        return this.http.get<Game>(this.gamesUrl + "/featuredgame")
        .pipe(
            retry(3),
            catchError(this.handleError)
        )
    }

    /**
     * Handles an error with getting games from the back-end API
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