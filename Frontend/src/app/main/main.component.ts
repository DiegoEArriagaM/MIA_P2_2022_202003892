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
  doc:any
  nameDoc:any=""

  entrada:any={
    comando:'',
    idU: 0,
	  idG:0,
	  idMoun:" ",
	  nombreU:" ",
    login:false

  }

  terminal:any
  resp:any=""
  t:any=1
  

  

  constructor(private servicios:ServicioService) {
    if(sessionStorage['idU']==undefined){
      this.entrada.idU=0
      this.entrada.idG=0
      this.entrada.idMoun=" "
      this.entrada.nombreU=" "
      this.entrada.login=false
    }else{
      this.entrada.idU=sessionStorage['idU']
      this.entrada.idG=sessionStorage['idG']
      this.entrada.idMoun=sessionStorage['idMoun']
      this.entrada.nombreU=sessionStorage['nombreU']
      this.entrada.login=sessionStorage['login']
    }
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
    const r=this.t
    this.entrada.comando=this.limpiar(this.entrada.comando)
    const command=this.getComand();
    this.servicios.Entrada(this.entrada).subscribe(
      data=>{
        let datos:any=data
        this.terminal+=this.t+")"+this.entrada.comando+"\n"
        if(datos.res!=""){
          this.resp+=r+")"+datos.res+"\n"
        }
        this.entrada.comando=""
        this.t=this.t+1
        this.entrada.comando=""
        /*const dir="http://localhost:8000/Reportes/"+"ejemplo.pdf"
        window.open(dir)*/
      },
      err=>{
        console.log(err)
      },() => {
        // 'onCompleted' callback.
        // No errors, route to new page here
      }
    )
  }

  onFileSelected(event:Event){
    const targe=event.target as HTMLInputElement
    this.doc=targe.files as FileList
    this.nameDoc=this.doc[0].name
  }

  ejecutar(){
    //Se informa que archivo se guardo
    const name:string[]=this.doc[0].name.split(".")
    if (name.length>1){
      if (name[1]=="script"){
        const file: File = this.doc.item(0)
        const reader: FileReader = new FileReader();
        reader.readAsText(file);
        reader.onload=(e)=>{
          const content:string=reader.result as string
          const lines: string[] = content.split('\n');
          let j=0
          while(j<lines.length){
            this.exec(lines[j])
            j=j+1
          }
        }
      }else{
        alert("NO SE SELECCIONO UN ARCHIVO .script")
      }
    }else{
      alert("NO SE SELECCIONO UN ARCHIVO .script")
    }
  }

  exec(comando:any){
    const r=this.t
    this.terminal+=this.t+")"+comando+"\n"
    this.t=this.t+1
    console.log(comando)
    this.entrada.comando=comando
    this.entrada.comando=this.limpiar(this.entrada.comando)
    const command=this.getComand();
    let datos:any
    this.servicios.Entrada(this.entrada).subscribe(
      data=>{
        datos=data
        if(datos.res!=""){
          this.resp+=r+")"+datos.res+"\n"
        }
        this.entrada.comando=""
      },
      err=>{
        console.log(err)
      }
    )
  }

  limpiar(comando:any):string{
    let str=comando
    str=str.replaceAll('\n',' ')
    str=str.replaceAll('\t',' ')
    str=str.replaceAll('\r',' ')

    while(this.entrada.comando.charAt(0)==' '){
      this.entrada.comando=this.entrada.comando.replace(' ','')
    }
    return str
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
