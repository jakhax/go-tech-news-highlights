import { Component } from '@angular/core';

@Component({
  selector: 'ngx-footer',
  styleUrls: ['./footer.component.scss'],
  template: `
    <span class="created-by">Created by <b><a href="https://github.com/jakhax" target="_blank">Jakhax</a></b> 2018</span>
  `,
})
export class FooterComponent {
}
