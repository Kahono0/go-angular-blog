import {
  Component,
  OnInit,
  ViewChild,
  ElementRef,
  AfterViewInit,
} from '@angular/core';
import { BlogService } from '../../services/blog.service';
import { Blog } from '../../models/blog';
import { RouterLink } from '@angular/router';
import { DatePipe } from '@angular/common';
import { Subject, debounceTime, distinctUntilChanged } from 'rxjs';

@Component({
  selector: 'app-blog-search',
  standalone: true,
  imports: [RouterLink, DatePipe],
  templateUrl: './blog-search.component.html',
  styleUrl: './blog-search.component.scss',
})
export class BlogSearchComponent implements AfterViewInit {
  @ViewChild('searchInput') searchInputRef!: ElementRef<HTMLInputElement>;
  blogs: Blog[] = [];
  query: string = '';
  loading: boolean = false;
  error: string | null = null;

  private searchSubject = new Subject<string>();

  constructor(private blogService: BlogService) {
    this.searchSubject
      .pipe(debounceTime(300), distinctUntilChanged())
      .subscribe((query) => {
        this.query = query;
        this.searchBlogs();
      });
  }

  ngAfterViewInit(): void {
    this.searchInputRef.nativeElement.focus();
  }

  onSearchChange(value: string) {
    this.searchSubject.next(value);
  }

  searchBlogs(): void {
    if (!this.query.trim()) {
      this.blogs = [];
      return;
    }
    this.loading = true;
    this.blogService.search(this.query).subscribe({
      next: (blogs) => {
        this.blogs = blogs;
        this.loading = false;
        this.error = null;
      },
      error: (err) => {
        this.error = 'Failed to search blogs';
        this.loading = false;
        this.blogs = [];
      }
    });
  }
}
