package database

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
	"time"
	"menchaca-health/internal/models"
	"golang.org/x/crypto/argon2"
	"github.com/jackc/pgx/v5"
)

type UsuarioRepository struct {
	db *Database
}

func NewUsuarioRepository(db *Database) *UsuarioRepository {
	return &UsuarioRepository{db: db}
}

// HashPassword cifra la contraseña usando Argon2
func (r *UsuarioRepository) HashPassword(contrasena string) (string, error) {
	salt := make([]byte, 16)
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}
	
	// Parámetros Argon2: time=1, memory=64MB, threads=4, keyLen=32
	hash := argon2.IDKey([]byte(contrasena), salt, 1, 64*1024, 4, 32)
	
	// Combinar salt + hash y codificar en base64
	combined := append(salt, hash...)
	return base64.StdEncoding.EncodeToString(combined), nil
}

// VerifyPassword verifica si la contraseña coincide con el hash
func (r *UsuarioRepository) VerifyPassword(contrasena, contrasenaHasheada string) bool {
	decoded, err := base64.StdEncoding.DecodeString(contrasenaHasheada)
	if err != nil || len(decoded) < 48 {
		return false
	}
	
	salt := decoded[:16]
	hash := decoded[16:]
	
	newHash := argon2.IDKey([]byte(contrasena), salt, 1, 64*1024, 4, 32)
	
	return string(hash) == string(newHash)
}

// Create crea un nuevo usuario
func (r *UsuarioRepository) Create(usuario *models.Usuario) error {
	hashedPassword, err := r.HashPassword(usuario.Contrasena)
	if err != nil {
		return fmt.Errorf("error al cifrar contraseña: %v", err)
	}
	
	query := `
		INSERT INTO usuarios (nombre, apellido, tipo_usuario, correo_electronico, telefono, fecha_nacimiento, contrasena)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id`
	
	err = r.db.DB.QueryRow(context.Background(), query,
		usuario.Nombre,
		usuario.Apellido,
		usuario.TipoUsuario,
		usuario.CorreoElectronico,
		usuario.Telefono,
		usuario.FechaNacimiento.Time, // Use .Time to get the underlying time.Time
		hashedPassword,
	).Scan(&usuario.ID)
	
	if err != nil {
		return fmt.Errorf("error al crear usuario: %v", err)
	}
	
	return nil
}

// GetByID obtiene un usuario por ID
func (r *UsuarioRepository) GetByID(id int) (*models.Usuario, error) {
	usuario := &models.Usuario{}
	query := `
		SELECT id, nombre, apellido, tipo_usuario, correo_electronico, telefono, fecha_nacimiento
		FROM usuarios WHERE id = $1`
	
	err := r.db.DB.QueryRow(context.Background(), query, id).Scan(
		&usuario.ID,
		&usuario.Nombre,
		&usuario.Apellido,
		&usuario.TipoUsuario,
		&usuario.CorreoElectronico,
		&usuario.Telefono,
		&usuario.FechaNacimiento,
	)
	
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("usuario no encontrado")
		}
		return nil, fmt.Errorf("error al obtener usuario: %v", err)
	}
	
	return usuario, nil
}

// Update actualiza un usuario
func (r *UsuarioRepository) Update(usuario *models.Usuario) error {
	var query string
	var args []interface{}
	
	// Convertir la fecha a solo YYYY-MM-DD para PostgreSQL DATE
	fechaNacimientoDate := usuario.FechaNacimiento.Format("2006-01-02")
	
	if usuario.Contrasena != "" {
		// Si se proporciona nueva contraseña, cifrarla
		hashedPassword, err := r.HashPassword(usuario.Contrasena)
		if err != nil {
			return fmt.Errorf("error al cifrar contraseña: %v", err)
		}
		
		query = `
			UPDATE usuarios 
			SET nombre = $1, apellido = $2, tipo_usuario = $3, correo_electronico = $4, 
			    telefono = $5, fecha_nacimiento = $6, contrasena = $7
			WHERE id = $8`
		
		args = []interface{}{
			usuario.Nombre,
			usuario.Apellido,
			usuario.TipoUsuario,
			usuario.CorreoElectronico,
			usuario.Telefono,
			fechaNacimientoDate, // Usar string formateado
			hashedPassword,
			usuario.ID,
		}
	} else {
		// Sin cambio de contraseña
		query = `
			UPDATE usuarios 
			SET nombre = $1, apellido = $2, tipo_usuario = $3, correo_electronico = $4, 
			    telefono = $5, fecha_nacimiento = $6
			WHERE id = $7`
		
		args = []interface{}{
			usuario.Nombre,
			usuario.Apellido,
			usuario.TipoUsuario,
			usuario.CorreoElectronico,
			usuario.Telefono,
			fechaNacimientoDate, // Usar string formateado
			usuario.ID,
		}
	}
	
	result, err := r.db.DB.Exec(context.Background(), query, args...)
	if err != nil {
		return fmt.Errorf("error al actualizar usuario: %v", err)
	}
	
	if result.RowsAffected() == 0 {
		return ErrNoRecord
	}
	
	return nil
}

