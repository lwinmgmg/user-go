package env

import (
	"os"

	yamlfile "github.com/lwinmgmg/user-go/pkg/yaml_file"
)

type Server struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type RedisServer struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Db       int    `yaml:"db"`
}

type DbServer struct {
	Host        string `yaml:"host"`
	Port        int    `yaml:"port"`
	User        string `yaml:"user"`
	Password    string `yaml:"password"`
	DbName      string `yaml:"db_name"`
	TablePrefix string `yaml:"table_prefix"`
}

type EmailServer struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Email    string `yaml:"email"`
	Password string `yaml:"password"`
	Enable   bool   `yaml:"enable"`
}

type OtpService struct {
	Issuer string `yaml:"name"`
	Skew   uint   `yaml:"skew"`
}

type JwtService struct {
	Key           string `yaml:"key"`
	CacheDuration uint   `yaml:"cache_duration"`
	LoginDuration uint   `yaml:"login_duration"`
}

type Settings struct {
	Service          string      `yaml:"service"`
	HttpServer       Server      `yaml:"http_server"`
	GrpcServer       Server      `yaml:"grpc_server"`
	Db               DbServer    `yaml:"db"`
	RoDb             DbServer    `yaml:"ro_db"`
	Redis            RedisServer `yaml:"redis"`
	LoginEmailServer EmailServer `yaml:"login_mail_server"`
	OtpService       OtpService  `yaml:"otp"`
	JwtService       JwtService  `yaml:"jwt"`
}

func LoadSettings() (Settings, error) {
	settings := Settings{}
	path, ok := os.LookupEnv("USER_SETTING_PATH")
	if !ok {
		path = "settings.yaml"
	}
	if err := yamlfile.LoadFile(path, &settings); err != nil {
		return settings, err
	}
	return settings, nil
}
