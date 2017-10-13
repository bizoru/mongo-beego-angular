import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { MigaComponent } from './miga.component';

describe('MigaComponent', () => {
  let component: MigaComponent;
  let fixture: ComponentFixture<MigaComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ MigaComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(MigaComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should be created', () => {
    expect(component).toBeTruthy();
  });
});
