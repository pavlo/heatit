package app

import (
	"errors"
	"fmt"
	"github.com/pavlo/heatit/utils"
	"strings"
)

const (
	TypeSimple int = 0
)

type Parameters struct {
	data map[string]Param
}

type Param struct {
	name      string
	value     string
	paramType int
	resolved  bool
}

func NewParameters(yamlFilename string, overrides []string) (*Parameters, error) {

	result := &Parameters{data: make(map[string]Param)}

	if yamlFilename == Empty {
		return result, nil
	}

	data, err := utils.ParseYamlFile(yamlFilename)
	if err != nil {
		return nil, err
	}

	for _, overrideItem := range overrides {
		overrideKey, overrideValue := toParameter(overrideItem)
		data[overrideKey] = overrideValue
	}

	for k, v := range data {
		key := k.(string)
		value := fmt.Sprintf("%v", v)

		result.data[key] = Param{
			name:      key,
			value:     value,
			paramType: TypeSimple,
			resolved:  true,
		}
	}

	return result, nil
}

// The `source` is passed as --override-param argument
// it looks like `key=value` string, we're splitting it below
func toParameter(source string) (key, value string) {
	segments := strings.SplitN(source, EqualSign, 2)
	if len(segments) == 2 {
		return segments[0], segments[1]
	}
	return segments[0], Empty
}

func (params *Parameters) getValue(name string) (value string, err error) {
	p := params.data[name]

	if p.name == Empty {
		return Empty, errors.New("No value found for key: " + name)
	}

	return p.value, nil
}
