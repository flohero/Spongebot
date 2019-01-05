import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import {LoginComponent} from './login/login.component';
import {DashboardComponent} from './dashboard/dashboard.component';
import {CommandsComponent} from './commands/commands.component';
import {RouterGuard} from './_guards/router.guard';
import {CommandDetailComponent} from './command-detail/command-detail.component';
import {CommandCreateComponent} from './command-create/command-create.component';
import {AccountsComponent} from './accounts/accounts.component';
import {AdminGuard} from './_guards/admin.guard';
import {AccountCreateComponent} from './account-create/account-create.component';
import {NewPasswordComponent} from './new-password/new-password.component';

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
      { path: 'commands/:id', component: CommandDetailComponent, canActivate: [RouterGuard]},
      { path: 'accounts', component: AccountsComponent, canActivate: [AdminGuard, RouterGuard]},
      { path: 'accounts/create', component: AccountCreateComponent, canActivate: [AdminGuard, RouterGuard]},
      { path: 'accounts/update', component: NewPasswordComponent, canActivate: [RouterGuard]}
    ]
  }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
