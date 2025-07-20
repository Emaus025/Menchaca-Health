import { Component, OnInit, Inject, PLATFORM_ID, inject } from '@angular/core';
import { CommonModule, isPlatformBrowser } from '@angular/common';
import { Router } from '@angular/router';
import { CardModule } from 'primeng/card';
import { ButtonModule } from 'primeng/button';
import { MenubarModule } from 'primeng/menubar';
import { MenuItem } from 'primeng/api';
import { AvatarModule } from 'primeng/avatar';
import { BadgeModule } from 'primeng/badge';
import { RippleModule } from 'primeng/ripple';
import { TooltipModule } from 'primeng/tooltip';
import { AuthService } from '../auth/auth.service';

@Component({
  selector: 'app-home',
  standalone: true,
  imports: [
    CommonModule,
    CardModule,
    ButtonModule,
    MenubarModule,
    AvatarModule,
    BadgeModule,
    RippleModule,
    TooltipModule
  ],
  templateUrl: './home.component.html',
  styleUrl: './home.component.css'
})
export class HomeComponent implements OnInit {
  items: MenuItem[] = [];
  userInfo: any = {};
  private isBrowser: boolean;
  private authService = inject(AuthService);

  constructor(
    private router: Router,
    @Inject(PLATFORM_ID) private platformId: Object
  ) {
    this.isBrowser = isPlatformBrowser(this.platformId);
  }

  ngOnInit() {
    this.loadUserInfo();
    this.initializeMenu();
  }

  initializeMenu() {
    const userType = this.userInfo.tipo_usuario?.toLowerCase();
    
    switch (userType) {
      case 'paciente':
        this.items = this.getPatientMenu();
        break;
      case 'medico':
        this.items = this.getDoctorMenu();
        break;
      case 'administrador':
        this.items = this.getAdminMenu();
        break;
      case 'enfermero':
        this.items = this.getNurseMenu();
        break;
      default:
        this.items = this.getDefaultMenu();
    }
  }

  getPatientMenu(): MenuItem[] {
    return [
      {
        label: 'Mi Panel',
        icon: 'pi pi-home',
        routerLink: '/home'
      },
      {
        label: 'Mis Citas',
        icon: 'pi pi-calendar',
        routerLink: '/mis-citas'
      },
      {
        label: 'Mis Recetas',
        icon: 'pi pi-file-pdf',
        routerLink: '/mis-recetas'
      },
      {
        label: 'Mi Perfil',
        icon: 'pi pi-user',
        routerLink: '/perfil'
      }
    ];
  }

  getDoctorMenu(): MenuItem[] {
    return [
      {
        label: 'Mi Consultorio',
        icon: 'pi pi-home',
        routerLink: '/home'
      },
      {
        label: 'Pacientes',
        icon: 'pi pi-users',
        items: [
          {
            label: 'Lista de Pacientes',
            icon: 'pi pi-list',
            routerLink: '/pacientes'
          },
          {
            label: 'Expedientes',
            icon: 'pi pi-folder',
            routerLink: '/expedientes'
          }
        ]
      },
      {
        label: 'Citas',
        icon: 'pi pi-calendar',
        items: [
          {
            label: 'Mi Agenda',
            icon: 'pi pi-calendar-plus',
            routerLink: '/citas'
          },
          {
            label: 'Nueva Cita',
            icon: 'pi pi-plus',
            routerLink: '/citas/nueva'
          }
        ]
      },
      {
        label: 'Recetas',
        icon: 'pi pi-file-pdf',
        items: [
          {
            label: 'Crear Receta',
            icon: 'pi pi-plus',
            routerLink: '/recetas/nueva'
          },
          {
            label: 'Mis Recetas',
            icon: 'pi pi-list',
            routerLink: '/recetas'
          }
        ]
      },
      {
        label: 'Mi Consultorio',
        icon: 'pi pi-building',
        routerLink: '/consultorio'
      }
    ];
  }

