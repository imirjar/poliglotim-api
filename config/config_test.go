// config/config_test.go
package config

import (
	"os"
	"testing"
)

func TestConfig_New(t *testing.T) {
	// Создаем временный конфиг файл с правильным YAML
	yamlContent := `server:
  port: "8080"
database:
  name: "test_db"
  user: "test_user"
  password: "test_pass"
  host: "localhost"
  port: "5432"
`

	tmpFile, err := os.CreateTemp("", "test_config_*.yaml")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpFile.Name())

	if _, err := tmpFile.WriteString(yamlContent); err != nil {
		t.Fatal(err)
	}
	if err := tmpFile.Close(); err != nil {
		t.Fatal(err)
	}

	// Тестируем загрузку из файла
	cfg := &Config{}
	err = cfg.readFile(tmpFile.Name())
	if err != nil {
		t.Fatalf("readFile() error = %v", err)
	}

	if cfg.Server.Port != "8080" {
		t.Errorf("Expected server port 8080, got '%s'", cfg.Server.Port)
	}

	if cfg.Database.Name != "test_db" {
		t.Errorf("Expected db name test_db, got '%s'", cfg.Database.Name)
	}
}

func TestConfig_ReadEnv(t *testing.T) {
	// Устанавливаем env переменные
	t.Setenv("SERVER_PORT", "9090")
	t.Setenv("DB_NAME", "env_db")
	t.Setenv("DB_USER", "env_user")

	cfg := &Config{
		Server:   ServerConf{Port: "8080"}, // значение по умолчанию
		Database: StorageConf{Name: "default_db"},
	}

	err := cfg.readEnv()
	if err != nil {
		t.Errorf("readEnv() error = %v", err)
	}

	// Env должен переопределить значения
	if cfg.Server.Port != "9090" {
		t.Errorf("Expected server port 9090 from env, got '%s'", cfg.Server.Port)
	}

	if cfg.Database.Name != "env_db" {
		t.Errorf("Expected db name env_db from env, got '%s'", cfg.Database.Name)
	}
}

func TestStorageConf_GetConnString(t *testing.T) {
	tests := []struct {
		name string
		db   StorageConf
		want string
	}{
		{
			name: "basic connection string",
			db: StorageConf{
				Name: "mydb",
				User: "user",
				Pswd: "pass",
				Host: "localhost",
				Port: "5432",
			},
			want: "postgresql://user:pass@localhost:5432/mydb?sslmode=disable",
		},
		{
			name: "empty password",
			db: StorageConf{
				Name: "mydb",
				User: "user",
				Pswd: "",
				Host: "localhost",
				Port: "5432",
			},
			want: "postgresql://user:@localhost:5432/mydb?sslmode=disable",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.db.GetConnString(); got != tt.want {
				t.Errorf("GetConnString() = %v, want %v", got, tt.want)
			}
		})
	}
}
