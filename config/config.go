package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	dotenv "github.com/dotenv-org/godotenvvault"
	"github.com/joeshaw/envdecode"
)

type envars struct {
	CSRFKey     string `env:"CSRF_KEY,required"`
	BindAddress string `env:"BIND_ADDRESS,default=127.0.0.1:3000,strict"`
	Database    struct {
		Host        string `env:"DB_HOST,default=127.0.0.1,strict"`
		Port        string `env:"DB_PORT,default=5432,strict"`
		Name        string `env:"DB_NAME,required"`
		Username    string `env:"DB_USERNAME,required"`
		Password    string `env:"DB_PASSWORD,required"`
		AutoMigrate bool   `env:"DATABASE_AUTO_MIGRATE,default=false,strict"`
	}
	Email struct {
		Provider     string `env:"EMAIL_PROVIDER"` // possible values: smtp, postmark, awsses
		MailFromName string `env:"EMAIL_FROM_NAME,default=Administrator,strict"`
		MailFromAddr string `env:"EMAIL_FROM_ADDRESS,default=admin@example.com,strict"`
		SMTP         struct {
			Host           string `env:"EMAIL_SMTP_HOST"`
			Port           string `env:"EMAIL_SMTP_PORT"`
			Username       string `env:"EMAIL_SMTP_USERNAME"`
			Password       string `env:"EMAIL_SMTP_PASSWORD"`
			EnableStartTLS bool   `env:"EMAIL_SMTP_ENABLE_TLS,default=true"`
		}
		Postmark struct {
			APIKey string `env:"EMAIL_POSTMARK_API_KEY"`
		}
		AWSSES struct {
			Region          string `env:"EMAIL_AWSSES_REGION"`
			AccessKeyID     string `env:"EMAIL_AWSSES_ACCESS_KEY_ID"`
			SecretAccessKey string `env:"EMAIL_AWSSES_SECRET_ACCESS_KEY"`
		}
	}
}

// Env is a strongly typed reference to all configuration parsed from Environment Variables
var Env envars

func init() {
	loadEnvVariables()
	Reload()
}

// Reload configuration from current Enviornment Variables
func Reload() {
	Env = envars{}
	err := envdecode.Decode(&Env)
	if err != nil {
		log.Fatalf("failed to parse envars: %s", err)
		panic(err)
	}

	// Email Provider can be inferred if absense
	if Env.Email.Provider == "" {
		if Env.Email.Postmark.APIKey != "" {
			Env.Email.Provider = "postmark"
		} else if Env.Email.AWSSES.AccessKeyID != "" {
			Env.Email.Provider = "awsses"
		} else {
			Env.Email.Provider = "smtp"
		}
	}

	emailType := Env.Email.Provider
	if emailType == "postmark" {
		mustBeSet("EMAIL_POSTMARK_API_KEY")
	} else if emailType == "awsses" {
		mustBeSet("EMAIL_AWSSES_REGION")
		mustBeSet("EMAIL_AWSSES_ACCESS_KEY_ID")
		mustBeSet("EMAIL_AWSSES_SECRET_ACCESS_KEY")
	} else if emailType == "smtp" {
		mustBeSet("EMAIL_SMTP_HOST")
		mustBeSet("EMAIL_SMTP_PORT")
	}
}

func mustBeSet(key string) {
	_, ok := os.LookupEnv(key)
	if !ok {
		panic(fmt.Errorf("could not find environment variable named '%s'", key))
	}
}

func loadEnvVariables() {
	// Get the current working directory
	wd, err := os.Getwd()
	if err != nil {
		log.Println("error getting current working directory:", err)
		return
	}

	// Determine the paths to .env files in current and parent directories
	envFiles := []string{filepath.Join(".", ".env")}

	// Attempt to load environment variables from .env files
	for _, envFile := range envFiles {
		err := dotenv.Load(filepath.Join(wd, envFile))
		if err == nil {
			// If successfully loaded, break the loop
			break
		}
	}

	if err != nil {
		log.Fatalf("error loading .env file %s", err.Error())
	}
}
