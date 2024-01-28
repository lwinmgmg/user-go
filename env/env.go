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

type Settings struct {
	HttpServer Server      `yaml:"http_server"`
	GrpcServer Server      `yaml:"grpc_server"`
	Db         DbServer    `yaml:"db"`
	Redis      RedisServer `yaml:"redis"`
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
