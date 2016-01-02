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
