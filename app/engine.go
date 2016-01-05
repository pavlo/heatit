package app

import (
	"github.com/codegangsta/cli"
	"github.com/pavlo/heatit/utils"
	"log"
	"strings"

	"bytes"
	"fmt"
	"github.com/pavlo/heatit/directives"
	"gopkg.in/yaml.v2"
	"os"
	"regexp"
)

var paramDirectiveRegexp, _ = regexp.Compile(fmt.Sprintf("(%s%s[a-z-]*)",
	directives.ParamDirectiveTag, directives.DirectiveSeparator))

type Engine struct {
	params          *Parameters
	sourceFile      string
	destinationFile string
}

func NewEngine(c *cli.Context) *Engine {
	return engine(
		c.String("source"),
		c.String("destination"),
		c.String("params"),
	)
}

func engine(source string, destination string, params string) *Engine {
	p, err := NewParameters(params)
	if err != nil {
		log.Fatalf("Failed to parse parameters file! %v", err)
	}
	performer := Engine{
		params:          p,
		sourceFile:      source,
		destinationFile: destination,
	}
	return &performer
}

func (engine *Engine) Process() {

	data, err := utils.ReadTextFile(engine.sourceFile)
	if err != nil {
		log.Fatalf("Can not read source YAML! %v\n%s", err, engine.sourceFile)
	}

	data, err = processInserts(data, 0)
	if err != nil {
		log.Fatalf("Can to process inserts! %v\n%s", err, data)
	}

	data, err = processParams(data, engine.params)
	if err != nil {
		log.Fatalf("Can to process params! %v\n%s", err, data)
	}

	var tmp map[interface{}]interface{}
	err = yaml.Unmarshal([]byte(data), &tmp)
	if err != nil {
		utils.DescribeUnmarshalError(data, err)
		os.Exit(1)
	}

	bytes, err := yaml.Marshal(tmp)
	if err != nil {
		log.Fatalf("Failed to marshal the result to YAML! %v", err)
	}

	err = utils.WriteTextFile(engine.destinationFile, bytes)
	if err != nil {
		log.Fatalf("Failed to save the result to %s,", engine.destinationFile)
	}
}

func processInserts(data string, indent int) (string, error) {
	var result bytes.Buffer
	lines := strings.Split(data, NewLine)

	for _, line := range lines {
		cleanLine := strings.TrimSpace(line)

		if len(cleanLine) == 0 {
			continue
		}

		if strings.Index(cleanLine, directives.InsertDirectiveTag) == 0 {
			err := handleSingleInsertion(line, indent, &result)
			if err != nil {
				return Empty, err
			}
		} else {
			result.WriteString(strings.Repeat(Space, indent))
			result.WriteString(line)
			result.WriteString(NewLine)
		}
	}

	return result.String(), nil
}

func processParams(data string, params *Parameters) (string, error) {
	var result bytes.Buffer

	lines := strings.Split(data, NewLine)
	for _, line := range lines {
		matches := paramDirectiveRegexp.FindAllString(line, -1)
		if matches != nil {
			for _, m := range matches {

				directive, err := directives.NewParamDirective(m)
				if err != nil {
					log.Printf("Invalid @param directive format: %s", m)
					return Empty, err
				}

				value, err := params.getValue(directive.Name)
				if err != nil {
					log.Printf("Got a param: '%s', but had no value to replace it with.", m)
					return Empty, err

				}
				line = strings.Replace(line, m, value, -1)
			}
		}
		result.WriteString(line)
		result.WriteString(NewLine)
	}

	return result.String(), nil
}

func handleSingleInsertion(line string, indent int, result *bytes.Buffer) error {
	insertion, err := directives.NewInsertDirective(line)

	if err != nil {
		log.Printf("Failed to create InsertDirective for line %s", line)
		return err
	}

	if insertion.SourceType == directives.InsertDirectiveFileType {
		content, err := utils.ReadTextFile(insertion.SourceValue)

		if err != nil {
			log.Printf("Failed to read %s file for insertion!", insertion.SourceValue)
			return err
		}

		for _, contentLine := range strings.Split(content, NewLine) {
			contentLine, err = processInserts(contentLine, insertion.Indent+indent)
			if err != nil {
				log.Println("Failed to execute handleSingleInsertion")
				return err
			}
			result.WriteString(contentLine)
		}
	}

	return nil
}
