package yamlfile_test

import (
	"errors"
	"os"
	"path"
	"testing"

	yamlfile "github.com/lwinmgmg/user-go/pkg/yaml_file"
)

func helpWrongFIle(t *testing.T) {
	file_data := `abc=
  aaa: 10
  bbb: abcd`
	tmpDir := os.TempDir()
	filename := "test_load_file0001.yaml"
	path := path.Join(tmpDir, filename)
	file, err := os.Create(path)
	if err != nil {
		t.Errorf("Error on opening yaml test file : %v", err)
		return
	}
	defer os.Remove(path)
	defer file.Close()
	_, err = file.Write([]byte(file_data))
	if err != nil {
		t.Errorf("Error on writing file : %v", err)
		return
	}
	test_data := struct {
		Abc struct {
			Aaa int    `yaml:"aaa"`
			Bbb string `yaml:"bbb"`
		} `yaml:"abc"`
	}{}
	if err := yamlfile.LoadFile(path, &test_data); err == nil {
		t.Error("Expected error on loading wrong format file")
		return
	}
}

func TestLoadFile(t *testing.T) {
	file_data := `abc:
  aaa: 10
  bbb: abcd`
	tmpDir := os.TempDir()
	filename := "test_load_file0001.yaml"
	path := path.Join(tmpDir, filename)
	file, err := os.Create(path)
	if err != nil {
		t.Errorf("Error on opening yaml test file : %v", err)
		return
	}
	defer os.Remove(path)
	defer file.Close()
	_, err = file.Write([]byte(file_data))
	if err != nil {
		t.Errorf("Error on writing file : %v", err)
		return
	}
	test_data := struct {
		Abc struct {
			Aaa int    `yaml:"aaa"`
			Bbb string `yaml:"bbb"`
		} `yaml:"abc"`
	}{}
	if err := yamlfile.LoadFile(path, &test_data); err != nil {
		t.Errorf("Error on loading yaml file : %v", err)
		return
	}
	if test_data.Abc.Aaa != 10 {
		t.Errorf("Expected 10, getting : %v", test_data.Abc.Aaa)
		return
	}
	if test_data.Abc.Bbb != "abcd" {
		t.Errorf("Expected abcd, getting : %v", test_data.Abc.Bbb)
		return
	}
	if err := yamlfile.LoadFile("No file 0000000001", &test_data); !errors.Is(err, os.ErrNotExist) {
		t.Errorf("Expected no file found error : %v", err)
	}
	helpWrongFIle(t)
}
