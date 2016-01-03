package directives

import (
	"errors"
	"fmt"
	"strings"
)

const PARAM_DIRECTIVE = DIRECTIVE_INDICATOR + "param"

type ParamDirective struct {
	Name string
}

func NewParamDirective(source string) (*ParamDirective, error) {
	result := &ParamDirective{}
	name := strings.Split(source, DIRECTIVE_SEPARATOR)

	if len(name) != 2 {
		return nil, errors.New(fmt.Sprintf("Invlid ParameterDirective declaration: %s", source))
	}

	result.Name = strings.TrimSpace(name[1])
	return result, nil
}
