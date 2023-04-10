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
export class GameComponent {

  dataSource: Game[] | undefined;

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
    
    const gameSlug = this.route.snapshot.paramMap.get('game-slug');

    this.gameService.getGame(gameSlug).subscribe((data: Game[]) => { this.dataSource = data; });

  }

  

}
