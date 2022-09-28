import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { MainComponent } from './main/main.component';
import { ReportesComponent } from './reportes/reportes.component';

const routes: Routes = [
  {
    path:'',
    component:MainComponent,
    pathMatch:'full'
  },
  {
    path:'Reportes',
    component:ReportesComponent,
  }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }