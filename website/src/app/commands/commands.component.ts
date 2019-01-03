import { Component, OnInit } from '@angular/core';
import {CommandService} from '../_services/command.service';
import {Command} from '../model/command.model';
import {AuthService} from '../_services/auth.service';
import {Router} from '@angular/router';
import {Title} from '@angular/platform-browser';

@Component({
  selector: 'app-commands',
  templateUrl: './commands.component.html',
  styleUrls: ['./commands.component.css']
})
export class CommandsComponent implements OnInit {
  commands: Command[];
  constructor(private commandsService: CommandService, private router: Router,
              private authService: AuthService, private titleService: Title) { }

  ngOnInit() {
    this.titleService.setTitle('Spongebot: Commands');
    this.loadCommands();
  }
  onDelete(id: number) {
    this.commandsService.deleteCommand(id)
      .subscribe(
        data => {
          if (data.status === 204) {
            this.loadCommands();
          }
        },
        error => {
          this.authService.errorHandler(error.status);
        }
      );
  }
  loadCommands() {
    this.commandsService.getAllCommands()
      .subscribe(
        data => {
          this.commands = data.body.sort((a, b) => a.id - b.id);
        },
        error => {
          this.authService.errorHandler(error.status);
          console.error(
            `Backend returned code ${error.status}, ` +
            `body was: ${error.error}`);
        }
      );
  }
}
