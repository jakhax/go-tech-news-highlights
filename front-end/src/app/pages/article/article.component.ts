import { Component, OnInit } from '@angular/core';
import {ActivatedRoute,Router} from "@angular/router";
import {GoNewsApiService} from "../technews/services/go-news-api.service"


@Component({
  selector: 'article',
  templateUrl: './article.component.html',
  styleUrls: ['./article.component.scss']
})
export class ArticleComponent implements OnInit {

  p:number;

  goHome(){
    this.router.navigate(["/pages/technews"])
  }
  pageChanged(p){
    console.log(p)
  }
  constructor(private route: ActivatedRoute, public getArticle:GoNewsApiService,private router:Router) {
    this.route.params.subscribe( params => this.getArticle.getArticles(params["id"]).then(()=>{
      console.log(this.getArticle.articles)
    }) );
}

  ngOnInit() {
  }

}
