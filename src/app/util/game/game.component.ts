import { Component } from '@angular/core';
import { Game, GameService } from './game.service';

@Component({
  selector: 'app-game',
  templateUrl: './game.component.html',
  styleUrls: ['./game.component.css']
})
export class GameComponent {
  error: any;
  headers: string[] = [];
  games: Game[] | undefined;

  constructor(private gameService: GameService) { this.showGames() }

  clear() {
    this.games = undefined;
    this.error = undefined;
    this.headers = [];
  }

  showGames() {
    this.gameService.getGames()
    .subscribe((data: Game[]) => this.games = data );
  }

}
