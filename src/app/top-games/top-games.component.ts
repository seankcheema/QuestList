import { Component } from '@angular/core';
import { GameComponent } from '../util/game/game.component';

@Component({
  selector: 'app-top-games',
  templateUrl: './top-games.component.html',
  styleUrls: ['./top-games.component.css']
})
export class TopGamesComponent extends GameComponent {

  

  ngOnInit()
  {
    this.showGames(1);
  }
}
