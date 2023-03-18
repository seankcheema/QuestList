import { Component, AfterViewInit, ViewChild, OnInit } from '@angular/core';
import {MatTableDataSource} from '@angular/material/table';
import {MatPaginator} from '@angular/material/paginator';
import { GameService } from '../util/game/game.service';
import { Game } from '../util/game/game.service';
import {CollectionViewer, DataSource} from '@angular/cdk/collections';
import { Observable } from 'rxjs';

@Component({
  selector: 'app-top-games',
  templateUrl: './top-games.component.html',
  styleUrls: ['./top-games.component.css'],
  providers: [GameService]
})

export class TopGamesComponent implements OnInit{
  data = new TableDataSource(this.web);
  displayedColumns: string[] = ['id', 'name', 'rating'];


  constructor(private web : GameService) {}

  ngOnInit(){}
}

export class TableDataSource extends DataSource<any> {
  constructor(private web : GameService){
    super();
  }
  connect(): Observable<Game[]> {
    return this.web.getGames(null);
    // need to find argument 
  }
  disconnect() {}
}

/*

first attempt for inserting games into table 

export class TopGamesComponent implements OnInit{
  data: Game[] = [];
  displayedColumns: string[] = ['id', 'name', 'rating'];


  constructor(private web : GameService) {
    this.web.getGames().subscribe(x => {
      this.data = x;
      console.log(this.data);
    })
  }
}


original (pre games into table)

export class TopGamesComponent implements AfterViewInit{
  displayedColumns: string[] = ['position', 'name', 'score'];
  dataSource = new MatTableDataSource<TopGames>(topgames_data);

  @ViewChild(MatPaginator) paginator:any = MatPaginator;

  ngAfterViewInit() {
    this.dataSource.paginator = this.paginator;
}
}


export interface TopGames {
  name: string;
  position: number;
  score: number;
}

const topgames_data: TopGames[]= [
  {position: 1, name: 'Best game of all time', score: 100},
  {position: 2, name: 'Minecraft', score : 99},
  {position: 3, name: 'COD', score: 98},
  {position: 4, name: 'Halo', score : 97},
  {position: 5, name: 'WOW', score: 90},
  {position: 6, name: 'Overwatch', score : 80},
  {position: 7, name: 'Fortnite', score: 79},
  {position: 8, name: 'Mario', score : 70},
  {position: 9, name: 'Zelda', score: 69},
  {position: 10, name: 'Apex', score : 60},
  {position: 11, name: 'PUBG', score: 50},
  {position: 12, name: 'Skyrim', score : 1},
];

*/


