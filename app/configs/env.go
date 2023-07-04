package configs

import (
	"github.com/caarlos0/env/v8"
	_ "github.com/joho/godotenv/autoload"
)

type Environment string

const (
	Development Environment = "development"
	Test        Environment = "test"
	Staging     Environment = "staging"
	Production  Environment = "production"
)

type Config struct {
	UP   bool        `env:"UP" envDefault:"false"`
	Env  Environment `env:"ENV" envDefault:"development"`
	DIR  string      `env:"DIR" envDefault:"/Users/anuragupadhyay/Desktop/Anurag/work/rampnow-be"`
	Port int         `env:"PORT" envDefault:"3000"`

	DBHost     string `env:"DB_HOST" envDefault:"smart-blob-2709.7s5.cockroachlabs.cloud"`
	DBPort     int    `env:"DB_PORT" envDefault:"26257"`
	DBUser     string `env:"DB_USER" envDefault:"anurag"`
	DBPassword string `env:"DB_PASSWORD" envDefault:"zIJRTuQmeHcRyYN7xNdLUg"`
	DBName     string `env:"DB_NAME" envDefault:"defaultdb"`
	DBSSLMode  string `env:"DB_SSL_MODE" envDefault:"verify-full"`
	DBUrl      string `env:"DB_URL" envDefault:"root@tcp(127.0.0.1:3306)/rapid?parseTime=true"`

	JwtSecret string `env:"JWT_SECRET" envDefault:"secret"`

	RedisHost     string `env:"REDIS_HOST" envDefault:"127.0.0.1"`
	RedisPort     string `env:"REDIS_PORT" envDefault:"6379"`
	RedisDB       int    `env:"REDIS_DB" envDefault:"1"`
	RedisPassword string `env:"REDIS_PASSWORD" envDefault:""`
	RedisMaxRetry int    `env:"REDIS_MAX_RETRIES" envDefault:"10"`
	RedisPoolSize int    `env:"REDIS_POOL_SIZE" envDefault:"5"`

	Region           string `env:"REGION" envDefault:"eu-north-1"`
	AccessKeyID      string `env:"ACCESS_KEY_ID" envDefault:"AJBHSYSAJGASGAGA"`
	SecretAccessKey  string `env:"SECRET_ACCESS_KEY" envDefault:"bhajdbajbajhbasdbasdasdabsjdbasjdash"`
	DocumentS3Bucket string `env:"DOCUMENT_S3_BUCKET" envDefault:"docs-rampnow"`

	SmtpHost     string `env:"SMTP_PORT" envDefault:"email-smtp.eu-north-1.amazonaws.com"`
	SmtpPort     string `env:"SMTP_ENDPOINT" envDefault:"2465"`
	SmtpUsername string `env:"SMTP_USERNAME" envDefault:"ieijadlnieahfai"`
	SmtpPassword string `env:"SMTP_PASSWORD" envDefault:"dbaudegdahcaucnachacheichehc"`

	WorkerRedisDB int `env:"WORKER_REDIS_DB" envDefault:"2"`
}

func (c *Config) Load() error {
	if err := env.Parse(c); err != nil {
		return err
	}
	return nil
}
