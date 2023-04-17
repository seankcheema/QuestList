import { Component, OnInit } from '@angular/core';
import { Game, GameService } from '../util/game/game.service';
import { Observable } from 'rxjs';
import { Review, ReviewService } from '../util/review/review.service';

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.css']
})

export class HomeComponent implements OnInit {

  gameSource: Game[] | undefined;
  recentGameSource: Game[] | undefined;
  upcomingGameSource: Game[] | undefined;
  topRatedGameSource: Game[] | undefined;
  featuredGameSource: Game | undefined;
  recentReviewSource: Review[] | undefined;

  constructor(private gameService: GameService, private reviewService: ReviewService) { }

  ngOnInit(): void { 

    this.gameService.getGames(null)
    .subscribe((data: Game[]) => this.gameSource = data );

    this.gameService.getRecentGames()
    .subscribe((data: Game[]) => this.recentGameSource = data );

    this.gameService.getUpcomingGames()
    .subscribe((data: Game[]) => this.upcomingGameSource = data );

    this.gameService.getTopRatedGames()
    .subscribe((data: Game[]) => this.topRatedGameSource = data );

    this.gameService.getFeaturedGame()
    .subscribe((data: Game) => this.featuredGameSource = data );

    this.reviewService.getRecentReviews()
    .subscribe((data: Review[]) => this.recentReviewSource = data );
  }

}
