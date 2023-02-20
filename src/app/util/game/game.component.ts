import { Component } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
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
  page: number = 1;

  constructor(private gameService: GameService, private route: ActivatedRoute) { }

  clear() {
    this.game = undefined;
    this.error = undefined;
    this.headers = [];
  }

  showGames(page: number) {
    this.gameService.getGames(page)
    .subscribe((data: Game) => this.game = { ...data });
  }

}