// Delete elimina un usuario
func (r *UsuarioRepository) Delete(id int) error {
	query := `DELETE FROM usuarios WHERE id = $1`
	
	result, err := r.db.DB.Exec(context.Background(), query, id)
	if err != nil {
		return fmt.Errorf("error al eliminar usuario: %v", err)
	}
	
	if result.RowsAffected() == 0 {
		return ErrNoRecord
	}
	
	return nil
}

// List obtiene todos los usuarios
func (r *UsuarioRepository) List() ([]*models.Usuario, error) {
	query := `
		SELECT id, nombre, apellido, tipo_usuario, correo_electronico, telefono, fecha_nacimiento
		FROM usuarios ORDER BY nombre, apellido`
	
	rows, err := r.db.DB.Query(context.Background(), query)
	if err != nil {
		return nil, fmt.Errorf("error al listar usuarios: %v", err)
	}
	defer rows.Close()
	
	var usuarios []*models.Usuario
	for rows.Next() {
		usuario := &models.Usuario{}
		err := rows.Scan(
			&usuario.ID,
			&usuario.Nombre,
			&usuario.Apellido,
			&usuario.TipoUsuario,
			&usuario.CorreoElectronico,
			&usuario.Telefono,
			&usuario.FechaNacimiento,
		)
		if err != nil {
			return nil, fmt.Errorf("error al escanear usuario: %v", err)
		}
		usuarios = append(usuarios, usuario)
	}
	
	if err = rows.Err(); err != nil {
		return nil, err
	}
	
	return usuarios, nil
}

// GetByRole obtiene usuarios por tipo/rol
func (r *UsuarioRepository) GetByRole(role string) ([]*models.Usuario, error) {
	query := `
		SELECT id, nombre, apellido, tipo_usuario, correo_electronico, telefono, fecha_nacimiento
		FROM usuarios WHERE tipo_usuario = $1 ORDER BY nombre, apellido`
	
	rows, err := r.db.DB.Query(context.Background(), query, role)
	if err != nil {
		return nil, fmt.Errorf("error al obtener usuarios por rol: %v", err)
	}
	defer rows.Close()
	
	var usuarios []*models.Usuario
	for rows.Next() {
		usuario := &models.Usuario{}
		err := rows.Scan(
			&usuario.ID,
			&usuario.Nombre,
			&usuario.Apellido,
			&usuario.TipoUsuario,
			&usuario.CorreoElectronico,
			&usuario.Telefono,
			&usuario.FechaNacimiento,
		)
		if err != nil {
			return nil, fmt.Errorf("error al escanear usuario: %v", err)
		}
		usuarios = append(usuarios, usuario)
	}
	
	if err = rows.Err(); err != nil {
		return nil, err
	}
	
	return usuarios, nil
}

// ValidateLogin valida las credenciales de login
func (r *UsuarioRepository) ValidateLogin(correoElectronico, contrasena string) (*models.Usuario, error) {
	log.Printf("[DEBUG] ValidateLogin - Buscando usuario con email: %s", correoElectronico)
	
	usuario, err := r.GetByEmail(correoElectronico)
	if err != nil {
		log.Printf("[ERROR] GetByEmail falló para %s: %v", correoElectronico, err)
		return nil, err
	}
	
	log.Printf("[DEBUG] Usuario encontrado en BD - ID: %d, Email: %s", 
		usuario.ID, usuario.CorreoElectronico)
	
	// Verificar contraseña
	passwordValid := r.VerifyPassword(contrasena, usuario.Contrasena)
	log.Printf("[DEBUG] Verificación de contraseña para %s: %t", 
		correoElectronico, passwordValid)
	
	if !passwordValid {
		log.Printf("[ERROR] Contraseña incorrecta para usuario: %s", correoElectronico)
		return nil, fmt.Errorf("credenciales inválidas")
	}
	
	// Limpiar contraseña antes de retornar
	usuario.Contrasena = ""
	log.Printf("[DEBUG] Login validado exitosamente para: %s", correoElectronico)
	return usuario, nil
}

