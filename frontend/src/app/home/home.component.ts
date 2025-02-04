import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.css']
})
export class HomeComponent implements OnInit {
  name: string | null = '';
  userId: string | null = '';
  favouriteColour: string | null = '';

  constructor(private route: ActivatedRoute) { }

  ngOnInit(): void {
    this.name = this.route.snapshot.queryParamMap.get('name');
    this.userId = this.route.snapshot.queryParamMap.get('userId');
    this.favouriteColour = this.route.snapshot.queryParamMap.get('favouriteColour');
  }

}
