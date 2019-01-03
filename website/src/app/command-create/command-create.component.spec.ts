import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { CommandCreateComponent } from './command-create.component';

describe('CommandCreateComponent', () => {
  let component: CommandCreateComponent;
  let fixture: ComponentFixture<CommandCreateComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ CommandCreateComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(CommandCreateComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
