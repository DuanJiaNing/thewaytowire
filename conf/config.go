package conf

import (
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type File *os.File

type Config struct {
	Addr string `yaml:"addr"`
}

type ConfigFilePath string

func Load(cp ConfigFilePath) (*Config, error) {
	path := string(cp)
	log.Println("read config from:", path)

	bs, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cf Config
	if err = yaml.Unmarshal(bs, &cf); err != nil {
		return nil, err
	}

	return &cf, nil
}
