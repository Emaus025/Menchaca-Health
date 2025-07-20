package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
    "github.com/gorilla/mux"
    "menchaca-health/internal/database"
    "menchaca-health/internal/handlers"
    "menchaca-health/internal/middleware"
    "menchaca-health/internal/services"
)

func main() {
    // Inicializar la base de datos
    log.Println("[DEBUG] Iniciando conexión a base de datos")
    db, err := database.NewDatabase()
    if err != nil {
        log.Fatalf("Failed to connect to the database: %v", err)
    }
    defer db.Close()
    log.Println("[DEBUG] Conexión a base de datos establecida")

    // Inicializar repositorios
    log.Printf("[DEBUG] Inicializando repositorios y handlers")
    appointmentRepo := database.NewAppointmentRepository(db)
    usuarioRepo := database.NewUsuarioRepository(db)
    // TODO: Agregar otros repositorios cuando se creen
    // pacienteRepo := database.NewPacienteRepository(db)
    // recetaRepo := database.NewRecetaRepository(db)
    // expedienteRepo := database.NewExpedienteRepository(db)
    // consultorioRepo := database.NewConsultorioRepository(db)
    // horarioRepo := database.NewHorarioRepository(db)

    // Inicializar servicios
    emailService := services.NewEmailService()

    // Inicializar handlers
    appointmentHandler := handlers.NewAppointmentHandler(appointmentRepo)
    usuarioHandler := handlers.NewUsuarioHandler(usuarioRepo, emailService)
    // TODO: Agregar otros handlers cuando se creen
    // pacienteHandler := handlers.NewPacienteHandler(pacienteRepo)
    // recetaHandler := handlers.NewRecetaHandler(recetaRepo)
    // expedienteHandler := handlers.NewExpedienteHandler(expedienteRepo)
    // consultorioHandler := handlers.NewConsultorioHandler(consultorioRepo)
    // horarioHandler := handlers.NewHorarioHandler(horarioRepo)

    // Configurar router
    router := mux.NewRouter()
    router.Use(middleware.CORS)

    // Manejo global de OPTIONS
    router.Methods("OPTIONS").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
    })

    // ==================== RUTAS DE AUTENTICACIÓN ====================
    authRouter := router.PathPrefix("/api/auth").Subrouter()
    authRouter.HandleFunc("/login", usuarioHandler.Login).Methods("POST")
    authRouter.HandleFunc("/register", usuarioHandler.Create).Methods("POST")
    authRouter.HandleFunc("/verify-email", usuarioHandler.VerifyEmail).Methods("GET")
    authRouter.HandleFunc("/logout", usuarioHandler.Logout).Methods("POST")
    authRouter.HandleFunc("/refresh-token", usuarioHandler.RefreshToken).Methods("POST")

    // ==================== RUTAS DE USUARIOS ====================
    usuarioRouter := router.PathPrefix("/api/usuarios").Subrouter()
    usuarioRouter.HandleFunc("", usuarioHandler.Create).Methods("POST")
    usuarioRouter.HandleFunc("", usuarioHandler.List).Methods("GET")
    usuarioRouter.HandleFunc("/{id}", usuarioHandler.GetByID).Methods("GET")
    usuarioRouter.HandleFunc("/{id}", usuarioHandler.Update).Methods("PUT")
    usuarioRouter.HandleFunc("/{id}", usuarioHandler.Delete).Methods("DELETE")
    usuarioRouter.HandleFunc("/role/{role}", usuarioHandler.GetByRole).Methods("GET")
    usuarioRouter.HandleFunc("/profile", usuarioHandler.GetProfile).Methods("GET")
    usuarioRouter.HandleFunc("/profile", usuarioHandler.UpdateProfile).Methods("PUT")
    usuarioRouter.HandleFunc("/change-password", usuarioHandler.ChangePassword).Methods("PUT")

    // ==================== RUTAS DE PACIENTES ====================
    pacienteRouter := router.PathPrefix("/api/pacientes").Subrouter()
    // TODO: Implementar cuando se cree el handler
    // pacienteRouter.HandleFunc("", pacienteHandler.Create).Methods("POST")
    // pacienteRouter.HandleFunc("", pacienteHandler.List).Methods("GET")
    // pacienteRouter.HandleFunc("/{id}", pacienteHandler.GetByID).Methods("GET")
    // pacienteRouter.HandleFunc("/{id}", pacienteHandler.Update).Methods("PUT")
    // pacienteRouter.HandleFunc("/{id}", pacienteHandler.Delete).Methods("DELETE")
    // pacienteRouter.HandleFunc("/search", pacienteHandler.Search).Methods("GET")
    // pacienteRouter.HandleFunc("/medico/{medicoId}", pacienteHandler.GetByMedico).Methods("GET")

    // ==================== RUTAS DE CITAS ====================
    appointmentRouter := router.PathPrefix("/api/citas").Subrouter()
    appointmentRouter.HandleFunc("", appointmentHandler.Create).Methods("POST")
    appointmentRouter.HandleFunc("", appointmentHandler.List).Methods("GET")
    appointmentRouter.HandleFunc("/{id}", appointmentHandler.GetByID).Methods("GET")
    appointmentRouter.HandleFunc("/{id}", appointmentHandler.Update).Methods("PUT")
    appointmentRouter.HandleFunc("/{id}", appointmentHandler.Delete).Methods("DELETE")
    appointmentRouter.HandleFunc("/patient/{patientId}", appointmentHandler.GetByPatientID).Methods("GET")
    appointmentRouter.HandleFunc("/date/{date}", appointmentHandler.GetByDate).Methods("GET")
    appointmentRouter.HandleFunc("/medico/{medicoId}", appointmentHandler.GetByMedico).Methods("GET")
    appointmentRouter.HandleFunc("/today", appointmentHandler.GetToday).Methods("GET")
    appointmentRouter.HandleFunc("/upcoming", appointmentHandler.GetUpcoming).Methods("GET")
    appointmentRouter.HandleFunc("/{id}/confirm", appointmentHandler.Confirm).Methods("PUT")
    appointmentRouter.HandleFunc("/{id}/cancel", appointmentHandler.Cancel).Methods("PUT")

    // ==================== RUTAS DE RECETAS ====================
    recetaRouter := router.PathPrefix("/api/recetas").Subrouter()
    // TODO: Implementar cuando se cree el handler
    // recetaRouter.HandleFunc("", recetaHandler.Create).Methods("POST")
    // recetaRouter.HandleFunc("", recetaHandler.List).Methods("GET")
    // recetaRouter.HandleFunc("/{id}", recetaHandler.GetByID).Methods("GET")
    // recetaRouter.HandleFunc("/{id}", recetaHandler.Update).Methods("PUT")
    // recetaRouter.HandleFunc("/{id}", recetaHandler.Delete).Methods("DELETE")
    // recetaRouter.HandleFunc("/paciente/{pacienteId}", recetaHandler.GetByPaciente).Methods("GET")
    // recetaRouter.HandleFunc("/medico/{medicoId}", recetaHandler.GetByMedico).Methods("GET")
    // recetaRouter.HandleFunc("/{id}/pdf", recetaHandler.GeneratePDF).Methods("GET")

    // ==================== RUTAS DE EXPEDIENTES ====================
    expedienteRouter := router.PathPrefix("/api/expedientes").Subrouter()
    // TODO: Implementar cuando se cree el handler
    // expedienteRouter.HandleFunc("", expedienteHandler.Create).Methods("POST")
    // expedienteRouter.HandleFunc("", expedienteHandler.List).Methods("GET")
    // expedienteRouter.HandleFunc("/{id}", expedienteHandler.GetByID).Methods("GET")
    // expedienteRouter.HandleFunc("/{id}", expedienteHandler.Update).Methods("PUT")
    // expedienteRouter.HandleFunc("/{id}", expedienteHandler.Delete).Methods("DELETE")
    // expedienteRouter.HandleFunc("/paciente/{pacienteId}", expedienteHandler.GetByPaciente).Methods("GET")
    // expedienteRouter.HandleFunc("/search", expedienteHandler.Search).Methods("GET")

    // ==================== RUTAS DE CONSULTORIOS ====================
    consultorioRouter := router.PathPrefix("/api/consultorios").Subrouter()
    // TODO: Implementar cuando se cree el handler
    // consultorioRouter.HandleFunc("", consultorioHandler.Create).Methods("POST")
    // consultorioRouter.HandleFunc("", consultorioHandler.List).Methods("GET")
    // consultorioRouter.HandleFunc("/{id}", consultorioHandler.GetByID).Methods("GET")
    // consultorioRouter.HandleFunc("/{id}", consultorioHandler.Update).Methods("PUT")
    // consultorioRouter.HandleFunc("/{id}", consultorioHandler.Delete).Methods("DELETE")
    // consultorioRouter.HandleFunc("/medico/{medicoId}", consultorioHandler.GetByMedico).Methods("GET")
    // consultorioRouter.HandleFunc("/disponibles", consultorioHandler.GetDisponibles).Methods("GET")

    // ==================== RUTAS DE HORARIOS ====================
    horarioRouter := router.PathPrefix("/api/horarios").Subrouter()
    // TODO: Implementar cuando se cree el handler
    // horarioRouter.HandleFunc("", horarioHandler.Create).Methods("POST")
    // horarioRouter.HandleFunc("", horarioHandler.List).Methods("GET")
    // horarioRouter.HandleFunc("/{id}", horarioHandler.GetByID).Methods("GET")
    // horarioRouter.HandleFunc("/{id}", horarioHandler.Update).Methods("PUT")
    // horarioRouter.HandleFunc("/{id}", horarioHandler.Delete).Methods("DELETE")
    // horarioRouter.HandleFunc("/medico/{medicoId}", horarioHandler.GetByMedico).Methods("GET")
    // horarioRouter.HandleFunc("/consultorio/{consultorioId}", horarioHandler.GetByConsultorio).Methods("GET")

    // ==================== RUTAS DE REPORTES ====================
    reporteRouter := router.PathPrefix("/api/reportes").Subrouter()
    // TODO: Implementar cuando se cree el handler
    // reporteRouter.HandleFunc("/citas", reporteHandler.CitasReport).Methods("GET")
    // reporteRouter.HandleFunc("/pacientes", reporteHandler.PacientesReport).Methods("GET")
    // reporteRouter.HandleFunc("/medicos", reporteHandler.MedicosReport).Methods("GET")
    // reporteRouter.HandleFunc("/ingresos", reporteHandler.IngresosReport).Methods("GET")
    // reporteRouter.HandleFunc("/estadisticas", reporteHandler.EstadisticasGenerales).Methods("GET")

    // ==================== RUTAS DE CONFIGURACIÓN ====================
    configRouter := router.PathPrefix("/api/configuracion").Subrouter()
    // TODO: Implementar cuando se cree el handler
    // configRouter.HandleFunc("/sistema", configHandler.GetSistema).Methods("GET")
    // configRouter.HandleFunc("/sistema", configHandler.UpdateSistema).Methods("PUT")
    // configRouter.HandleFunc("/notificaciones", configHandler.GetNotificaciones).Methods("GET")
    // configRouter.HandleFunc("/notificaciones", configHandler.UpdateNotificaciones).Methods("PUT")

    // Iniciar servidor
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    serverAddr := fmt.Sprintf(":%s", port)
    log.Printf("[INFO] Servidor iniciando en puerto %s", port)
    log.Printf("Server starting on %s", serverAddr)

    if err := http.ListenAndServe(serverAddr, router); err != nil {
        log.Fatalf("Error starting server: %v", err)
    }
}