import { Component, OnInit } from '@angular/core';
import { FormBuilder } from '@angular/forms'

@Component({
  selector: 'app-sign-up',
  templateUrl: './sign-up.component.html',
  styleUrls: ['./sign-up.component.css']
})
export class SignUpComponent{

  /**
   * Constructor for SignUpComponent class
   * @param formBuilder FormBuilder used to create the sign-up form
   */
  constructor(private formBuilder:FormBuilder) { }
  
  // Form group for sign-up form
  profileForm = this.formBuilder.group({
    username:[''],
    email:[''],
    password:[''],
    dob:['']
  })

  //TODO: Send form data to back-end API
  signup() : void {
    console.log('Form data is ', this.profileForm.value)
  }
}
