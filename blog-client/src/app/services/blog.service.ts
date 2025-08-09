import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { Blog, BlogResponse } from '../models/blog';
import { environment } from '../../environments/environment';

@Injectable({
  providedIn: 'root',
})
export class BlogService {
  private base = environment.baseUrl;
  constructor(private http: HttpClient) {}

  list(page: number, itemsPerPage: number): Observable<BlogResponse> {
    return this.http.get<BlogResponse>(`${this.base}/blogs`, {
      params: {
        page: page.toString(),
        itemsPerPage: itemsPerPage.toString(),
      },
    });
  }

  get(slug: string): Observable<Blog> {
    return this.http.get<Blog>(`${this.base}/blogs/${slug}`);
  }

  createBlog(blogData: any): Observable<any> {
    const formData = new FormData();
    formData.append('title', blogData.title);
    formData.append('createdAt', blogData.createdAt);
    formData.append('slug', blogData.slug);
    formData.append('content', blogData.content);
    if (blogData.image) {
      formData.append('image', blogData.image);
    }

    return this.http.post(`${this.base}/blog`, formData);
  }

  search(query: string): Observable<Blog[]> {
    return this.http.get<Blog[]>(`${this.base}/blogs/search`, {
      params: { q: query },
    });
  }
}
