import { NgModule } from '@angular/core';
import { ThemeModule } from '../../@theme/theme.module';
import { TechNewsComponent } from './technews.component';



@NgModule({
  imports: [
    ThemeModule,
  ],
  declarations: [
    TechNewsComponent,
  ],
})
export class TechNewsModule { }
