import { TestBed } from '@angular/core/testing';
import { HttpClientTestingModule, HttpTestingController } from '@angular/common/http/testing';
import { RouterTestingModule } from '@angular/router/testing';
import { GameComponent } from "./game.component";
import { Game, GameService } from './game.service';

/**
 * GameComponent Cypress test
 */
describe('GameComponent', () => {
    
    // Before each test, set up the TestBed
    beforeEach(() => TestBed.configureTestingModule({
        imports: [HttpClientTestingModule, RouterTestingModule],
        providers: [GameComponent, GameService]
    }));

    // Test that the component mounts
    it('mounts', () => {
        cy.mount(GameComponent);
    })

    // Test that the component can communicate with the back-end API
    it('should get games', () => {
        cy.request('http://localhost:8080/allGames/1').as('games');
        cy.get('@games').its('status').should('eq', 200);
    });
});