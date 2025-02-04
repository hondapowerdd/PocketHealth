import { Component, OnInit } from '@angular/core';
import { NgForm } from '@angular/forms';
import { Router } from '@angular/router';
import { UserService } from '../services/user.service';

@Component({
  selector: 'app-register',
  templateUrl: './register.component.html',
  styleUrls: ['./register.component.css']
})
export class RegisterComponent implements OnInit {
  constructor(
    private userService: UserService,
    private router: Router,
  ) { }

  ngOnInit(): void { }

  onFormSubmit(form: NgForm) {
    const name = form.value.name;
    const email = form.value.email;

    // modify to add favourite colour
    const favouriteColour = form.value.FavouriteColour.trim();
    // data validation for name
    const isValidName = /^[A-Za-z\s]+$/.test(name);
    if (!isValidName) {
      alert("Name should only contain letters and spaces. Please enter a valid name.");
      return;
    }
    // data validation for email
    const isValidEmail = /^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$/.test(email);
    if (!isValidEmail) {
      alert("Please enter a valid email address.");
      return;
    }
    // data validation for favourite colour
    const isValidColour = /^[A-Za-z\s]+$/.test(favouriteColour);
    if (!isValidColour) {
      alert("Favourite colour should only contain letters. Please enter a valid colour.");
      return; // Stop submission if the input is invalid
    }

    this.userService.postRegister(name, email, favouriteColour).subscribe((response: any) => {
      console.log("Response: ", response);
      if (response && response.user_id) {
        // modify this function to pass the user ID to the home component
        console.log("Name: ", name);
        console.log("User ID: ", response.user_id);
        this.router.navigate(['/home'], {
          queryParams: { name: name, userId: response.user_id, favouriteColour: favouriteColour }
        });
      } else {
        console.error("no user ID");
      }
    })
  }

}
