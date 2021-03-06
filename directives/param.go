package directives

import (
	"fmt"
	"strings"
)

const ParamDirectiveTag = DirectiveIndicator + "param"

type ParamDirective struct {
	Name string
}

func NewParamDirective(source string) (*ParamDirective, error) {
	result := &ParamDirective{}
	name := strings.Split(source, DirectiveSeparator)

	if len(name) != 2 {
		return nil, fmt.Errorf("invalid ParameterDirective declaration: %s", source)
	}

	result.Name = strings.TrimSpace(name[1])
	return result, nil
}
