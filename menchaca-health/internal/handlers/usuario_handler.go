package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"menchaca-health/internal/database"
	"menchaca-health/internal/models"
	"github.com/gorilla/mux"
	"menchaca-health/internal/services"
)

type UsuarioHandler struct {
	repo         *database.UsuarioRepository
	emailService *services.EmailService
}

func NewUsuarioHandler(repo *database.UsuarioRepository, emailService *services.EmailService) *UsuarioHandler {
	return &UsuarioHandler{
		repo:         repo,
		emailService: emailService,
	}
}

// Modificar el método Create con logs detallados
func (h *UsuarioHandler) Create(w http.ResponseWriter, r *http.Request) {
	log.Println("[DEBUG] Iniciando creación de usuario")
	
	var usuario models.Usuario
	if err := json.NewDecoder(r.Body).Decode(&usuario); err != nil {
		log.Printf("[ERROR] Error decodificando JSON: %v", err)
		http.Error(w, "Datos inválidos", http.StatusBadRequest)
		return
	}

	log.Printf("[DEBUG] Usuario recibido: %+v", usuario)

	// Validaciones básicas
	if usuario.Nombre == "" || usuario.CorreoElectronico == "" || usuario.Contrasena == "" {
		log.Printf("[ERROR] Campos requeridos faltantes. Nombre: '%s', Email: '%s', Password: '%s'", 
			usuario.Nombre, usuario.CorreoElectronico, "[OCULTO]")
		http.Error(w, "Campos requeridos faltantes", http.StatusBadRequest)
		return
	}

	log.Println("[DEBUG] Validaciones básicas pasadas")

	// Generar token de verificación
	log.Println("[DEBUG] Generando token de verificación")
	token, err := h.emailService.GenerateVerificationToken()
	if err != nil {
		log.Printf("[ERROR] Error generando token: %v", err)
		http.Error(w, "Error generando token", http.StatusInternalServerError)
		return
	}
	log.Printf("[DEBUG] Token generado exitosamente: %s", token[:10]+"...")

	// Crear usuario con verificación pendiente
	log.Println("[DEBUG] Creando usuario en base de datos")
	if err := h.repo.CreateWithVerification(&usuario, token); err != nil {
		log.Printf("[ERROR] Error creando usuario en BD: %v", err)
		if strings.Contains(err.Error(), "duplicate key") {
			log.Printf("[ERROR] Email duplicado: %s", usuario.CorreoElectronico)
			http.Error(w, "El correo electrónico ya está registrado", http.StatusConflict)
			return
		}
		log.Printf("[ERROR] Error de base de datos: %v", err)
		http.Error(w, fmt.Sprintf("Error creando usuario: %v", err), http.StatusInternalServerError)
		return
	}
	log.Printf("[DEBUG] Usuario creado exitosamente con ID: %d", usuario.ID)

	// Enviar email de verificación
	log.Println("[DEBUG] Enviando email de verificación")
	if err := h.emailService.SendVerificationEmail(usuario.CorreoElectronico, token, usuario.Nombre); err != nil {
		// Log del error pero no fallar el registro
		log.Printf("[WARNING] Error enviando email de verificación: %v", err)
		log.Printf("[WARNING] Usuario creado pero email no enviado")
	} else {
		log.Println("[DEBUG] Email de verificación enviado exitosamente")
	}

	// No devolver la contraseña
	usuario.Contrasena = ""
	
	log.Println("[DEBUG] Proceso de creación completado exitosamente")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Usuario registrado. Revisa tu email para verificar tu cuenta.",
		"usuario": usuario,
	})
}

// Nuevo método para verificar email
func (h *UsuarioHandler) VerifyEmail(w http.ResponseWriter, r *http.Request) {
	token := r.URL.Query().Get("token")
	if token == "" {
		http.Error(w, "Token requerido", http.StatusBadRequest)
		return
	}

	if err := h.repo.VerifyEmail(token); err != nil {
		http.Error(w, "Token inválido o expirado", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Email verificado exitosamente. Ya puedes iniciar sesión.",
	})
}

