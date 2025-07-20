package config

import (
    "context"
    "fmt"
    "log"
    "os"
    "strings"
    "github.com/jackc/pgx/v5/pgxpool"
    "github.com/joho/godotenv"
)

var Conn *pgxpool.Pool

type Config struct {
    SupabaseURL      string
    SupabaseKey      string
    SupabasePassword string
}

func InitDatabase() error {
    err := godotenv.Load()
    if err != nil {
        return fmt.Errorf("error loading .env file: %v", err)
    }

    log.Println("[DEBUG] Iniciando conexión a base de datos")
    
    // Usar GetDatabaseURL() en lugar de construcción hardcodeada
    connStr := GetDatabaseURL()
    
    // Agregar statement_cache_mode=disabled si no está presente
    if !strings.Contains(connStr, "statement_cache_mode") {
        if strings.Contains(connStr, "?") {
            connStr += "&statement_cache_mode=disabled"
        } else {
            connStr += "?statement_cache_mode=disabled"
        }
    }
    
    log.Printf("[DEBUG] Cadena de conexión: %s", connStr)
    
    poolConfig, err := pgxpool.ParseConfig(connStr)
    if err != nil {
        log.Printf("Error al parsear la configuración: %v", err)
        return err
    }

    Conn, err = pgxpool.NewWithConfig(context.Background(), poolConfig)
    if err != nil {
        log.Printf("Error al conectar a la base de datos: %v", err)
        return err
    }
    
    log.Println("[DEBUG] Conexión a base de datos exitosa")
    return nil
}

func LoadConfig() (*Config, error) {
    err := godotenv.Load()
    if err != nil {
        return nil, fmt.Errorf("error loading .env file: %v", err)
    }

    return &Config{
        SupabaseURL:      os.Getenv("SUPABASE_URL"),
        SupabaseKey:      os.Getenv("SUPABASE_KEY"),
        SupabasePassword: os.Getenv("SUPABASE_PASSWORD"),
    }, nil
}


func GetDatabaseURL() string {
    // Usar variables de entorno
    dbURL := os.Getenv("DATABASE_URL")
    if dbURL != "" {
        // Agregar pool_mode si no está presente
        if !strings.Contains(dbURL, "pool_mode") {
            if strings.Contains(dbURL, "?") {
                dbURL += "&pool_mode=transaction"
            } else {
                dbURL += "?pool_mode=transaction"
            }
        }
        return dbURL
    }
    
    // Fallback a construcción manual
    host := os.Getenv("DB_HOST")
    port := os.Getenv("DB_PORT")
    user := os.Getenv("DB_USER")
    password := os.Getenv("DB_PASSWORD")
    dbname := os.Getenv("DB_NAME")
    
    if host == "" {
        host = "aws-0-us-east-2.pooler.supabase.com"
    }
    if port == "" {
        port = "6543"
    }
    if dbname == "" {
        dbname = "postgres"
    }
    
    return fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?pool_mode=transaction", user, password, host, port, dbname)
}