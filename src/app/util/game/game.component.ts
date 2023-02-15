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
  game: Game | undefined;

  constructor(private gameService: GameService) { }

  clear() {
    this.game = undefined;
    this.error = undefined;
    this.headers = [];
  }

  showGames() {
    this.gameService.getGames()
    .subscribe((data: Game) => this.game = { ...data });
  }
}
