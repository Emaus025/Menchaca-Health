# 🏥 Menchaca Health - Sistema de Gestión Médica
## 📋 Descripción del Proyecto
Menchaca Health es un sistema integral de gestión médica diseñado para optimizar la administración de clínicas y consultorios médicos. El sistema permite gestionar pacientes, médicos, citas, expedientes médicos, recetas y horarios de manera eficiente y segura.

### 🎯 Objetivo Principal
Digitalizar y centralizar todos los procesos administrativos y médicos de una clínica, proporcionando:

- Gestión eficiente de citas médicas
- Expedientes médicos digitales
- Control de usuarios por roles
- Generación de reportes y estadísticas
- Gestión de consultorios y horarios
- Sistema de recetas médicas
### 🌟 Funcionalidades Principales 👥 Gestión de Usuarios
- Registro y autenticación de usuarios
- Roles diferenciados: Administrador, Médico, Paciente, Recepcionista
- Gestión de perfiles y cambio de contraseñas
- Verificación de email 📅 Sistema de Citas
- Programación de citas médicas
- Confirmación y cancelación de citas
- Vista de citas por médico, paciente y fecha
- Citas del día y próximas citas 📋 Expedientes Médicos
- Historial clínico completo
- Antecedentes médicos
- Información de seguro médico
- Búsqueda avanzada de expedientes 💊 Gestión de Recetas
- Creación de recetas médicas
- Historial de recetas por paciente
- Generación de PDF para recetas 🏢 Administración de Consultorios
- Gestión de consultorios disponibles
- Asignación de médicos a consultorios
- Control de disponibilidad ⏰ Gestión de Horarios
- Horarios de médicos por día de la semana
- Turnos matutinos y vespertinos
- Control de disponibilidad por consultorio 📊 Reportes y Estadísticas
- Reportes de citas por período
- Estadísticas de pacientes
- Reportes de ingresos
- Dashboard con métricas generales
## 🏗️ Arquitectura del Sistema
### 🎨 Frontend - Angular 18
Ubicación: menchaca-health-frontend/
 📁 Estructura del Proyecto 🛠️ Tecnologías Frontend
- Angular 18 - Framework principal
- PrimeNG - Biblioteca de componentes UI
- RxJS - Programación reactiva
- TypeScript - Lenguaje de programación
- Angular Router - Navegación y guards
- HttpClient - Comunicación con API 🔐 Sistema de Autenticación
- Guards de autenticación para rutas protegidas
- Guards de roles para control de acceso
- Interceptors para manejo de tokens
- Gestión de sesiones con localStorage
### ⚙️ Backend - Go (Golang)
Ubicación: menchaca-health/
 📁 Estructura del Proyecto 🛠️ Tecnologías Backend
- Go 1.24 - Lenguaje de programación
- Gorilla Mux - Router HTTP
- PostgreSQL - Base de datos
- pgx/v5 - Driver de PostgreSQL
- Argon2 - Hash de contraseñas
- UUID - Identificadores únicos
- CORS - Configuración de CORS 🗄️ Base de Datos
- PostgreSQL como sistema de gestión de base de datos
- Migraciones para control de versiones del esquema
- Conexión pooling para optimización de rendimiento
## 🚀 Instalación y Configuración
### 📋 Prerrequisitos
- Node.js (v18 o superior)
- Angular CLI (v18)
- Go (v1.24 o superior)
- PostgreSQL (v12 o superior)
- Git
### 🔧 Configuración del Backend
1. 1.
   Clonar el repositorio:
2. 1.
   Configurar variables de entorno:
3. 1.
   Instalar dependencias:
4. 1.
   Configurar base de datos:
5. 1.
   Ejecutar el servidor:
El servidor estará disponible en http://localhost:8080

### 🎨 Configuración del Frontend
1. 1.
   Navegar al directorio del frontend:
2. 1.
   Instalar dependencias:
3. 1.
   Configurar entorno:
4. 1.
   Ejecutar la aplicación:
La aplicación estará disponible en http://localhost:4200

## 📡 API Endpoints
### 🔐 Autenticación
### 👥 Usuarios
### 📅 Citas
### 📋 Expedientes (Próximamente)
```
GET    /api/expedientes         # 
Listar expedientes
POST   /api/expedientes         # 
Crear expediente
GET    /api/expedientes/{id}    # 
Obtener expediente
PUT    /api/expedientes/{id}    # 
Actualizar expediente
DELETE /api/expedientes/{id}    # 
Eliminar expediente
GET    /api/expedientes/paciente/
{id} # Expediente por paciente
```
## 👤 Roles y Permisos
### 🔑 Administrador
- Gestión completa de usuarios
- Acceso a todos los reportes
- Configuración del sistema
- Gestión de consultorios y horarios
### 👨‍⚕️ Médico
- Gestión de sus citas
- Acceso a expedientes de sus pacientes
- Creación de recetas
- Gestión de sus horarios
### 🏥 Recepcionista
- Gestión de citas
- Registro de pacientes
- Consulta de horarios
- Reportes básicos
### 🧑‍🤝‍🧑 Paciente
- Ver sus citas
- Actualizar perfil
- Ver su expediente médico
- Ver sus recetas
## 🧪 Testing
### Backend Testing
### Frontend Testing
## 📦 Deployment
### 🐳 Docker (Recomendado)
1. 1.
   Backend Dockerfile:
2. 1.
   Frontend Dockerfile:
3. 1.
   Docker Compose:
## 🔧 Desarrollo
### 📝 Convenciones de Código Backend (Go)
- Seguir las convenciones estándar de Go
- Usar gofmt para formateo
- Documentar funciones públicas
- Manejo de errores explícito Frontend (Angular)
- Seguir Angular Style Guide
- Usar TypeScript estricto
- Componentes standalone
- Servicios inyectables
### 🌿 Git Workflow
## 🤝 Contribución
1. 1.
   Fork el proyecto
2. 2.
   Crear rama para la funcionalidad ( git checkout -b feature/AmazingFeature )
3. 3.
   Commit los cambios ( git commit -m 'Add some AmazingFeature' )
4. 4.
   Push a la rama ( git push origin feature/AmazingFeature )
5. 5.
   Abrir un Pull Request
