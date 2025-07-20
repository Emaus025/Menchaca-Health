import { Injectable, inject, PLATFORM_ID } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable, BehaviorSubject } from 'rxjs';
import { tap } from 'rxjs/operators';
import { isPlatformBrowser } from '@angular/common';

export interface Usuario {
  id: number;
  nombre: string;
  apellido: string;
  tipo_usuario: string;
  correo_electronico: string;
  telefono: string;
  fecha_nacimiento: string;
}

export interface LoginRequest {
  correo_electronico: string;
  contrasena: string;
}

export interface RegisterRequest {
  nombre: string;
  apellido: string;
  tipo_usuario: string;
  correo_electronico: string;
  telefono: string;
  fecha_nacimiento: string;
  contrasena: string;
}

export interface LoginResponse {
  usuario: Usuario;
  token?: string;
  message: string;
}

@Injectable({
  providedIn: 'root'
})
export class AuthService {
  private http = inject(HttpClient);
  private platformId = inject(PLATFORM_ID);
  private apiUrl = 'http://localhost:8080/api';
  private currentUserSubject = new BehaviorSubject<Usuario | null>(null);
  public currentUser$ = this.currentUserSubject.asObservable();

  constructor() {
    // Solo acceder a localStorage en el navegador
    if (isPlatformBrowser(this.platformId)) {
      const storedUser = localStorage.getItem('currentUser');
      if (storedUser) {
        this.currentUserSubject.next(JSON.parse(storedUser));
      }
    }
  }

  login(credentials: LoginRequest): Observable<LoginResponse> {
    return this.http.post<LoginResponse>(`${this.apiUrl}/auth/login`, credentials)
      .pipe(
        tap((response: LoginResponse) => {
          if (response.usuario && isPlatformBrowser(this.platformId)) {
            localStorage.setItem('currentUser', JSON.stringify(response.usuario));
            this.currentUserSubject.next(response.usuario);
          }
        })
      );
  }

  register(userData: RegisterRequest): Observable<any> {
    return this.http.post(`${this.apiUrl}/usuarios`, userData);
  }

  logout(): void {
    if (isPlatformBrowser(this.platformId)) {
      // Limpiar todos los datos almacenados
      localStorage.removeItem('currentUser');
      localStorage.removeItem('token');
      localStorage.removeItem('user');
    }
    this.currentUserSubject.next(null);
  }

  getCurrentUser(): Usuario | null {
    return this.currentUserSubject.value;
  }

  isAuthenticated(): boolean {
    return this.getCurrentUser() !== null;
  }

  getUsersByRole(role: string): Observable<Usuario[]> {
    return this.http.get<Usuario[]>(`${this.apiUrl}/usuarios/role/${role}`);
  }
}
