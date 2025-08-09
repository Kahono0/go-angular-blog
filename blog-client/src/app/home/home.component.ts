import { Component, OnInit } from '@angular/core';
import { Blog } from '../models/blog';
import { BlogService } from '../services/blog.service';
import { DatePipe } from '@angular/common';
import { RouterLink } from '@angular/router';

@Component({
  selector: 'app-home',
  standalone: true,
  imports: [DatePipe, RouterLink],
  templateUrl: './home.component.html',
  styleUrl: './home.component.scss',
})
export class HomeComponent implements OnInit {
  blogs: Blog[] = [];
  totalPages: number = 0;
  currentPage: number = 1;
  itemsPerPage: number = 10;
  loading = false;
  error: string | null = null;

  windowSize: number = 3;

  constructor(private blogService: BlogService) {}

  ngOnInit(): void {
    this.loadBlogs();
  }

  changePage(page: number): void {
    if (page < 1 || page > this.totalPages) return;
    this.currentPage = page;
    this.loadBlogs();
  }

  loadBlogs(): void {
    this.loading = true;
    this.blogService.list(this.currentPage, this.itemsPerPage).subscribe({
      next: (response) => {
        this.blogs = response.data;
        this.totalPages = response.totalPages;
        this.currentPage = response.currentPage;
        this.loading = false;

        setTimeout(() => {
          window.scrollTo({ top: 0, behavior: 'smooth' });
        });
      },
      error: (err) => {
        this.error = 'Failed to load blogs';
        this.loading = false;
      },
    });
  }

  get toShowPages(): number[] {
    const pages: number[] = [];
    const startPage = Math.max(
      1,
      this.currentPage - Math.floor(this.windowSize / 2)
    );
    const endPage = Math.min(this.totalPages, startPage + this.windowSize - 1);

    for (let i = startPage; i <= endPage; i++) {
      pages.push(i);
    }

    return pages;
  }
}
