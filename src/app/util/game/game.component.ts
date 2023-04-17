import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { Game, GameService } from './game.service';
import { UserAuthService } from '../user-auth/user-auth.service';
import { FormBuilder, FormControl } from '@angular/forms';
import { Review, ReviewService } from '../review/review.service';

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
  isLoggedIn: boolean = false;

  /**
   * Constuctor for GameComponent class
   * @param gameService Injectable GameService that provides game data Observable
   * @param route Used to get query parameters from URL
   */
  constructor(private formBuilder: FormBuilder, private gameService: GameService, private route: ActivatedRoute, private userAuthService: UserAuthService, private reviewService: ReviewService) { }

  // Form group for sign-up form
  ratingControl = new FormControl();

  /**
   * OnInit lifecycle hook
   */
  ngOnInit(): void {
    
    const gameSlug = this.route.snapshot.paramMap.get('game-slug');

    this.gameService.getGame(gameSlug).subscribe((data: Game[]) => { this.dataSource = data; });

    this.isLoggedIn = this.userAuthService.isUserLoggedIn(sessionStorage.getItem('username') || '');

    this.ratingControl.setValue(5);

  }

  addReview(reviewDescription: string): void {

    if(this.isLoggedIn === false) {
      window.location.href = '/sign-in';
    }
    else {
      const gameSlug = this.route.snapshot.paramMap.get('game-slug');
      const username = sessionStorage.getItem('username') || '';
      const rating = this.ratingControl.value;

      console.log("addReview: " + gameSlug + ' ' + username + ' ' + rating + ' ' + reviewDescription);

      this.reviewService.addReview({GameSlug: gameSlug, Username: username, Rating: rating, Description: reviewDescription})
    }
  }

}
