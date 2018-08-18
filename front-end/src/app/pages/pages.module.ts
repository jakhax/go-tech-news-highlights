import { NgModule } from '@angular/core';

import { PagesComponent } from './pages.component';
import { TechNewsModule } from './technews/technews.module';
import { PagesRoutingModule } from './pages-routing.module';
import { ThemeModule } from '../@theme/theme.module';
import { MiscellaneousModule } from './miscellaneous/miscellaneous.module';
import { ArticleComponent } from './article/article.component';
import {NgxPaginationModule} from 'ngx-pagination';



const PAGES_COMPONENTS = [
  PagesComponent,
];

@NgModule({
  imports: [
    PagesRoutingModule,
    ThemeModule,
    TechNewsModule,
    MiscellaneousModule,
    NgxPaginationModule,
  ],
  declarations: [
    ...PAGES_COMPONENTS,
    ArticleComponent,
  ],
})
export class PagesModule {
}
