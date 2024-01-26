package env

import (
	"os"

	yamlfile "github.com/lwinmgmg/user-go/pkg/yaml_file"
)

type Server struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type Settings struct {
	HttpServer Server `yaml:"http_server"`
	GrpcServer Server `yaml:"grpc_server"`
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
