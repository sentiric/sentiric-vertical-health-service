// sentiric-vertical-health-service/internal/config/config.go
package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

type Config struct {
	GRPCPort string
	HttpPort string
	CertPath string
	KeyPath  string
	CaPath   string
	LogLevel string
	Env      string

	// Sağlık servisi bağımlılıkları (Placeholder)
	EMRAdapter string // Elektronik Tıbbi Kayıt sistemi (Epic, Cerner, vb.)
}

func Load() (*Config, error) {
	godotenv.Load()

	// Harmonik Mimari Portlar (Dikey Servisler, 202XX bloğu atandı)
	return &Config{
		GRPCPort: GetEnv("VERTICAL_HEALTH_SERVICE_GRPC_PORT", "20211"),
		HttpPort: GetEnv("VERTICAL_HEALTH_SERVICE_HTTP_PORT", "20210"),

		CertPath: GetEnvOrFail("VERTICAL_HEALTH_SERVICE_CERT_PATH"),
		KeyPath:  GetEnvOrFail("VERTICAL_HEALTH_SERVICE_KEY_PATH"),
		CaPath:   GetEnvOrFail("GRPC_TLS_CA_PATH"),
		LogLevel: GetEnv("LOG_LEVEL", "info"),
		Env:      GetEnv("ENV", "production"),

		EMRAdapter: GetEnv("EMR_ADAPTER", "mock_emr"),
	}, nil
}

func GetEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

func GetEnvOrFail(key string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		log.Fatal().Str("variable", key).Msg("Gerekli ortam değişkeni tanımlı değil")
	}
	return value
}
