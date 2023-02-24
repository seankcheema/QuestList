import { TestBed } from '@angular/core/testing';
import { HttpClientTestingModule, HttpTestingController } from '@angular/common/http/testing';
import { RouterTestingModule } from '@angular/router/testing';
import { GameComponent } from "./game.component";
import { GameService } from './game.service';

describe('GameComponent', () => {

    beforeEach(() => TestBed.configureTestingModule({
        imports: [HttpClientTestingModule, RouterTestingModule],
        providers: [GameService]
    }));

    it('mounts', () => {
        cy.mount(GameComponent);
    })

    it('should get games', () => {

    });
});