// GetByEmail obtiene un usuario por email (para login)
func (r *UsuarioRepository) GetByEmail(correoElectronico string) (*models.Usuario, error) {
	log.Printf("[DEBUG] Buscando usuario con email: '%s'", correoElectronico)
	
	usuario := &models.Usuario{}
	query := `
		SELECT id, nombre, apellido, tipo_usuario, correo_electronico, telefono, fecha_nacimiento, contrasena
		FROM usuarios WHERE correo_electronico = $1`
	
	log.Printf("[DEBUG] Ejecutando query: %s con parámetro: '%s'", query, correoElectronico)
	
	err := r.db.DB.QueryRow(context.Background(), query, correoElectronico).Scan(
		&usuario.ID,
		&usuario.Nombre,
		&usuario.Apellido,
		&usuario.TipoUsuario,
		&usuario.CorreoElectronico,
		&usuario.Telefono,
		&usuario.FechaNacimiento,
		&usuario.Contrasena,
	)
	
	if err != nil {
		if err == pgx.ErrNoRows {
			log.Printf("[DEBUG] No se encontró usuario con email: '%s'", correoElectronico)
			return nil, fmt.Errorf("usuario no encontrado")
		}
		log.Printf("[ERROR] Error en query GetByEmail: %v", err)
		return nil, fmt.Errorf("error al obtener usuario: %v", err)
	}
	
	log.Printf("[DEBUG] Usuario encontrado exitosamente: ID=%d, Email='%s'", usuario.ID, usuario.CorreoElectronico)
	return usuario, nil
}

// IsEmailVerified verifica si el email está verificado
func (r *UsuarioRepository) IsEmailVerified(correoElectronico string) (bool, error) {
    var verificado bool
    query := "SELECT email_verificado FROM usuarios WHERE correo_electronico = $1"
    
    err := r.db.DB.QueryRow(context.Background(), query, correoElectronico).Scan(&verificado)
    if err != nil {
        return false, err
    }
    
    return verificado, nil
}

func (r *UsuarioRepository) CreateWithVerification(usuario *models.Usuario, token string) error {
	log.Println("[DEBUG] Iniciando CreateWithVerification")
	
	hashedPassword, err := r.HashPassword(usuario.Contrasena)
	if err != nil {
		log.Printf("[ERROR] Error hasheando password: %v", err)
		return err
	}
	log.Println("[DEBUG] Password hasheado exitosamente")

	// Convertir la fecha a solo YYYY-MM-DD para PostgreSQL DATE
	fechaNacimientoDate := usuario.FechaNacimiento.Format("2006-01-02")
	log.Printf("[DEBUG] Fecha formateada: %s", fechaNacimientoDate)

	query := `
		INSERT INTO usuarios (nombre, apellido, tipo_usuario, correo_electronico, telefono, 
						fecha_nacimiento, contrasena, email_verificado, token_verificacion, 
						token_expiracion, fecha_creacion, fecha_actualizacion)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
		RETURNING id
	`
	
	tokenExpiration := time.Now().Add(24 * time.Hour)
	now := time.Now()
	
	log.Printf("[DEBUG] Ejecutando query INSERT con email: %s", usuario.CorreoElectronico)
	log.Printf("[DEBUG] Token expiration: %v", tokenExpiration)
	
	err = r.db.DB.QueryRow(
		context.Background(),
		query,
		usuario.Nombre,
		usuario.Apellido,
		usuario.TipoUsuario,
		usuario.CorreoElectronico,
		usuario.Telefono,
		fechaNacimientoDate, // <-- Usar la fecha formateada
		hashedPassword,
		false, // email_verificado
		token,
		tokenExpiration,
		now,
		now,
	).Scan(&usuario.ID)

	if err != nil {
		log.Printf("[ERROR] Error ejecutando INSERT: %v", err)
		log.Printf("[ERROR] Query: %s", query)
		return err
	}
	
	log.Printf("[DEBUG] Usuario insertado exitosamente con ID: %d", usuario.ID)
	return nil
}

func (r *UsuarioRepository) VerifyEmail(token string) error {
    query := `
        UPDATE usuarios 
        SET email_verificado = true, 
            token_verificacion = NULL, 
            token_expiracion = NULL,
            fecha_actualizacion = $1
        WHERE token_verificacion = $2 
        AND token_expiracion > $3
        AND email_verificado = false
    `
    
    result, err := r.db.DB.Exec(
        context.Background(),
        query,
        time.Now(),
        token,
        time.Now(),
    )
    
    if err != nil {
        return err
    }
    
    rowsAffected := result.RowsAffected()
    if rowsAffected == 0 {
        return fmt.Errorf("token inválido o expirado")
    }
    
    return nil
}