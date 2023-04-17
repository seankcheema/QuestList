import { Component } from '@angular/core';
import { DomSanitizer, SafeResourceUrl } from '@angular/platform-browser';

@Component({
  selector: 'app-docs',
  templateUrl: './docs.component.html',
  styleUrls: ['./docs.component.css']
})
export class DocsComponent {

  constructor(private sanitizer: DomSanitizer) { }

  signUpUrl: SafeResourceUrl | undefined;
  signInUrl: SafeResourceUrl | undefined;
  topGamesUrl: SafeResourceUrl | undefined;
  searchGamesUrl: SafeResourceUrl | undefined;

  ngOnInit() {

    this.signUpUrl = this.sanitizer.bypassSecurityTrustResourceUrl('../assets/doc-videos/sign-up.mp4');
    this.signInUrl = this.sanitizer.bypassSecurityTrustResourceUrl('../assets/doc-videos/sign-in.mp4');
    this.topGamesUrl = this.sanitizer.bypassSecurityTrustResourceUrl('../assets/doc-videos/top-games.mp4');
    this.searchGamesUrl = this.sanitizer.bypassSecurityTrustResourceUrl('../assets/doc-videos/game-search.mp4');
  }

}
