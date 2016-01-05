package directives

import (
	"fmt"
	"strings"
)

const InsertDirectiveTag = DirectiveIndicator + "insert"

type InsertDirective struct {
	SourceType  string
	SourceValue string
	Indent      int
}

func NewInsertDirective(source string) (*InsertDirective, error) {
	result := &InsertDirective{}

	s := strings.Replace(source, "\"", "", -1)
	s = strings.Replace(s, "'", "", -1)

	result.Indent = strings.Index(s, DirectiveIndicator)

	s = strings.TrimSpace(s)
	segments := strings.SplitN(s, DirectiveSeparator, 2)
	args := strings.TrimSpace(segments[1])

	argSegments := strings.Split(args, DirectiveSeparator)

	switch len(argSegments) {
	case 1:
		result.SourceType = InsertDirectiveFileType
		result.SourceValue = argSegments[0]
	case 2:
		result.SourceType = strings.TrimSpace(argSegments[0])
		result.SourceValue = strings.TrimSpace(argSegments[1])
	default:
		return nil, fmt.Errorf("failed to get directive from %s", source)
	}

	return result, nil
}
