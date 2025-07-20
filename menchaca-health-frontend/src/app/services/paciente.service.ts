import { Injectable, inject } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { environment } from '../../environments/environment';

export interface Paciente {
  id?: number;
  nombre: string;
  apellido: string;
  correo_electronico: string;
  telefono: string;
  fecha_nacimiento: string;
  direccion?: string;
  seguro_medico?: string;
  contacto_emergencia?: string;
}

@Injectable({
  providedIn: 'root'
})
export class PacienteService {
  private http = inject(HttpClient);
  private apiUrl = `${environment.apiUrl}/pacientes`;

  getAll(): Observable<Paciente[]> {
    return this.http.get<Paciente[]>(this.apiUrl);
  }

  getById(id: number): Observable<Paciente> {
    return this.http.get<Paciente>(`${this.apiUrl}/${id}`);
  }

  create(paciente: Paciente): Observable<Paciente> {
    return this.http.post<Paciente>(this.apiUrl, paciente);
  }

  update(id: number, paciente: Paciente): Observable<Paciente> {
    return this.http.put<Paciente>(`${this.apiUrl}/${id}`, paciente);
  }

  delete(id: number): Observable<void> {
    return this.http.delete<void>(`${this.apiUrl}/${id}`);
  }

  search(query: string): Observable<Paciente[]> {
    return this.http.get<Paciente[]>(`${this.apiUrl}/search?q=${query}`);
  }

  getByMedico(medicoId: number): Observable<Paciente[]> {
    return this.http.get<Paciente[]>(`${this.apiUrl}/medico/${medicoId}`);
  }
}