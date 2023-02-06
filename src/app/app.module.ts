import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { MatButtonModule} from '@angular/material/button';
import { MatToolbarModule } from '@angular/material/toolbar';
import { HttpClientModule } from '@angular/common/http';
import { MatGridListModule } from '@angular/material/grid-list';
<<<<<<< HEAD
import { MatTabsModule } from '@angular/material/tabs';
=======
import { MatListModule } from '@angular/material/list';
import { MatCardModule } from '@angular/material/card';
>>>>>>> de5d00418163a50d5dc3bb79cc74eed7a5d07465

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { NavbarComponent } from './util/navbar/navbar.component';
import { AboutUsComponent } from './about-us/about-us.component';
import { SignUpComponent } from './sign-up/sign-up.component';
import { HomeComponent } from './home/home.component';
import { CommunityComponent } from './community/community.component';
import { TopGamesComponent } from './top-games/top-games.component';
import { GamesComponent } from './util/games/games.component';

@NgModule({
  declarations: [
    AppComponent,
    NavbarComponent,
    AboutUsComponent,
    SignUpComponent,
    HomeComponent,
    CommunityComponent,
    TopGamesComponent,
    GamesComponent
  ],
  imports: [
    BrowserModule,
    HttpClientModule,
    AppRoutingModule,
    BrowserAnimationsModule,
    MatButtonModule,
    MatToolbarModule,
    MatGridListModule,
<<<<<<< HEAD
    MatTabsModule
=======
    MatListModule,
    MatCardModule,
>>>>>>> de5d00418163a50d5dc3bb79cc74eed7a5d07465
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
