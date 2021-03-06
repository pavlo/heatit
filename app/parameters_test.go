package app

import (
	"testing"
)

const ValidYaml = "../fixtures/parameters/params.yaml"

func TestNewParameters(t *testing.T) {
	p, err := NewParameters(ValidYaml, make([]string, 0))

	if err != nil {
		t.Errorf("Failed to create an instance of Parameters!")
	}

	if len(p.data) != 6 {
		t.Errorf("Wrong number of parameters!")
	}

	assertParam(p, "private-network-uuid", "00497c93-978b-4ec8-b3f2-7fd0ea738ef4", TypeSimple, t)
	assertParam(p, "network-interface", "eth2", TypeSimple, t)
	assertParam(p, "coreos-token", "954398c993934acf5aedd1315a42d15d", TypeSimple, t)
}

func TestNewParametersWithOverride(t *testing.T) {
	p, err := NewParameters(ValidYaml, []string{"coreos-token=overriden", "a-new-param=NewValue"})

	if err != nil {
		t.Errorf("Failed to create an instance of Parameters!")
	}

	if len(p.data) != 7 {
		t.Errorf("Wrong number of parameters!")
	}

	assertParam(p, "private-network-uuid", "00497c93-978b-4ec8-b3f2-7fd0ea738ef4", TypeSimple, t)
	assertParam(p, "network-interface", "eth2", TypeSimple, t)
	assertParam(p, "coreos-token", "overriden", TypeSimple, t)
	assertParam(p, "a-new-param", "NewValue", TypeSimple, t)
}

func TestNewParametersNoYamlFile(t *testing.T) {
	p, err := NewParameters("", make([]string, 0))
	if err != nil {
		t.Errorf("Failed to create an instance of Parameters!")
	}

	if len(p.data) != 0 {
		t.Errorf("Parameters should be empty!")
	}
}

func TestNewParametersInvalidYamlFile(t *testing.T) {
	_, err := NewParameters("../fixtures/invalid_yaml_file.yaml", make([]string, 0))
	if err == nil {
		t.Errorf("Expected to receive an error, because YAML file is not parseable!")
	}
}

func TestParametersGetValue(t *testing.T) {
	p, _ := NewParameters(ValidYaml, make([]string, 0))

	v, err := p.getValue("network-interface")

	if err != nil {
		t.Errorf("TestGetValue failed, error occured!")
	}
	if v != "eth2" {
		t.Errorf("TestGetValue failed, returned value is wrong!")
	}
}

func TestParametersGetNonExistentValue(t *testing.T) {
	p, _ := NewParameters(ValidYaml, make([]string, 0))

	_, err := p.getValue("does-not-exits")

	if err == nil {
		t.Errorf("TestGetNonExistentValue should fail but it did not!")
	}
}

func assertParam(params *Parameters, name string, value string, paramType int, t *testing.T) {
	p := params.data[name]
	if !p.resolved {
		t.Errorf("%s.resolved is wrong!", name)
	}
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
