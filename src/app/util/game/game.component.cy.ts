import { TestBed } from '@angular/core/testing';
import { HttpClientTestingModule, HttpTestingController } from '@angular/common/http/testing';
import { RouterTestingModule } from '@angular/router/testing';
import { GameComponent } from "./game.component";
import { Game, GameService } from './game.service';

describe('GameComponent', () => {

    let gameService: GameService;

    beforeEach(() => TestBed.configureTestingModule({
        imports: [HttpClientTestingModule, RouterTestingModule],
        providers: [GameComponent, GameService]
    }));

    it('mounts', () => {
        cy.mount(GameComponent);
    })

    it('should get games', () => {

        cy.request('http://localhost:8080/allGames/1').as('games');

        cy.get('@games').its('status').should('eq', 200);

    });
});