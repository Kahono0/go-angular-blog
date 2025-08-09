import { Routes } from '@angular/router';
import { HomeComponent } from './home/home.component';
import { BlogViewComponent } from './components/blog-view/blog-view.component';
import { BlogCreateComponent } from './components/blog-create/blog-create.component';
import { BlogSearchComponent } from './components/blog-search/blog-search.component';

export const routes: Routes = [
    {path: '', pathMatch: 'full', component: HomeComponent},
    {path: 'blog/:slug', component: BlogViewComponent},
    {path: 'create', component: BlogCreateComponent},
    {path: 'search', component: BlogSearchComponent},
    {path: '**', redirectTo: '', pathMatch: 'full'},
];
