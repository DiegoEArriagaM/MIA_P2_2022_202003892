import { Component, OnInit } from '@angular/core';
import { ServicioService } from '../Coneccion/servicio.service';
import {Router} from '@angular/router';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {

  

  entrada:any={
    comando:'',
    idParticion:'',
    usuario:'',
    password:'',
    idU: 0,
	  idG:0,
	  idMoun:" ",
	  nombreU:" ",
    login:false
  }

  constructor(private route:Router,private servicios:ServicioService) { 
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
    console.log(this.entrada)
    console.log(sessionStorage)
  }

  ngOnInit(): void {
  }

  actualizarU(datos:any){
    this.entrada.idU=datos['id_u']
    this.entrada.idG=datos['id_g']
    this.entrada.idMoun=datos['id_mount']
    this.entrada.nombreU=datos['nombre_u']
    this.entrada.login=datos['login']
    sessionStorage['idU']=datos['id_u']
    sessionStorage['idG']=datos['id_g']
    sessionStorage['idMoun']=datos['id_mount']
    sessionStorage['nombreU']=datos['nombre_u']
    sessionStorage['login']=datos['login']
    console.log(this.entrada)
    console.log(sessionStorage)
  }

  mandarDatos(){
    this.entrada.comando="login -usuario="+this.entrada.usuario+" -password="+this.entrada.password+" -id="+this.entrada.idParticion
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
    this.route.navigate([''])
  }

}
