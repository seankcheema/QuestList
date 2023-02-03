import { Component, Injectable } from '@angular/core'
import { HttpClient, HttpErrorResponse } from '@angular/common/http'

import { Observable, throwError } from 'rxjs'
import { catchError, retry } from 'rxjs/operators'

export interface Games {
    textFile: string;
}

@Component({
    selector: 'app-games',
    templateUrl: './games.component.html',
    providers: [GameService],
    styles: ['.error {color:red;}']
})

@Injectable()
export class GameService {
    gamesURL = 'http://localhost:8080/games';

    constructor(private http: HttpClient) { }

    getGames() {
        return this.http.get<GameService>('gamesURL');
    }

    private handleError(error: HttpErrorResponse) {
        if(error.status == 0)
            console.error('An error occurred:', error.error);
        else
            console.error(
                `Backend returned code ${error.status}, ` +
                `body was: ${error.error}`);

        return throwError(() => new Error('Something bad happened; please try again later.'));
    }

}