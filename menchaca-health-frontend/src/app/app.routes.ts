import { Routes } from '@angular/router';
import { AuthGuard } from './guards/auth.guard';
import { RoleGuard } from './guards/role.guard';

export const routes: Routes = [
  // ==================== RUTAS PÚBLICAS ====================
  {
    path: 'auth',
    children: [
      {
        path: 'login',
        loadComponent: () => import('./auth/login/login.component').then(m => m.LoginComponent)
      },
      {
        path: 'register',
        loadComponent: () => import('./auth/register/register.component').then(m => m.RegisterComponent)
      },
      {
        path: 'verify-email',
        loadComponent: () => import('./auth/verify-email/verify-email.component').then(m => m.VerifyEmailComponent)
      },
      {
        path: 'forgot-password',
        loadComponent: () => import('./auth/forgot-password/forgot-password.component').then(m => m.ForgotPasswordComponent)
      },
      {
        path: '',
        redirectTo: 'login',
        pathMatch: 'full'
      }
    ]
  },

  // ==================== RUTAS PROTEGIDAS ====================
  {
    path: 'home',
    loadComponent: () => import('./home/home.component').then(m => m.HomeComponent),
    canActivate: [AuthGuard]
  },

  // ==================== RUTAS DE USUARIOS (Solo Admin) ====================
  {
    path: 'usuarios',
    canActivate: [AuthGuard, RoleGuard],
    data: { roles: ['administrador'] },
    children: [
      {
        path: '',
        loadComponent: () => import('./usuarios/usuario-list/usuario-list.component').then(m => m.UsuarioListComponent)
      },
      {
        path: 'nuevo',
        loadComponent: () => import('./usuarios/usuario-form/usuario-form.component').then(m => m.UsuarioFormComponent)
      },
      {
        path: 'editar/:id',
        loadComponent: () => import('./usuarios/usuario-form/usuario-form.component').then(m => m.UsuarioFormComponent)
      },
      {
        path: 'detalle/:id',
        loadComponent: () => import('./usuarios/usuario-detail/usuario-detail.component').then(m => m.UsuarioDetailComponent)
      }
    ]
  },

  // ==================== RUTAS DE PACIENTES ====================
  {
    path: 'pacientes',
    canActivate: [AuthGuard, RoleGuard],
    data: { roles: ['administrador', 'medico', 'enfermero'] },
    children: [
      {
        path: '',
        loadComponent: () => import('./pacientes/paciente-list/paciente-list.component').then(m => m.PacienteListComponent)
      },
      {
        path: 'nuevo',
        loadComponent: () => import('./pacientes/paciente-form/paciente-form.component').then(m => m.PacienteFormComponent),
        canActivate: [RoleGuard],
        data: { roles: ['administrador', 'medico'] }
      },
      {
        path: 'editar/:id',
        loadComponent: () => import('./pacientes/paciente-form/paciente-form.component').then(m => m.PacienteFormComponent),
        canActivate: [RoleGuard],
        data: { roles: ['administrador', 'medico'] }
      },
      {
        path: 'detalle/:id',
        loadComponent: () => import('./pacientes/paciente-detail/paciente-detail.component').then(m => m.PacienteDetailComponent)
      }
    ]
  },

  // ==================== RUTAS DE CITAS ====================
  {
    path: 'citas',
    canActivate: [AuthGuard],
    children: [
      {
        path: '',
        loadComponent: () => import('./citas/cita-list/cita-list.component').then(m => m.CitaListComponent)
      },
      {
        path: 'nueva',
        loadComponent: () => import('./citas/cita-form/cita-form.component').then(m => m.CitaFormComponent),
        canActivate: [RoleGuard],
        data: { roles: ['administrador', 'medico'] }
      },
      {
        path: 'editar/:id',
        loadComponent: () => import('./citas/cita-form/cita-form.component').then(m => m.CitaFormComponent),
        canActivate: [RoleGuard],
        data: { roles: ['administrador', 'medico'] }
      },
      {
        path: 'detalle/:id',
        loadComponent: () => import('./citas/cita-detail/cita-detail.component').then(m => m.CitaDetailComponent)
      },
      {
        path: 'agenda',
        loadComponent: () => import('./citas/agenda/agenda.component').then(m => m.AgendaComponent)
      }
    ]
  },

  // ==================== RUTAS DE CITAS PARA PACIENTES ====================
  {
    path: 'mis-citas',
    loadComponent: () => import('./pacientes/mis-citas/mis-citas.component').then(m => m.MisCitasComponent),
    canActivate: [AuthGuard, RoleGuard],
    data: { roles: ['paciente'] }
  },

  // ==================== RUTAS DE RECETAS ====================
  {
    path: 'recetas',
    canActivate: [AuthGuard, RoleGuard],
    data: { roles: ['administrador', 'medico'] },
    children: [
      {
        path: '',
        loadComponent: () => import('./recetas/receta-list/receta-list.component').then(m => m.RecetaListComponent)
      },
      {
        path: 'nueva',
        loadComponent: () => import('./recetas/receta-form/receta-form.component').then(m => m.RecetaFormComponent)
      },
      {
        path: 'editar/:id',
        loadComponent: () => import('./recetas/receta-form/receta-form.component').then(m => m.RecetaFormComponent)
      },
      {
        path: 'detalle/:id',
        loadComponent: () => import('./recetas/receta-detail/receta-detail.component').then(m => m.RecetaDetailComponent)
      }
    ]
  },

  // ==================== RUTAS DE RECETAS PARA PACIENTES ====================
  {
    path: 'mis-recetas',
    loadComponent: () => import('./pacientes/mis-recetas/mis-recetas.component').then(m => m.MisRecetasComponent),
    canActivate: [AuthGuard, RoleGuard],
    data: { roles: ['paciente'] }
  },

  // ==================== RUTAS DE EXPEDIENTES ====================
  {
    path: 'expedientes',
    canActivate: [AuthGuard, RoleGuard],
    data: { roles: ['administrador', 'medico', 'enfermero'] },
    children: [
      {
        path: '',
        loadComponent: () => import('./expedientes/expediente-list/expediente-list.component').then(m => m.ExpedienteListComponent)
      },
      {
        path: 'nuevo',
        loadComponent: () => import('./expedientes/expediente-form/expediente-form.component').then(m => m.ExpedienteFormComponent),
        canActivate: [RoleGuard],
        data: { roles: ['administrador', 'medico'] }
      },
      {
        path: 'editar/:id',
        loadComponent: () => import('./expedientes/expediente-form/expediente-form.component').then(m => m.ExpedienteFormComponent),
        canActivate: [RoleGuard],
        data: { roles: ['administrador', 'medico'] }
      },
      {
        path: 'detalle/:id',
        loadComponent: () => import('./expedientes/expediente-detail/expediente-detail.component').then(m => m.ExpedienteDetailComponent)
      }
    ]
  },

  // ==================== RUTAS DE CONSULTORIOS ====================
  {
    path: 'consultorios',
    canActivate: [AuthGuard, RoleGuard],
    data: { roles: ['administrador', 'medico', 'enfermero'] },
    children: [
      {
        path: '',
        loadComponent: () => import('./consultorios/consultorio-list/consultorio-list.component').then(m => m.ConsultorioListComponent)
      },
      {
        path: 'nuevo',
        loadComponent: () => import('./consultorios/consultorio-form/consultorio-form.component').then(m => m.ConsultorioFormComponent),
        canActivate: [RoleGuard],
        data: { roles: ['administrador'] }
      },
      {
        path: 'editar/:id',
        loadComponent: () => import('./consultorios/consultorio-form/consultorio-form.component').then(m => m.ConsultorioFormComponent),
        canActivate: [RoleGuard],
        data: { roles: ['administrador'] }
      },
      {
        path: 'detalle/:id',
        loadComponent: () => import('./consultorios/consultorio-detail/consultorio-detail.component').then(m => m.ConsultorioDetailComponent)
      }
    ]
  },

  // ==================== RUTAS DE HORARIOS ====================
  {
    path: 'horarios',
    canActivate: [AuthGuard, RoleGuard],
    data: { roles: ['administrador', 'medico', 'enfermero'] },
    children: [
      {
        path: '',
        loadComponent: () => import('./horarios/horario-list/horario-list.component').then(m => m.HorarioListComponent)
      },
      {
        path: 'nuevo',
        loadComponent: () => import('./horarios/horario-form/horario-form.component').then(m => m.HorarioFormComponent),
        canActivate: [RoleGuard],
        data: { roles: ['administrador', 'medico'] }
      },
      {
        path: 'editar/:id',
        loadComponent: () => import('./horarios/horario-form/horario-form.component').then(m => m.HorarioFormComponent),
        canActivate: [RoleGuard],
        data: { roles: ['administrador', 'medico'] }
      },
      {
        path: 'detalle/:id',
        loadComponent: () => import('./horarios/horario-detail/horario-detail.component').then(m => m.HorarioDetailComponent)
      }
    ]
  },

  // ==================== RUTAS DE REPORTES ====================
  {
    path: 'reportes',
    canActivate: [AuthGuard, RoleGuard],
    data: { roles: ['administrador'] },
    children: [
      {
        path: '',
        loadComponent: () => import('./reportes/reporte-dashboard/reporte-dashboard.component').then(m => m.ReporteDashboardComponent)
      },
      {
        path: 'citas',
        loadComponent: () => import('./reportes/reporte-citas/reporte-citas.component').then(m => m.ReporteCitasComponent)
      },
      {
        path: 'pacientes',
        loadComponent: () => import('./reportes/reporte-pacientes/reporte-pacientes.component').then(m => m.ReportePacientesComponent)
      },
      {
        path: 'medicos',
        loadComponent: () => import('./reportes/reporte-medicos/reporte-medicos.component').then(m => m.ReporteMedicosComponent)
      },
      {
        path: 'ingresos',
        loadComponent: () => import('./reportes/reporte-ingresos/reporte-ingresos.component').then(m => m.ReporteIngresosComponent)
      }
    ]
  },

  // ==================== RUTAS DE CONFIGURACIÓN ====================
  {
    path: 'configuracion',
    canActivate: [AuthGuard, RoleGuard],
    data: { roles: ['administrador'] },
    children: [
      {
        path: '',
        loadComponent: () => import('./configuracion/config-dashboard/config-dashboard.component').then(m => m.ConfigDashboardComponent)
      },
      {
        path: 'sistema',
        loadComponent: () => import('./configuracion/config-sistema/config-sistema.component').then(m => m.ConfigSistemaComponent)
      },
      {
        path: 'notificaciones',
        loadComponent: () => import('./configuracion/config-notificaciones/config-notificaciones.component').then(m => m.ConfigNotificacionesComponent)
      },
      {
        path: 'usuarios',
        loadComponent: () => import('./configuracion/config-usuarios/config-usuarios.component').then(m => m.ConfigUsuariosComponent)
      }
    ]
  },

  // ==================== RUTAS DE PERFIL ====================
  {
    path: 'perfil',
    canActivate: [AuthGuard],
    children: [
      {
        path: '',
        loadComponent: () => import('./perfil/perfil-view/perfil-view.component').then(m => m.PerfilViewComponent)
      },
      {
        path: 'editar',
        loadComponent: () => import('./perfil/perfil-edit/perfil-edit.component').then(m => m.PerfilEditComponent)
      },
      {
        path: 'cambiar-password',
        loadComponent: () => import('./perfil/cambiar-password/cambiar-password.component').then(m => m.CambiarPasswordComponent)
      }
    ]
  },

  // ==================== RUTAS DE MI CONSULTORIO (Solo Médicos) ====================
  {
    path: 'consultorio',
    loadComponent: () => import('./medicos/mi-consultorio/mi-consultorio.component').then(m => m.MiConsultorioComponent),
    canActivate: [AuthGuard, RoleGuard],
    data: { roles: ['medico'] }
  },

  // ==================== RUTAS POR DEFECTO ====================
  {
    path: '',
    redirectTo: '/home',
    pathMatch: 'full'
  },
  {
    path: '**',
    loadComponent: () => import('./shared/not-found/not-found.component').then(m => m.NotFoundComponent)
  }
];
