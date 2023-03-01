import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { Game, GameService } from './game.service';

@Component({
  selector: 'app-game',
  templateUrl: './game.component.html',
  styleUrls: ['./game.component.css']
})

/**
 * GameComponent
 */
export class GameComponent implements OnInit {
  error: any;
  headers: string[] = [];
  games: Game[] | undefined;

  /**
   * Constuctor for GameComponent class
   * @param gameService Injectable GameService that provides game data Observable
   * @param route Used to get query parameters from URL
   */
  constructor(private gameService: GameService, private route: ActivatedRoute) { }

  /**
   * OnInit lifecycle hook
   */
  ngOnInit(): void {
    
    const page = this.route.snapshot.queryParamMap.get('page');

    this.showGames(page);
  }

  /**
   * Clears the games array and error message
   */
  clear() : void {
    this.games = undefined;
    this.error = undefined;
    this.headers = [];
  }

  /**
   * Subscribes to the gameService Observable and sets the games array to the data returned
   * @param page the page number to query in the back-end API
   */
  showGames(page: string | null) : void {
    this.gameService.getGames(page)
    .subscribe((data: Game[]) => this.games = data );
  }

}
