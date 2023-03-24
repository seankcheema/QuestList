import { HttpClientModule } from '@angular/common/http';
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { RouterTestingModule } from '@angular/router/testing';
import { AppRoutingModule } from '../app-routing.module';
import { GameService } from '../util/game/game.service';

import { TopGamesComponent } from './top-games.component';

describe('TopGamesComponent', () => {
  let component: TopGamesComponent;
  let fixture: ComponentFixture<TopGamesComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [HttpClientModule, RouterTestingModule],
      declarations: [ TopGamesComponent ],
      providers: [GameService]
    })
    .compileComponents();

    fixture = TestBed.createComponent(TopGamesComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
