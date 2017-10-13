import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { CommonModule } from '@angular/common';
/*
import { AgendaNewComponent } from '../agenda/agenda-new/agenda-new.component';
import { AgendaComponent } from '../agenda/agenda-view/agenda-view.component';
import { AgendaEditComponent } from '../agenda/agenda-edit/agenda-edit.component';
*/
/*
const routes: Routes = [
  { path: 'agenda', component: AgendaComponent },
  { path: 'agenda/new', component: AgendaNewComponent },
  { path: 'agenda/edit/:id', component: AgendaEditComponent }
];
*/

@NgModule({
  imports: [
    CommonModule,
    RouterModule.forRoot(routes)
  ],
  exports: [RouterModule],
  declarations: []
})
export class RoutingModule { }
