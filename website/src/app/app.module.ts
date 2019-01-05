import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { HttpClientModule } from '@angular/common/http';
import { FormsModule } from '@angular/forms';
import { AppRoutingModule } from './app-routing.module';
import { ReactiveFormsModule } from '@angular/forms';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { AppComponent } from './app.component';
import { CommandsComponent } from './commands/commands.component';
import { LoginComponent } from './login/login.component';
import { DashboardComponent } from './dashboard/dashboard.component';
import { SidebarComponent } from './sidebar/sidebar.component';
import { CommandDetailComponent } from './command-detail/command-detail.component';
import { CommandCreateComponent } from './command-create/command-create.component';
import { AccountsComponent } from './accounts/accounts.component';
import { AccountCreateComponent } from './account-create/account-create.component';
import { NewPasswordComponent } from './new-password/new-password.component';

@NgModule({
  declarations: [
    AppComponent,
    CommandsComponent,
    LoginComponent,
    DashboardComponent,
    SidebarComponent,
    CommandDetailComponent,
    CommandCreateComponent,
    AccountsComponent,
    AccountCreateComponent,
    NewPasswordComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    HttpClientModule,
    FormsModule,
    ReactiveFormsModule,
    BrowserAnimationsModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
