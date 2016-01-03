package app

import (
	"github.com/codegangsta/cli"
	"github.com/pavlo/heatit/utils"
	"log"
	"strings"

	"bytes"
	"fmt"
	"github.com/pavlo/heatit/directives"
	"regexp"
)

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

func (p *Engine) Process() error {

	data, err := utils.ReadTextFile(p.sourceFile)
	if err != nil {
		return err
	}

	data, err = processInserts(data, 0)
	if err != nil {
		return err
	}

	data, err = processParams(data, p.params)

	log.Printf("%+v", p)
	return nil
}

func processInserts(data string, indent int) (string, error) {
	var result bytes.Buffer
	lines := strings.Split(data, NEW_LINE)

	for _, line := range lines {
		cleanLine := strings.TrimSpace(line)

		if len(cleanLine) == 0 {
			continue
		}

		if strings.Index(cleanLine, directives.INSERT_DIRECTIVE) == 0 {
			err := handleSingleInsertion(line, indent, &result)
			if err != nil {
				return EMPTY, err
			}
		} else {
			result.WriteString(strings.Repeat(SPACE, indent))
			result.WriteString(line)
			result.WriteString(NEW_LINE)
		}
	}

	return result.String(), nil
}

func processParams(data string, params *Parameters) (string, error) {
	var result bytes.Buffer

	reg, _ := regexp.Compile(fmt.Sprintf("(%s%s[a-z-]*)", directives.PARAM_DIRECTIVE, directives.DIRECTIVE_SEPARATOR))
	lines := strings.Split(data, NEW_LINE)

	for _, line := range lines {
		matches := reg.FindAllString(line, -1)
		if matches != nil {
			for _, m := range matches {

				directive, err := directives.NewParamDirective(m)
				if err != nil {
					log.Printf("Invalid @param directive format: ", m)
					return EMPTY, err
				}

				value, err := params.getValue(directive.Name)
				if err != nil {
					log.Printf("Got a param: '%s', but had no value to replace it with.", m)
					return EMPTY, err

				}
				line = strings.Replace(line, m, value, -1)
			}
		}
		result.WriteString(line)
		result.WriteString(NEW_LINE)
	}

	return result.String(), nil
}

func handleSingleInsertion(line string, indent int, result *bytes.Buffer) error {
	insertion, err := directives.NewInsertDirective(line)

	if err != nil {
		log.Printf("Failed to create InsertDirective for line %s", line)
		return err
	}

	if insertion.SourceType == directives.INSERT_DIRECTIVE_TYPE_FILE {
		content, err := utils.ReadTextFile(insertion.SourceValue)

		if err != nil {
			log.Printf("Failed to read %s file for insertion!", insertion.SourceValue)
			return err
		}

		for _, contentLine := range strings.Split(content, NEW_LINE) {
			contentLine, err = processInserts(contentLine, insertion.Indent+indent)
			if err != nil {
				log.Printf("Failed to execute processRecursiveInserts")
				return err
			}
			result.WriteString(contentLine)
		}
	}

	return nil
}
