import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders} from '@angular/common/http';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class ServicioService {

  headers=new HttpHeaders().set("Access-Control-Allow-Origin", "*");

  constructor(private http:HttpClient) { }

  Inicio(){
    return this.http.get(`http://localhost:8000`)
  }

  ListR(){
    return this.http.get(`http://localhost:8000/ListRep`)
  }

  Entrada(datos:any){
    return this.http.post(`http://localhost:8000/Entrada`,datos,{headers:this.headers});
  }
}
