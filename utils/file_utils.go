package utils

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

func ParseYamlFile(path string) (map[interface{}]interface{}, error) {
	var data map[interface{}]interface{}
	source, err := ioutil.ReadFile(path)

	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(source, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func ReadTextFile(filename string) (string, error) {
	b, e := ioutil.ReadFile(filename)
	return string(b), e
}

func WriteTextFile(filename string, content []byte) error {
	err := ioutil.WriteFile(filename, content, 0644)
	if err != nil {
		return err
	}
	return nil
}
