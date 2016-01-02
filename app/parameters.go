package app

import (
	"github.com/pavlo/heatit/utils"
	"errors"
)

const EMPTY = ""

const (
	TypeSimple int = 0
)

type Parameters struct {
	data map[string] Param
}

type Param struct {
	name string
	value string
	paramType int
}

func NewParameters(yamlFilename string) (*Parameters, error)  {

	result := &Parameters { data: make(map[string]Param) }

	if (yamlFilename == EMPTY) {
		return result, nil
	}

	data, err := utils.ParseYamlFile(yamlFilename)

	if err != nil {
		return nil, err
	}

	for k, v := range data {
		key := k.(string)
		value := v.(string)

		result.data[key] = Param {
			name: key,
			value: value,
			paramType: TypeSimple,
		}
	}

	return result, nil
}

func (params *Parameters) getValue(name string) (value string, err error) {
	p := params.data[name]

	if p.name == EMPTY {
		return EMPTY, errors.New("No value found for key: " + name)
	}

	return p.value, nil
}

