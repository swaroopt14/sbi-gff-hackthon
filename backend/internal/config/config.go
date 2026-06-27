package config

import (
	"fmt"
	"strings"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	App      AppConfig
	Database DatabaseConfig
	Redis    RedisConfig
	JWT      JWTConfig
	Agents   AgentsConfig
	Storage  StorageConfig
	Kafka    KafkaConfig
	CORS     CORSConfig
}

type AppConfig struct {
	Env      string `mapstructure:"env"`
	Port     string `mapstructure:"port"`
	LogLevel string `mapstructure:"log_level"`
}

type DatabaseConfig struct {
	URL          string `mapstructure:"url"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
}

type RedisConfig struct {
	URL      string `mapstructure:"url"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

type JWTConfig struct {
	PrivateKeyPath   string        `mapstructure:"private_key_path"`
	PublicKeyPath    string        `mapstructure:"public_key_path"`
	AccessTokenTTL  time.Duration `mapstructure:"access_token_ttl"`
	RefreshTokenTTL time.Duration `mapstructure:"refresh_token_ttl"`
}

type AgentsConfig struct {
	IntentURL        string        `mapstructure:"intent_url"`
	QualifyURL       string        `mapstructure:"qualify_url"`
	PersonaliseURL   string        `mapstructure:"personalise_url"`
	ConversationURL  string        `mapstructure:"conversation_url"`
	DocumentURL      string        `mapstructure:"document_url"`
	ComplianceURL    string        `mapstructure:"compliance_url"`
	ExplainURL       string        `mapstructure:"explain_url"`
	AuditURL         string        `mapstructure:"audit_url"`
	TimeoutSeconds   int           `mapstructure:"timeout_seconds"`
}

type StorageConfig struct {
	Endpoint   string `mapstructure:"endpoint"`
	AccessKey  string `mapstructure:"access_key"`
	SecretKey  string `mapstructure:"secret_key"`
	Bucket     string `mapstructure:"bucket"`
	UseSSL     bool   `mapstructure:"use_ssl"`
}

type KafkaConfig struct {
	Brokers  []string `mapstructure:"brokers"`
	GroupID  string   `mapstructure:"group_id"`
	UseMock  bool     `mapstructure:"use_mock"`
}

type CORSConfig struct {
	AllowedOrigins []string `mapstructure:"allowed_origins"`
}

func Load() (*Config, error) {
	v := viper.New()

	v.SetConfigName("app")
	v.SetConfigType("yaml")
	v.AddConfigPath("./configs")
	v.AddConfigPath("../configs")

	v.SetEnvPrefix("APERTURE")
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()

	setDefaults(v)

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, fmt.Errorf("reading config: %w", err)
		}
	}

	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("unmarshalling config: %w", err)
	}

	return &cfg, nil
}

func setDefaults(v *viper.Viper) {
	v.SetDefault("app.env", "development")
	v.SetDefault("app.port", "8080")
	v.SetDefault("app.log_level", "info")

	v.SetDefault("database.max_open_conns", 25)
	v.SetDefault("database.max_idle_conns", 10)

	v.SetDefault("redis.db", 0)

	v.SetDefault("jwt.access_token_ttl", 15*time.Minute)
	v.SetDefault("jwt.refresh_token_ttl", 7*24*time.Hour)

	v.SetDefault("agents.timeout_seconds", 30)
	v.SetDefault("agents.intent_url", "http://localhost:8101")
	v.SetDefault("agents.qualify_url", "http://localhost:8102")
	v.SetDefault("agents.personalise_url", "http://localhost:8103")
	v.SetDefault("agents.conversation_url", "http://localhost:8104")
	v.SetDefault("agents.document_url", "http://localhost:8105")
	v.SetDefault("agents.compliance_url", "http://localhost:8106")
	v.SetDefault("agents.explain_url", "http://localhost:8107")
	v.SetDefault("agents.audit_url", "http://localhost:8108")

	v.SetDefault("kafka.use_mock", true)
	v.SetDefault("kafka.group_id", "aperture-backend")

	v.SetDefault("cors.allowed_origins", []string{"http://localhost:3000"})
}
