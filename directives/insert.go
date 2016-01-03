package directives

import (
	"errors"
	"fmt"
	"strings"
)

const INSERT_DIRECTIVE = DIRECTIVE_INDICATOR + "insert" + DIRECTIVE_SEPARATOR

type InsertDirective struct {
	SourceType  string
	SourceValue string
	Indent      int
}

func NewInsertDirective(source string) (*InsertDirective, error) {
	result := &InsertDirective{}

	s := strings.Replace(source, "\"", "", -1)
	s = strings.Replace(s, "'", "", -1)

	result.Indent = strings.Index(s, DIRECTIVE_INDICATOR)

	s = strings.TrimSpace(s)
	segments := strings.SplitN(s, DIRECTIVE_SEPARATOR, 2)
	args := strings.TrimSpace(segments[1])

	argSegments := strings.Split(args, DIRECTIVE_SEPARATOR)

	switch len(argSegments) {
	case 1:
		result.SourceType = INSERT_DIRECTIVE_TYPE_FILE
		result.SourceValue = argSegments[0]
	case 2:
		result.SourceType = strings.TrimSpace(argSegments[0])
		result.SourceValue = strings.TrimSpace(argSegments[1])
	default:
		return nil, errors.New(fmt.Sprintf("Failed to get directive from %s!", source))
	}

	return result, nil
}
