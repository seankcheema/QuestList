import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { Game, GameService } from './game.service';

@Component({
  selector: 'app-game',
  templateUrl: './game.component.html',
  styleUrls: ['./game.component.css']
})

export class GameComponent implements OnInit {
  error: any;
  headers: string[] = [];
  games: Game[] | undefined;

  constructor(private gameService: GameService, private route: ActivatedRoute) { }

  ngOnInit(): void {
    
    const page = this.route.snapshot.queryParamMap.get('page');

    this.showGames(page);
  }

  clear() {
    this.games = undefined;
    this.error = undefined;
    this.headers = [];
  }

  showGames(page: string | null) {
    this.gameService.getGames(page)
    .subscribe((data: Game[]) => this.games = data );
  }

}
