import {Component,OnInit} from '@angular/core';
import {GoNewsApiService} from "./services/go-news-api.service"
import { Router} from '@angular/router';

@Component({
  selector: 'ngx-technews',
  templateUrl: './technews.component.html',
  styleUrls: ['./technews.component.scss'],
 
})
export class TechNewsComponent implements OnInit {

  urlToImage:string ="https://i.imgur.com/th5yZb8.jpg"
 

  constructor(public getNewsSources:GoNewsApiService,private router:Router) {
  
  }
  viewArticle(id){
    this.router.navigate(["/pages/articles",id])
  }
  ngOnInit(){
    this.getNewsSources.getSources().then(()=>{
      console.log(this.getNewsSources.sources)
    })
  }
 
}
