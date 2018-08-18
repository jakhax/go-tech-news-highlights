import { RouterModule, Routes } from '@angular/router';
import { NgModule } from '@angular/core';

import { PagesComponent } from './pages.component';
import { TechNewsComponent } from './technews/technews.component';
import { ArticleComponent } from './article/article.component';
import { NotFoundComponent } from './miscellaneous/not-found/not-found.component';

const routes: Routes = [{
  path: '',
  component: PagesComponent,
  children: [
    {
      path: 'technews',
      component: TechNewsComponent,
    },
    {
      path: 'articles/:id',
      component: ArticleComponent,
    },
    {
      path: '',
      redirectTo: 'technews',
      pathMatch: 'full',
    },
    {
      path: '**',
      component: NotFoundComponent,
    }
  ],
}];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule],
})
export class PagesRoutingModule {
}
