import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import {LoginComponent} from './login/login.component';
import {DashboardComponent} from './dashboard/dashboard.component';
import {CommandsComponent} from './commands/commands.component';
import {RouterGuard} from './_guards/router.guard';
import {CommandDetailComponent} from './command-detail/command-detail.component';
import {CommandCreateComponent} from './command-create/command-create.component';

const routes: Routes = [
  { path: 'login', component: LoginComponent },
  { path: '', redirectTo: '/login', pathMatch: 'full'},
  {
    path: 'dashboard',
    component: DashboardComponent,
    canActivate: [RouterGuard],
    children: [
      { path: 'commands', component: CommandsComponent, canActivate: [RouterGuard] },
      { path: '', redirectTo: '/dashboard/commands', pathMatch: 'full'},
      { path: 'commands/create', component: CommandCreateComponent, canActivate: [RouterGuard]},
      { path: 'commands/:id', component: CommandDetailComponent, canActivate: [RouterGuard]}
    ]
  }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
