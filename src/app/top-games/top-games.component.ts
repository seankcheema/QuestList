import { Component, OnInit } from '@angular/core';
import { Game, GameService } from '../util/game/game.service';
import { ActivatedRoute } from '@angular/router';

@Component({
  selector: 'app-top-games',
  templateUrl: './top-games.component.html',
  styleUrls: ['./top-games.component.css'],
  providers: [GameService]
})

export class TopGamesComponent implements OnInit {
  //Columns to display in the table
  displayedColumns: string[] = ['id', 'name', 'rating'];
  //Data source for the table
  dataSource: Game[] | undefined;

  //Inject the GameService and ActivatedRoute
  constructor(private gameService: GameService, private route: ActivatedRoute) {}

  //Get the games from the GameService
  ngOnInit(): void {

    const page = this.route.snapshot.queryParamMap.get('page');

    this.gameService.getGames(page)
    .subscribe((data: Game[]) => this.dataSource = data );
  }
}