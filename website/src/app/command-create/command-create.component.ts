import { Component, OnInit } from '@angular/core';
import {Command, isValid} from '../model/command.model';
import {ActivatedRoute, Router} from '@angular/router';
import {AuthService} from '../_services/auth.service';
import {CommandService} from '../_services/command.service';
import {Title} from '@angular/platform-browser';

@Component({
  selector: 'app-command-create',
  templateUrl: './command-create.component.html',
  styleUrls: ['./command-create.component.css']
})
export class CommandCreateComponent implements OnInit {
  command = new Command();
  errorMsg = '';
  constructor(private router: Router, private route: ActivatedRoute,
              private authService: AuthService, private commandService: CommandService, private titleService: Title) { }

  ngOnInit() {
    this.titleService.setTitle(`Spongebot: Create Command`);
  }
  onSubmit() {
    if (!isValid(this.command)) {
      this.errorMsg = 'Regular expressions like "*", "^*" or "+" are not allowed.';
      return;
    } else {
      this.errorMsg = '';
    }
    this.commandService.createCommand(this.command)
      .subscribe(
        data => {
          console.log(data);
          if (data.status === 201) {
            this.router.navigate(['/dashboard/commands/']);
          }
        },
        error => {
          this.authService.errorHandler(error.status);
        }
      );
  }

}
