import { Component, OnInit } from '@angular/core';
import {ActivatedRoute, Router} from '@angular/router';
import {AuthService} from '../_services/auth.service';
import {CommandService} from '../_services/command.service';
import {Command, isValid} from '../_model/command.model';
import {Title} from '@angular/platform-browser';

@Component({
  selector: 'app-command-detail',
  templateUrl: './command-detail.component.html',
  styleUrls: ['./command-detail.component.css']
})
export class CommandDetailComponent implements OnInit {
  command: Command;
  errorMsg = '';
  constructor(private router: Router, private route: ActivatedRoute,
              private authService: AuthService, private commandService: CommandService, private titleService: Title) { }

  ngOnInit() {
    this.route.params.subscribe(params => {
      const id = params['id'];
      this.commandService.getCommandById(id)
        .subscribe(
          data => {
            this.command = data.body;
            this.titleService.setTitle(`Spongebot: Command ${this.command.id}`);
          },
          error => {
            this.authService.errorHandler(error.status);
          }
        );
    });
  }
  onSubmit() {
    if (!isValid(this.command)) {
      this.errorMsg = 'Regular expressions like "*", "^*" or "+" are not allowed. All fields are required!';
      return;
    } else {
      this.errorMsg = '';
    }
    this.commandService.updateCommand(this.command)
      .subscribe(
        data => {
          if (data.status === 204) {
            this.router.navigateByUrl('/dashboard/commands');
          }
        },
        error => {
          this.authService.errorHandler(error.status);
        }
      );
  }
}


