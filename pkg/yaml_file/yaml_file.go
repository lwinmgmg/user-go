package yamlfile

import (
	"os"

	"gopkg.in/yaml.v3"
)

func LoadFile(path string, out any) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	if err := yaml.Unmarshal(data, out); err != nil {
		return err
	}
	return nil
}
