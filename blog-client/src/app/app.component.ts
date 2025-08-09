import { Component } from '@angular/core';
import { RouterLink, RouterLinkActive, RouterOutlet } from '@angular/router';
import { HomeComponent } from './home/home.component';
import { HeaderComponent } from './components/header/header.component';
import { FooterComponent } from './components/footer/footer.component';

@Component({
  selector: 'app-root',
  standalone: true,
  imports: [RouterOutlet, HeaderComponent, FooterComponent],
  template: `
    <main>
      <app-header></app-header>
      <section>
        <router-outlet></router-outlet>
      </section>
      <app-footer></app-footer>
    </main>
  `,
  styleUrl: './app.component.scss',
})
export class AppComponent {
  title = 'Angulat blog app';
}
