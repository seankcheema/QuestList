import { Component, OnInit } from '@angular/core';
import { Game, GameService } from '../util/game/game.service';

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.css']
})

export class HomeComponent implements OnInit {

  dataSource: Game[] | undefined;

  constructor(private gameService: GameService) { }

  ngOnInit(): void { 
    this.gameService.getRecentGames()
    .subscribe((data: Game[]) => this.dataSource = data );
  }

}
