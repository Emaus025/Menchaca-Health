import { Injectable, inject } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { environment } from '../../environments/environment';

export interface Cita {
  id?: number;
  paciente_id: number;
  medico_id: number;
  fecha: string;
  hora: string;
  duracion: number;
  tipo: string;
  estado: string;
  notas?: string;
  paciente?: any;
  medico?: any;
}

@Injectable({
  providedIn: 'root'
})
export class CitaService {
  private http = inject(HttpClient);
  private apiUrl = `${environment.apiUrl}/citas`;

  getAll(): Observable<Cita[]> {
    return this.http.get<Cita[]>(this.apiUrl);
  }

  getById(id: number): Observable<Cita> {
    return this.http.get<Cita>(`${this.apiUrl}/${id}`);
  }

  create(cita: Cita): Observable<Cita> {
    return this.http.post<Cita>(this.apiUrl, cita);
  }

  update(id: number, cita: Cita): Observable<Cita> {
    return this.http.put<Cita>(`${this.apiUrl}/${id}`, cita);
  }

  delete(id: number): Observable<void> {
    return this.http.delete<void>(`${this.apiUrl}/${id}`);
  }

  getByPaciente(pacienteId: number): Observable<Cita[]> {
    return this.http.get<Cita[]>(`${this.apiUrl}/patient/${pacienteId}`);
  }

  getByMedico(medicoId: number): Observable<Cita[]> {
    return this.http.get<Cita[]>(`${this.apiUrl}/medico/${medicoId}`);
  }

  getByDate(date: string): Observable<Cita[]> {
    return this.http.get<Cita[]>(`${this.apiUrl}/date/${date}`);
  }

  getToday(): Observable<Cita[]> {
    return this.http.get<Cita[]>(`${this.apiUrl}/today`);
  }

  getUpcoming(): Observable<Cita[]> {
    return this.http.get<Cita[]>(`${this.apiUrl}/upcoming`);
  }

  confirm(id: number): Observable<Cita> {
    return this.http.put<Cita>(`${this.apiUrl}/${id}/confirm`, {});
  }

  cancel(id: number): Observable<Cita> {
    return this.http.put<Cita>(`${this.apiUrl}/${id}/cancel`, {});
  }
}