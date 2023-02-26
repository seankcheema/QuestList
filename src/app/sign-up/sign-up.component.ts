import { Component, OnInit } from '@angular/core';
import { FormBuilder } from '@angular/forms'

@Component({
  selector: 'app-sign-up',
  templateUrl: './sign-up.component.html',
  styleUrls: ['./sign-up.component.css']
})
export class SignUpComponent{
  constructor(private formBuilder:FormBuilder){}
  
  profileForm = this.formBuilder.group({
    username:[''],
    email:[''],
    password:[''],
    dob:['']
  })

  signup() {
    console.log('Form data is ', this.profileForm.value)
  }
}