  getAdminMenu(): MenuItem[] {
    return [
      {
        label: 'Dashboard',
        icon: 'pi pi-home',
        routerLink: '/home'
      },
      {
        label: 'Usuarios',
        icon: 'pi pi-users',
        items: [
          {
            label: 'Gestionar Usuarios',
            icon: 'pi pi-users',
            routerLink: '/usuarios'
          },
          {
            label: 'Nuevo Usuario',
            icon: 'pi pi-user-plus',
            routerLink: '/usuarios/nuevo'
          }
        ]
      },
      {
        label: 'Pacientes',
        icon: 'pi pi-heart',
        items: [
          {
            label: 'Lista de Pacientes',
            icon: 'pi pi-list',
            routerLink: '/pacientes'
          },
          {
            label: 'Expedientes',
            icon: 'pi pi-folder',
            routerLink: '/expedientes'
          }
        ]
      },
      {
        label: 'Citas',
        icon: 'pi pi-calendar',
        items: [
          {
            label: 'Todas las Citas',
            icon: 'pi pi-calendar',
            routerLink: '/citas'
          },
          {
            label: 'Nueva Cita',
            icon: 'pi pi-plus',
            routerLink: '/citas/nueva'
          }
        ]
      },
      {
        label: 'Consultorios',
        icon: 'pi pi-building',
        routerLink: '/consultorios'
      },
      {
        label: 'Reportes',
        icon: 'pi pi-chart-bar',
        routerLink: '/reportes'
      },
      {
        label: 'Configuración',
        icon: 'pi pi-cog',
        routerLink: '/configuracion'
      }
    ];
  }

  getNurseMenu(): MenuItem[] {
    return [
      {
        label: 'Panel',
        icon: 'pi pi-home',
        routerLink: '/home'
      },
      {
        label: 'Expedientes',
        icon: 'pi pi-folder',
        routerLink: '/expedientes'
      },
      {
        label: 'Horarios',
        icon: 'pi pi-clock',
        routerLink: '/horarios'
      },
      {
        label: 'Consultorios',
        icon: 'pi pi-building',
        routerLink: '/consultorios'
      }
    ];
  }

  getDefaultMenu(): MenuItem[] {
    return [
      {
        label: 'Dashboard',
        icon: 'pi pi-home',
        routerLink: '/home'
      }
    ];
  }

  loadUserInfo() {
    if (this.isBrowser) {
      const userData = localStorage.getItem('user');
      if (userData) {
        try {
          this.userInfo = JSON.parse(userData);
        } catch (error) {
          console.error('Error parsing user data:', error);
          this.setDefaultUserInfo();
        }
      } else {
        this.setDefaultUserInfo();
      }
    } else {
      this.setDefaultUserInfo();
    }
  }

  private setDefaultUserInfo() {
    this.userInfo = {
      nombre: 'Usuario',
      apellido: 'Demo',
      tipo_usuario: 'administrador'
    };
  }

  logout() {
    this.authService.logout();
    
    if (this.isBrowser) {
      localStorage.removeItem('token');
      localStorage.removeItem('user');
    }
    
    this.router.navigate(['/auth/login']);
  }

  getUserInitials(): string {
    const nombre = this.userInfo.nombre || 'U';
    const apellido = this.userInfo.apellido || 'D';
    return nombre.charAt(0).toUpperCase() + apellido.charAt(0).toUpperCase();
  }

  // Métodos para verificar el tipo de usuario
  isPatient(): boolean {
    return this.userInfo.tipo_usuario?.toLowerCase() === 'paciente';
  }

  isDoctor(): boolean {
    return this.userInfo.tipo_usuario?.toLowerCase() === 'medico';
  }

  isAdmin(): boolean {
    return this.userInfo.tipo_usuario?.toLowerCase() === 'administrador';
  }

  isNurse(): boolean {
    return this.userInfo.tipo_usuario?.toLowerCase() === 'enfermero';
  }
}