import { Component } from '@angular/core';
import { Games, GameService } from './games.service';

@Component({
  selector: 'app-games',
  templateUrl: './games.component.html',
  styleUrls: ['./games.component.css']
})

export class GamesComponent {
  games: Games | undefined;
  error: any;

  constructor(private gameService: GameService) { }

  showGames() {
    this.gameService.getGames()
      .subscribe((data: Games) => this.games = {...data });
  }
}
