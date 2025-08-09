import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { Blog } from '../models/blog';

@Injectable({
  providedIn: 'root',
})
export class BlogService {
  private base = 'http://127.0.0.1:3000';
  constructor(private http: HttpClient) {}

  list(): Observable<Blog[]> {
    return this.http.get<Blog[]>(`${this.base}/blogs`);
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
}
