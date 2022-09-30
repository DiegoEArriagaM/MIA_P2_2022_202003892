import { Component, OnInit } from '@angular/core';
import { ServicioService } from '../Coneccion/servicio.service';
import {Router} from '@angular/router';
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

  execE:any={
    comandos:null,
    idU: 0,
	  idG:0,
	  idMoun:" ",
	  nombreU:" ",
    login:false,
    i:0
  }

  terminal:any
  resp:any=""
  t:any=1
  list:any
  

  constructor(private route:Router,private servicios:ServicioService) {
    if(sessionStorage['idU']==undefined){
      this.entrada.idU=0
      this.entrada.idG=0
      this.entrada.idMoun=" "
      this.entrada.nombreU=" "
      this.entrada.login=false
      this.execE.idU=0
      this.execE.idG=0
      this.execE.idMoun=" "
      this.execE.nombreU=" "
      this.execE.login=false
    }else{
      this.entrada.idU=sessionStorage['idU']
      this.entrada.idG=sessionStorage['idG']
      this.entrada.idMoun=sessionStorage['idMoun']
      this.entrada.nombreU=sessionStorage['nombreU']
      this.entrada.login=sessionStorage['login']
      this.execE.idU=sessionStorage['idU']
      this.execE.idG=sessionStorage['idG']
      this.execE.idMoun=sessionStorage['idMoun']
      this.execE.nombreU=sessionStorage['nombreU']
      this.execE.login=sessionStorage['login']
      console.log(this.entrada)
    }
    this.terminal=""
    this.inicio=""
    this.servicios.Inicio().subscribe(
      data=>{
        let datos:any=data;
        this.inicio=datos
        this.actualizarU(datos.usuario)
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

  actualizarU(datos:any){
    this.entrada.idU=datos['id_u']
    this.entrada.idG=datos['id_g']
    this.entrada.idMoun=datos['id_mount']
    this.entrada.nombreU=datos['nombre_u']
    this.entrada.login=datos['login']
    this.execE.idU=datos['id_u']
    this.execE.idG=datos['id_g']
    this.execE.idMoun=datos['id_mount']
    this.execE.nombreU=datos['nombre_u']
    this.execE.login=datos['login']
    sessionStorage['idU']=datos['id_u']
    sessionStorage['idG']=datos['id_g']
    sessionStorage['idMoun']=datos['id_mount']
    sessionStorage['nombreU']=datos['nombre_u']
    sessionStorage['login']=datos['login']
  }

  mandarComando(){
    const r=this.t
    this.entrada.comando=this.limpiar(this.entrada.comando)
    this.servicios.Entrada(this.entrada).subscribe(
      data=>{
        let datos:any=data
        this.actualizarU(datos.usuario)
        this.terminal+=this.t+")"+this.entrada.comando+"\n"
        if(datos.res!=""){
          this.resp+=r+")"+datos.res+"\n"
        }
        this.entrada.comando=""
        this.t=this.t+1
        this.entrada.comando=""
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
    const r=this.t
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
            this.terminal+=this.t+")"+lines[j]+"\n"
            lines[j]=this.limpiar(lines[j])
            this.t++
            j=j+1
          }
          this.execE.comandos=lines
          this.execE.i=r
          let datos:any
          this.servicios.Exec(this.execE).subscribe(
            data=>{
              datos=data
              this.resp+=datos.res
            },
            err=>{
              console.log(err)
            }
          )
          
        }
      }else{
        alert("NO SE SELECCIONO UN ARCHIVO .script")
      }
    }else{
      alert("NO SE SELECCIONO UN ARCHIVO .script")
    }
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

  getRep(){
    this.servicios.ListR().subscribe(
      data=>{
        this.list=data
        console.log(this.list)
      },
      err=>{
        console.log(err)
      }
    )
    console.log(this.list)
  }

  verRep(name:any){
    const dir="http://localhost:8000/Reportes/"+name
        window.open(dir)
  }

  irLogin(){
    this.route.navigate(['login'])
  }

  logout(){
    this.entrada.comando="logout"
    this.servicios.Entrada(this.entrada).subscribe(
      data=>{
        let datos:any=data
        alert(datos.res)
        this.actualizarU(datos.usuario)
      },
      err=>{
        console.log(err)
      },() => {
        // 'onCompleted' callback.
        // No errors, route to new page here
      }
    )
  }
}