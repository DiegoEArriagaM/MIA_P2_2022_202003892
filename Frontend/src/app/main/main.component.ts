import { Component, OnInit } from '@angular/core';
import { ServicioService } from '../Coneccion/servicio.service';
//import {graphviz} from 'd3-graphviz';

@Component({
  selector: 'app-main',
  templateUrl: './main.component.html',
  styleUrls: ['./main.component.css']
})
export class MainComponent implements OnInit {
  //JSON
  inicio:any

  entrada:any={
    comando:''
  }

  terminal:any

  

  constructor(private servicios:ServicioService) {
    this.terminal=""
    this.inicio=""
    this.servicios.Inicio().subscribe(
      data=>{
        let datos:any=data;
        this.inicio=datos
      },
      err=>{
        this.inicio={
          res:"No se logro conectar con el servidor"
        }
        console.log(err)
      }
    )
    
   }

  

  ngOnInit(): void {
    //graphviz("#graph").renderDot('digraph {a -> b}');
  }

  mandarComando(){
    this.terminal+=this.entrada.comando+"\n"
    this.limpiar()
    const command=this.getComand();
    this.servicios.Entrada(this.entrada).subscribe(
      data=>{
        
        let datos:any=data
        this.terminal+=datos.res+"\n"
        this.entrada.comando=""
        /*const dir="http://localhost:8000/Reportes/"+"ejemplo.pdf"
        window.open(dir)*/
      },
      err=>{
        console.log(err)
      }
    )
  }

  onFileSelected(event:Event){
    const targe=event.target as HTMLInputElement
    let doc:any=targe.files as FileList
    //Se informa que archivo se guardo
    const name:string[]=doc[0].name.split(".")
    if (name.length>1){
      if (name[1]=="mia"){
        const file: File = doc.item(0)
        const reader: FileReader = new FileReader();
        reader.readAsText(file);
        reader.onload=(e)=>{
          const content:string=reader.result as string
          const lines: string[] = content.split('\n');
          for(let line of lines){
            this.exec(line)
          }
        }
      }else{
        alert("NO SE SELECCIONO UN ARCHIVO .mia")
      }
    }else{
      alert("NO SE SELECCIONO UN ARCHIVO .mia")
    }
     
  }

  exec(comando:any){
    this.entrada.comando=comando
    this.terminal+=this.entrada.comando+"\n"
    this.limpiar()
    const command=this.getComand();
    this.servicios.Entrada(this.entrada).subscribe(
      data=>{
        let datos:any=data
        this.terminal+=datos.res
        this.entrada.comando=""
      },
      err=>{
        console.log(err)
      }
    )
  }

  limpiar(){
    const str=this.entrada.comando
    this.entrada.comando=str.replaceAll('\n',' ')
    this.entrada.comando=str.replaceAll('\t',' ')
    this.entrada.comando=str.replaceAll('\r',' ')

    while(this.entrada.comando.charAt(0)==' '){
      this.entrada.comando=this.entrada.comando.replace(' ','')
    }
  }

  getComand():any{
    let comand=""
    let str=this.entrada.comando
    while(str.charAt(0)!=' ' && str.length>0){
      comand+=str.slice(0,1)
      str=str.replace(str.slice(0,1),'');
    }
    return comand
  }

}
