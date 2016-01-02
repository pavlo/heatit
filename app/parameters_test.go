package app

import (
	"testing"
)

func TestNewParameters(t *testing.T) {
	p, err := NewParameters("../fixtures/params.yaml")

	if err != nil {
		t.Errorf("Failed to create an instance of Parameters!")
	}

	if len(p.data) != 3 {
		t.Errorf("Wrong number of parameters!")
	}

	assertCorrectValues(p, "private-network-uuid", "00497c93-978b-4ec8-b3f2-7fd0ea738ef4", TypeSimple, t)
	assertCorrectValues(p, "network-interface", "eth2", TypeSimple, t)
	assertCorrectValues(p, "coreos-token", "954398c993934acf5aedd1315a42d15d", TypeSimple, t)
}

func TestNewParametersNoYamlFile(t *testing.T) {
	p, err := NewParameters("")
	if err != nil {
		t.Errorf("Failed to create an instance of Parameters!")
	}

	if len(p.data) != 0 {
		t.Errorf("Parameters should be empty!")
	}
}

func TestNewParametersInvalidYamlFile(t *testing.T) {
	_, err := NewParameters("../fixtures/invalid-yaml-file.yaml")
	if err == nil {
		t.Errorf("Expected to receive an error, because YAML file is not parseable!")
	}
}

func TestGetValue(t *testing.T) {
	p, _ := NewParameters("../fixtures/params.yaml")

	v, err := p.getValue("network-interface")

	if err != nil {
		t.Errorf("TestGetValue failed, error occured!")
	}
	if v != "eth2" {
		t.Errorf("TestGetValue failed, returned value is wrong!")
	}
}

func TestGetNonExistentValue(t *testing.T) {
	p, _ := NewParameters("../fixtures/params.yaml")

	_, err := p.getValue("does-not-exits")

	if err == nil {
		t.Errorf("TestGetNonExistentValue should fail but it did not!")
	}
}

func assertCorrectValues(params *Parameters, name string, value string, paramType int, t *testing.T) {
	p := params.data[name]
	if p.name != name {
		t.Errorf("%s.name is wrong!", name)
	}
	if p.value != value {
		t.Errorf("%s.value is wrong!", name)
	}
	if p.paramType != paramType {
		t.Errorf("%s.paramType is wrong!", name)
	}
}