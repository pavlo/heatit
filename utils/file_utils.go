package utils

import (
	"github.com/pavlo/heatit/directives"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"net/http"
	"log"
	"fmt"
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

func GetContentForInsertion(insert *directives.InsertDirective) (string, error) {

	fmt.Print("HERE!!!")

	var content string
	var err error

	if insert.Scheme == directives.DefaultScheme {
		content, err = ReadTextFile(insert.SourceValue)
	} else {
		content, err = GetRequest(insert.SourceValue)
	}

	if err != nil {
		log.Printf("Failed to read %s file for insertion!", insert.SourceValue)
		return "", err
	}

	return content, nil
}

func ReadTextFile(filename string) (string, error) {
	b, e := ioutil.ReadFile(filename)
	return string(b), e
}

func GetRequest(url string) (string, error) {

	fmt.Print("HERE")
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return string(body), nil
}

func WriteTextFile(filename string, content []byte) error {
	err := ioutil.WriteFile(filename, content, 0644)
	if err != nil {
		return err
	}
	return nil
}