// Modificar ValidateLogin para verificar email
func (h *UsuarioHandler) ValidateLogin(w http.ResponseWriter, r *http.Request) {
	var loginData struct {
		CorreoElectronico string `json:"correo_electronico"`
		Contrasena        string `json:"contrasena"`
	}

	if err := json.NewDecoder(r.Body).Decode(&loginData); err != nil {
		http.Error(w, "Datos inválidos", http.StatusBadRequest)
		return
	}

	
	verified, err := h.repo.IsEmailVerified(loginData.CorreoElectronico)
	if err != nil {
		http.Error(w, "Error verificando usuario", http.StatusInternalServerError)
		return
	}

	if !verified {
		http.Error(w, "Email no verificado", http.StatusUnauthorized)
		return
	}

	usuario, err := h.repo.ValidateLogin(loginData.CorreoElectronico, loginData.Contrasena)
	if err != nil {
		http.Error(w, "Credenciales inválidas", http.StatusUnauthorized)
		return
	}

	// Respuesta exitosa
	response := map[string]interface{}{
		"message": "Login exitoso",
		"usuario": usuario,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// GetByID maneja la obtención de usuario por ID
func (h *UsuarioHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}
	
	usuario, err := h.repo.GetByID(id)
	if err != nil {
		if strings.Contains(err.Error(), "no encontrado") {
			http.Error(w, "Usuario no encontrado", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(usuario)
}

// Update maneja la actualización de usuarios
func (h *UsuarioHandler) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}
	
	var usuario models.Usuario
	if err := json.NewDecoder(r.Body).Decode(&usuario); err != nil {
		http.Error(w, "Datos inválidos", http.StatusBadRequest)
		return
	}
	
	usuario.ID = id
	
	if err := h.repo.Update(&usuario); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	// Obtener usuario actualizado
	usuarioActualizado, err := h.repo.GetByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(usuarioActualizado)
}

// Delete maneja la eliminación de usuarios
func (h *UsuarioHandler) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}
	
	if err := h.repo.Delete(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	w.WriteHeader(http.StatusNoContent)
}

// List maneja la obtención de todos los usuarios
func (h *UsuarioHandler) List(w http.ResponseWriter, r *http.Request) {
	usuarios, err := h.repo.List()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(usuarios)
}

// GetByRole maneja la obtención de usuarios por rol
func (h *UsuarioHandler) GetByRole(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	role := vars["role"]
	
	usuarios, err := h.repo.GetByRole(role)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(usuarios)
}

// Login maneja la autenticación de usuarios
func (h *UsuarioHandler) Login(w http.ResponseWriter, r *http.Request) {
	var credentials struct {
		CorreoElectronico string `json:"correo_electronico"`
		Contrasena        string `json:"contrasena"`
	}
	
	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		log.Printf("[ERROR] Error decodificando credenciales: %v", err)
		http.Error(w, "Datos inválidos", http.StatusBadRequest)
		return
	}
	
	// DEBUG: Log de credenciales recibidas (SIN mostrar la contraseña completa)
	log.Printf("[DEBUG] Intento de login - Email: %s, Password length: %d", 
		credentials.CorreoElectronico, len(credentials.Contrasena))
	
	if credentials.CorreoElectronico == "" || credentials.Contrasena == "" {
		log.Printf("[ERROR] Credenciales vacías - Email: '%s', Password empty: %t", 
			credentials.CorreoElectronico, credentials.Contrasena == "")
		http.Error(w, "Correo electrónico y contraseña requeridos", http.StatusBadRequest)
		return
	}
	
	// DEBUG: Verificar si el usuario existe antes de validar
	usuarioExiste, err := h.repo.GetByEmail(credentials.CorreoElectronico)
	if err != nil {
		log.Printf("[ERROR] Usuario no encontrado con email: %s - Error: %v", 
			credentials.CorreoElectronico, err)
		http.Error(w, "Credenciales inválidas", http.StatusUnauthorized)
		return
	}
	
	log.Printf("[DEBUG] Usuario encontrado - ID: %d, Email: %s, Tipo: %s", 
		usuarioExiste.ID, usuarioExiste.CorreoElectronico, usuarioExiste.TipoUsuario)
	
	usuario, err := h.repo.ValidateLogin(credentials.CorreoElectronico, credentials.Contrasena)
	if err != nil {
		log.Printf("[ERROR] Validación de login falló para email: %s - Error: %v", 
			credentials.CorreoElectronico, err)
		http.Error(w, "Credenciales inválidas", http.StatusUnauthorized)
		return
	}
	
	log.Printf("[DEBUG] Login exitoso para usuario: %s (ID: %d)", 
		usuario.CorreoElectronico, usuario.ID)
	
	// Respuesta exitosa
	response := map[string]interface{}{
		"message": "Login exitoso",
		"usuario": usuario,
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}