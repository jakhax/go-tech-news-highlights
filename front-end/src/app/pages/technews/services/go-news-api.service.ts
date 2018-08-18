import { Injectable } from '@angular/core';
import {HttpClient} from  '@angular/common/http'
import {environment} from "../../../../environments/environment";
import {Source} from "../data/source"
import {Article} from "../data/article"

@Injectable({
  providedIn: 'root'
})
export class GoNewsApiService {

  sources:any=[]
  articles: any=[]
  addSources(a){
    this.sources=[]
    for(var i=0; i<a["Sources"].length; i++){
      this.sources.push(new Source(a["Sources"][i]["Id"],a["Sources"][i]["Name"],a["Sources"][i]["Description"],a["Sources"][i]["Country"]))
      }
  }

  getSources(){
    var sourcesLink:string=environment.apiEndPoint+"/sources"
    let promise=new Promise((resolve,reject)=>{
      this.http.get(sourcesLink).toPromise().then(myResponse=>{
        this.addSources(myResponse)
        resolve()
      },
    error=>{
      console.log(error)
      reject()
    })
    })
    return promise   
  }


  addArticles(a){
    this.articles=[]
    for(var i=0; i<a["Articles"].length; i++){
      this.articles.push(new Article(a["Articles"][i]["Title"],a["Articles"][i]["Author"], a["Articles"][i]["Description"],
      a["Articles"][i]["Url"],a["Articles"][i]["UrlToImage"],a["Articles"][i]["PublishedAt"],a["Articles"][i]["Source"]["Name"]))
      }
  }

  getArticles(article_id){
    var sourcesLink:string=environment.apiEndPoint+"/articles/"+article_id
    let promise=new Promise((resolve,reject)=>{
      this.http.get(sourcesLink).toPromise().then(myResponse=>{
        this.addArticles(myResponse)
        resolve()
      },
    error=>{
      console.log(error)
      reject()
    })
    })
    return promise   
  }

  constructor(private http:HttpClient) { }
}
