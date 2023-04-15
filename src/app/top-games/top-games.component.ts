import { Component, OnInit } from '@angular/core';
import { Game, GameService } from '../util/game/game.service';
import { ActivatedRoute, Router } from '@angular/router';

@Component({
  selector: 'app-top-games',
  templateUrl: './top-games.component.html',
  styleUrls: ['./top-games.component.css'],
  providers: [GameService]
})

export class TopGamesComponent implements OnInit {
  //Columns to display in the table
  displayedColumns: string[] = ['name', 'rating'];
  //Data source for the table
  dataSource: Game[] | undefined;
  //Page counter
  pageCounter = 1;

  //Inject the GameService and ActivatedRoute
  constructor(private gameService: GameService, private route: ActivatedRoute, private router: Router) {}

  //Get the games from the GameService
  ngOnInit(): void {

    const page = this.route.snapshot.queryParamMap.get('page');

    this.gameService.getGames(page)
    .subscribe((data: Game[]) => this.dataSource = data );
  }

  forwardPage(): void {
    this.pageCounter++;

    this.router.navigate(['/top-games'], {queryParams: {page: this.pageCounter.toString()}, queryParamsHandling:'merge'});
    this.gameService.getGames(this.pageCounter.toString())
    .subscribe((data: Game[]) => this.dataSource = data );
  }

  previousPage(): void {
    this.pageCounter--;

    this.router.navigate(['/top-games'], {queryParams: {page: this.pageCounter.toString()}, queryParamsHandling: 'merge'});
    this.gameService.getGames(this.pageCounter.toString())
    .subscribe((data: Game[]) => this.dataSource = data );
  }
}