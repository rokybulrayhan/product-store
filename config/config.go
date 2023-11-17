package config

import (
	"log"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

const DEFAULT_MAX_OPEN_CONNS = 5

type EnumDimensionType string

const (
	EnumDimensionTypeSize   EnumDimensionType = "Size"
	EnumDimensionTypeWeight EnumDimensionType = "Weight"
)

type (

	// Config -.
	Config struct {
		//JwtToken string `mapstructure:"JWT_TOKEN" json:"JWT_TOKEN"`
		Auth      `mapstructure:",squash"`
		JwtSecret string `env:"JWT_SECRET"`
		//App      `mapstructure:",squash"`
		HTTP      `mapstructure:",squash"`
		GRPC      `mapstructure:",squash"`
		Database  `mapstructure:",squash"`
		Aws       `mapstructure:",squash"`
		Tracing   `mapstructure:",squash"`
		Logger    `mapstructure:",squash"`
		MapApiKey string `env:"MAP_API_KEY"`
		Image     `mapstructure:",squash"`
	}
	Auth struct {
	}
	// App -.
	App struct {
		Name    string `mapstructure:"NAME" json:"name"`
		Version string `mapstructure:"VERSION" json:"version"`
	}

	// HTTP -.
	HTTP struct {
		HTTPAddress string `env:"HTTP_ADDRESS"`
	}
	GRPC struct {
		GrpcPort string `env:"GRPC_PORT"`
	}

	// DB -.
	Database struct {
		DBHost          string `env:"DBHOST"`
		DbUser          string `env:"DBUSER"`
		DbPass          string `env:"DBPASS"`
		DbPort          string `env:"DBPORT"`
		DbName          string `env:"DBNAME"`
		DbSchema        string `env:"DBSCHEMA"`
		SetMaxOpenConns int    `env:"SETMAXOPENCONNS" env-default:"5"`
		Debug           bool   `env:"DEBUG" env-default:"false"`
	}
	// AWS Config
	Aws struct {
		AwsAccessKeyId       string `env:"AWS_ACCESS_KEY_ID"`
		AwsSecretAccessKey   string `env:"AWS_SECRET_ACCESS_KEY"`
		AwsDefaultRegion     string `env:"AWS_DEFAULT_REGION"`
		AwsStorageBucketName string `env:"AWS_STOARGE_BUCKET_NAME"`
		AwsS3BaseURL         string `env:"AWS_S3_BASE_URL"`
		AwsCloudfrontURL     string `env:"AWS_CLOUDFRONT_URL"`
	}
	Tracing struct {
		TracerProvider string `env:"TRACERPROVIDER"`
	}

	// Logger config
	Logger struct {
		Development       bool
		DisableCaller     bool
		DisableStacktrace bool
		Encoding          string
		Level             string
	}

	// Image config
	Image struct {
		ThumbnailWidth  int   `env:"THUMBNAIL_WIDTH"`
		ThumbnailHeight int   `env:"THUMBNAIL_HEIGHT"`
		ImazeMaxSize    int   `env:"IMAGE_MAX_SIZE"`
		VideoMaxSize    int64 `env:"VIDEO_MAX_SIZE"`
	}
)

// type Config struct {
// 	HTTPAddress     string `mapstructure:"HTTP_ADDRESS"`
// 	DBHost          string `mapstructure:"DBHOST"`
// 	DbUser          string `mapstructure:"DBUSER"`
// 	DbPass          string `mapstructure:"DBPASS"`
// 	DbPort          string `mapstructure:"DBPORT"`
// 	DbName          string `mapstructure:"DBNAME"`
// 	DbSchema        string `mapstructure:"DBSCHEMA"`
// 	TracerProvider  string `mapstructure:"TRACERPROVIDER"`
// 	SetMaxOpenConns int    `mapstructure:"SETMAXOPENCONNS"`
// 	Debug           bool   `mapstructure:"DEBUG"`
// }

// Read properties from config.env file
// Command line enviroment variable will overwrite config.env properties
func NewConfig(configFile string) *Config {
	config := Config{}
	godotenv.Load(configFile)
	err := cleanenv.ReadEnv(&config)
	if err != nil {
		log.Fatalln(err)
	}
	return &config
}
