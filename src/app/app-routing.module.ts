import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { AboutUsComponent } from './about-us/about-us.component';
import { CommunityComponent } from './community/community.component';
import { HomeComponent } from './home/home.component';
import { SignInComponent } from './sign-in/sign-in.component';
import { SignUpComponent } from './sign-up/sign-up.component';
import { TopGamesComponent } from './top-games/top-games.component';
import { UserComponent } from './util/user/user.component';
import { GameComponent } from './util/game/game.component';

// Routes for the application
const routes: Routes = [
  { path: '', redirectTo: '/home', pathMatch: 'full'},
  { path: 'home', component: HomeComponent},
  { path: 'top-games', component: TopGamesComponent},
  { path: 'community', component: CommunityComponent},
  { path: 'about-us', component: AboutUsComponent},
  { path: 'sign-up', component: SignUpComponent},
  { path: 'sign-in', component: SignInComponent},
  { path: 'user/:username', component: UserComponent },
  { path: 'game/:game-slug', component: GameComponent }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
