package directives

import (
	"fmt"
	"strings"
	"net/url"
)

const InsertDirectiveTag = DirectiveIndicator + "insert"

type InsertDirective struct {
	Scheme      string
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

	insertionUrl, err := url.Parse(args)
	if (err != nil) {
		return nil, fmt.Errorf("failed to get directive from %s, %v", source, err)
	}

	if (insertionUrl.Scheme == Empty) {
		result.Scheme = DefaultScheme
	} else {
		result.Scheme = insertionUrl.Scheme
	}

	sourceValue := insertionUrl.String()
	if (result.Scheme == DefaultScheme) {
		sourceValue = strings.Replace(sourceValue, DefaultScheme + "://", "", 1);
	}
	result.SourceValue = sourceValue
	return result, nil
}
