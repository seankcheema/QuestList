import { Component, OnInit } from '@angular/core';
import {FormControl} from '@angular/forms';
import {Observable} from 'rxjs';
import {map, startWith} from 'rxjs/operators';
import { Game, GameService } from '../game/game.service';
import { Router } from '@angular/router';


@Component({
  selector: 'app-search-bar',
  templateUrl: './search-bar.component.html',
  styleUrls: ['./search-bar.component.css'],
  providers: [{provide: Window, useValue: window}]
})

export class SearchBarComponent implements OnInit {

  dataSource: Game[] | undefined;
  myControl = new FormControl('');
  options: string[] = [];
  filteredOptions!: Observable<string[]>;

  constructor(private gameService: GameService, private window: Window) { }

  ngOnInit() {
    this.filteredOptions = this.myControl.valueChanges.pipe(
      startWith(''),
      map(value => this._filter(value || '')),
    );
  }

  private _filter(value: string): string[] {
    const filterValue = value.toLowerCase();

    this.gameService.getGame(filterValue).subscribe((data: Game[]) => { this.dataSource = data; });
    
    if(this.dataSource != undefined) {
      for(let i = 0; i < this.dataSource.length; i++){
        this.options[i] = this.dataSource[i].name;
      }
    }

    return this.options.filter(option => option.toLowerCase().includes(filterValue));
  }

  redirectToGame(game: string) {
    this.window.location.href = '/game/' + game;
  }

}
