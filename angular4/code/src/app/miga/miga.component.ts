import { Component, OnInit } from '@angular/core';
import { GlobalsComponent } from '../globals/globals.component';

@Component({
  selector: 'app-miga',
  templateUrl: './miga.component.html',
  styleUrls: ['./miga.component.css']
})
export class MigaComponent implements OnInit {

  currentModule: string;

  constructor(private globals: GlobalsComponent) {

    this.currentModule = globals.currentModule;
  }

  ngOnInit() {
  }


}
