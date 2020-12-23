import { Directive, HostListener, Self} from '@angular/core';
import { NgControl } from '@angular/forms';

@Directive({
  selector: '[appEmptyToNullDirectve]'
})
export class EmptyToNullDirectveDirective {

  constructor(@Self() private ngControl:NgControl) { }
  @HostListener('keyup', ['$event']) onKeyDowns(event: KeyboardEvent) {
    if (this.ngControl.value?.trim() === '') {
      alert("Inside directive" + this.ngControl.value)
      this.ngControl.reset(null);
    }
  }
}
