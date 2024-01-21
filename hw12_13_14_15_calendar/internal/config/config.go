package config

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/artembert/golang-pro-otus-hw/hw12_13_14_15_calendar/internal/interfaces/loglevel"
	"gopkg.in/yaml.v3"
)

type LoggerConf struct {
	Level      loglevel.LogLevel `yaml:"level"`
	OutputPath string            `yaml:"outputPath"`
}

type StorageConf struct {
	Type string `yaml:"type"`
}

type DBConf struct {
	Host         string `yaml:"host"`
	Port         string `yaml:"port"`
	DBName       string `yaml:"dbName"`
	User         string `yaml:"user"`
	Password     string `yaml:"password"`
	MigrationDir string `yaml:"migrationDir"`
}

type ServerConf struct {
	Host              string        `yaml:"host"`
	Port              string        `yaml:"port"`
	ReadHeaderTimeout time.Duration `yaml:"readHeaderTimeout"`
	ReadTimeout       time.Duration `yaml:"readTimeout"`
	WriteTimeout      time.Duration `yaml:"writeTimeout"`
}

type Config struct {
	Logger  LoggerConf
	Storage StorageConf
	DB      DBConf
	Server  ServerConf
}

func New(configFilePath string) (Config, error) {
	cfg := Config{}

	file, err := os.Open(configFilePath)
	if err != nil {
		return cfg, fmt.Errorf("failed to open config gile %w", err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Fatalf("Unable to close config file %s", err)
		}
	}()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return cfg, fmt.Errorf("failed to open config gile %w", err)
	}

	if err := yaml.Unmarshal(bytes, &cfg); err != nil {
		_ = fmt.Errorf("Failed to decode config file: %w\n", err)
		return Config{}, err
	}

	return cfg, nil
}

func (dbConfig *DBConf) BuildDBUrl() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%v/%s?sslmode=disable",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
}
