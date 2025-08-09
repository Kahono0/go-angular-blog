import { Component, Input, OnInit } from '@angular/core';
import { RouterLink } from '@angular/router';
import { BlogService } from '../../services/blog.service';
import { Blog } from '../../models/blog';
import { DatePipe } from '@angular/common';

@Component({
  selector: 'app-blog-view',
  standalone: true,
  imports: [RouterLink, DatePipe],
  templateUrl: './blog-view.component.html',
  styleUrl: './blog-view.component.scss',
})
export class BlogViewComponent implements OnInit {
  @Input() slug!: string;
  blog: Blog | null = null;
  loading = false;
  error: string | null = null;
  constructor(private blogService: BlogService) {}

  ngOnInit(): void {
    console.log('BlogViewComponent initialized with slug:', this.slug);
    this.loadBlog();
  }

  loadBlog(): void {
    this.loading = true;
    this.blogService.get(this.slug).subscribe({
      next: (blog) => {
        this.blog = blog;
        this.loading = false;
      },
      error: (err) => {
        this.error = 'Failed to load blog';
        this.loading = false;
      },
    });
  }
}
