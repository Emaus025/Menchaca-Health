package models

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"
)

// CustomDate handles multiple date formats from frontend
type CustomDate struct {
	time.Time
}

// Scan implements the sql.Scanner interface for database reads
func (cd *CustomDate) Scan(value interface{}) error {
	if value == nil {
		cd.Time = time.Time{}
		return nil
	}
	
	switch v := value.(type) {
	case time.Time:
		cd.Time = v
		return nil
	case string:
		// Try to parse string dates
		formats := []string{
			"2006-01-02",
			"2006-01-02 15:04:05",
			"2006-01-02T15:04:05Z",
			"2006-01-02T15:04:05.000Z",
			"2006-01-02T15:04:05-07:00",
		}
		
		for _, format := range formats {
			if parsedTime, err := time.Parse(format, v); err == nil {
				cd.Time = parsedTime
				return nil
			}
		}
		return fmt.Errorf("cannot parse date string: %v", v)
	default:
		return fmt.Errorf("cannot scan %T into CustomDate", value)
	}
}

// Value implements the driver.Valuer interface for database writes
func (cd CustomDate) Value() (driver.Value, error) {
	if cd.Time.IsZero() {
		return nil, nil
	}
	return cd.Time, nil
}

// UnmarshalJSON implements custom JSON unmarshaling for dates
func (cd *CustomDate) UnmarshalJSON(data []byte) error {
	// Remove quotes from JSON string
	var dateStr string
	if err := json.Unmarshal(data, &dateStr); err != nil {
		return err
	}
	
	// Try different date formats
	formats := []string{
		"2006-01-02",                 // YYYY-MM-DD
		"2006-01-02T15:04:05Z",       // ISO format with Z
		"2006-01-02T15:04:05.000Z",   // ISO format with milliseconds
		"2006-01-02T15:04:05-07:00",  // ISO format with timezone
	}
	
	for _, format := range formats {
		if parsedTime, err := time.Parse(format, dateStr); err == nil {
			cd.Time = parsedTime
			return nil
		}
	}
	
	return fmt.Errorf("error parsing date %s: unsupported format", dateStr)
}

// MarshalJSON implements custom JSON marshaling for dates
func (cd CustomDate) MarshalJSON() ([]byte, error) {
	return json.Marshal(cd.Time.Format("2006-01-02"))
}

type Usuario struct {
	ID                 int        `json:"id" db:"id"`
	Nombre             string     `json:"nombre" db:"nombre"`
	Apellido           string     `json:"apellido" db:"apellido"`
	TipoUsuario        string     `json:"tipo_usuario" db:"tipo_usuario"`
	CorreoElectronico  string     `json:"correo_electronico" db:"correo_electronico"`
	Telefono           string     `json:"telefono" db:"telefono"`
	FechaNacimiento    CustomDate `json:"fecha_nacimiento" db:"fecha_nacimiento"`
	Contrasena         string     `json:"contrasena,omitempty" db:"contrasena"`
	EmailVerificado    bool       `json:"email_verificado" db:"email_verificado"`
	TokenVerificacion  string     `json:"-" db:"token_verificacion"`
	TokenExpiracion    time.Time  `json:"-" db:"token_expiracion"`
	FechaCreacion      time.Time  `json:"fecha_creacion" db:"fecha_creacion"`
	FechaActualizacion time.Time  `json:"fecha_actualizacion" db:"fecha_actualizacion"`
}