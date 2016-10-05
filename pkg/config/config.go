package config

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Value1 string
	Value2 string
	Stages []map[string]interface{}
}

func readFile(filename string) ([]byte, error) {
	file, err := os.Open(filename)
	if err != nil {
		return []byte{}, err
	}

	return ioutil.ReadAll(file)
}

func Parse(filename string) (*Config, error) {
	bytes, err := readFile(filename)
	if err != nil {
		return nil, err
	}

	result := &Config{}
	yaml.Unmarshal(bytes, result)

	return result, nil
}
