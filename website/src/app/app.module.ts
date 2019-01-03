import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { HttpClientModule } from '@angular/common/http';
import { FormsModule } from '@angular/forms';
import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { CommandsComponent } from './commands/commands.component';
import { LoginComponent } from './login/login.component';
import { DashboardComponent } from './dashboard/dashboard.component';
import { SidebarComponent } from './sidebar/sidebar.component';
import { CommandDetailComponent } from './command-detail/command-detail.component';
import { CommandCreateComponent } from './command-create/command-create.component';

@NgModule({
  declarations: [
    AppComponent,
    CommandsComponent,
    LoginComponent,
    DashboardComponent,
    SidebarComponent,
    CommandDetailComponent,
    CommandCreateComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    HttpClientModule,
    FormsModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
