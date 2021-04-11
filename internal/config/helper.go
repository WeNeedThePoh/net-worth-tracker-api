package config

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/validator.v2"
	"gopkg.in/yaml.v2"
)

var configs Conf

// Conf yaml struct for configs
type Conf struct {
	Serve struct {
		Public struct {
			Port int64 `validate:"nonzero"`
		}
	}
	Log struct {
		Level string `validate:"nonzero"`
	}
	Db struct {
		Host     string `validate:"nonzero"`
		Port     int64  `validate:"nonzero"`
		Name     string `validate:"nonzero"`
		User     string `validate:"nonzero"`
		Password string `validate:"nonzero"`
		Ssl      string `validate:"nonzero"`
	}
}

// InitFromFile Parse and store configs from file
func InitFromFile(filePath string) error {
	buf, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	parsedConfig, err := ParseFile(buf)
	if err != nil {
		return fmt.Errorf("in file %q: %v", ".config.yaml", err)
	}

	configs = *parsedConfig
	return nil
}

// ParseFile Parse YAML configuration file and validate fields
func ParseFile(buffer []byte) (*Conf, error) {
	c := &Conf{}

	err := yaml.Unmarshal(buffer, c)
	if err != nil {
		return nil, err
	}

	err = validateFields(*c)
	if err != nil {
		return nil, err
	}

	return c, nil
}

// GetAll Get parsed configs
func GetAll() Conf {
	return configs
}

func validateFields(c Conf) error {
	err := validator.Validate(c)

	return err
}
