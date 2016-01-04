package utils

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

const (
	regularLineFormat = "%v:\t\t%v\n"
	errorLineFormat   = "%v: =>\t\t%v\n"
	newLine           = "\n"
)

var currentLineRegexp, _ = regexp.Compile("line (\\d+):")

// This means that it generated a content that can not be parsed into YAML, it is
// developer's issue and our job now is to help the developer out as much as possible.
// The `yaml` library prints out a error message with the line number in the content
// that it thinks had the error. So we dump the content line by line with line numbers
// printed at the left for easier navigation.
// See https://github.com/pavlo/heatit/issues/7
func DescribeUnmarshalError(data string, err error) {
	// Message usually looks like this:
	// "yaml: line 1: mapping values are not allowed in this context"
	msg := err.Error()

	lineNum := -1
	matches := currentLineRegexp.FindStringSubmatch(msg)

	if len(matches) == 2 {
		lineNum, _ = strconv.Atoi(matches[1])
	}

	log.Println("Failed to Unmarshal resulting data to YAML!")
	log.Printf("%v\n", err)

	for ind, str := range strings.Split(data, newLine) {
		format := regularLineFormat
		if lineNum == ind {
			format = errorLineFormat
		}
		fmt.Printf(format, ind, str)
	}
}
