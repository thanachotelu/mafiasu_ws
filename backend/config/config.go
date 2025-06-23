package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

// KeycloakConfig เก็บการตั้งค่าสำหรับเชื่อมต่อ Keycloak
type KeycloakConfig struct {
	BaseURL       string // URL ของ Keycloak Server
	Realm         string // ชื่อ Realm
	AdminClientID string // เช่น admin-cli
	AdminUser     string // Admin username
	AdminPass     string // Admin password
	ClientID      string // Client ID สำหรับแอปพลิเคชัน
}

// Config โครงสร้างหลักของการตั้งค่าแอป
type Config struct {
	AppPort           string         // พอร์ตที่แอปจะฟัง
	DatabaseURL       string         // URL สำหรับเชื่อมต่อ Postgres
	KeycloakPublicKey string         // public key สำหรับตรวจสอบ JWT จาก Keycloak
	Keycloak          KeycloakConfig // การตั้งค่า Keycloak ทั้งหมด
}

// New โหลดค่าต่างๆ จาก environment variables ด้วย Viper
func New() (*Config, error) {
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	cfg := &Config{
		AppPort:           viper.GetString("APP_PORT"),
		KeycloakPublicKey: viper.GetString("KEYCLOAK_PUBLIC_KEY"),
		Keycloak: KeycloakConfig{
			BaseURL:       viper.GetString("KEYCLOAK_BASE_URL"),
			Realm:         viper.GetString("KEYCLOAK_REALM"),
			AdminClientID: viper.GetString("KEYCLOAK_ADMIN_CLIENT_ID"),
			AdminUser:     viper.GetString("KEYCLOAK_ADMIN_USER"),
			AdminPass:     viper.GetString("KEYCLOAK_ADMIN_PASS"),
			 ClientID:      viper.GetString("KEYCLOAK_CLIENT_ID"),
		},
	}

	// สร้าง DatabaseURL
	cfg.DatabaseURL = fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		viper.GetString("POSTGRES_HOST"),
		viper.GetString("POSTGRES_PORT"),
		viper.GetString("POSTGRES_USER"),
		viper.GetString("POSTGRES_PASSWORD"),
		viper.GetString("POSTGRES_DB"),
		viper.GetString("POSTGRES_SSLMODE"),
	)

	// Validate required fields
	if cfg.AppPort == "" {
		return nil, fmt.Errorf("APP_PORT is required")
	}
	if cfg.Keycloak.BaseURL == "" || cfg.Keycloak.Realm == "" {
		return nil, fmt.Errorf("KEYCLOAK_BASE_URL and KEYCLOAK_REALM are required")
	}
	if cfg.KeycloakPublicKey == "" {
		return nil, fmt.Errorf("KEYCLOAK_PUBLIC_KEY is required")
	}

	return cfg, nil
}
