import { Component } from '@angular/core';
import {
  FormBuilder,
  FormGroup,
  FormsModule,
  ReactiveFormsModule,
  Validators,
} from '@angular/forms';
import { BlogService } from '../../services/blog.service';
import { Router } from '@angular/router';

@Component({
  selector: 'app-blog-create',
  standalone: true,
  imports: [ReactiveFormsModule],
  templateUrl: './blog-create.component.html',
  styleUrl: './blog-create.component.scss',
})
export class BlogCreateComponent {
  blogForm: FormGroup;
  imagePreview: string | null = null;
  maxContentLength = 100;

  constructor(private fb: FormBuilder, private blogService: BlogService, private router: Router) {
    const today = new Date().toISOString().split('T')[0];
    this.blogForm = this.fb.group({
      title: ['', Validators.required],
      createdAt: [{ value: today, disabled: true }, Validators.required],
      slug: [{ value: '', disabled: true }, Validators.required],
      image: [null, Validators.required],
      content: [
        '',
        [Validators.required, Validators.maxLength(this.maxContentLength)],
      ],
    });

    this.blogForm.get('title')?.valueChanges.subscribe((title) => {
      const slug = this.generateSlug(title);
      this.blogForm.get('slug')?.setValue(slug, { emitEvent: false });
    });
  }

  onFileSelect(event: Event) {
    const file = (event.target as HTMLInputElement)?.files?.[0];
    if (file && (file.type === 'image/png' || file.type === 'image/jpeg')) {
      const reader = new FileReader();
      reader.onload = () => {
        this.imagePreview = reader.result as string;
        this.blogForm.patchValue({ image: file });
      };
      reader.readAsDataURL(file);
    } else {
      alert('Please upload a valid PNG or JPEG image.');
    }
  }

  generateSlug(title: string): string {
    return title
      ? title
          .toLowerCase()
          .trim()
          .replace(/[^a-z0-9]+/g, '-')
          .replace(/^-+|-+$/g, '')
      : '';
  }

  removeImage() {
    this.imagePreview = null;
    this.blogForm.patchValue({ image: null });
  }

  get contentLength() {
    return this.blogForm.get('content')?.value?.length || 0;
  }

  onSubmit() {
    if (this.blogForm.valid) {
      const formData = this.blogForm.getRawValue();
      const payload = {
        ...formData,
        image: this.blogForm.get('image')?.value,
      };

      this.blogService.createBlog(payload).subscribe({
        next: (res) => {
          this.router.navigate(['/blog', res.slug]);
        },
        error: (err) => {
          console.error('Error creating blog:', err);
          alert('Error creating blog.');
        },
      });
    } else {
      this.blogForm.markAllAsTouched();
    }
  }
}
