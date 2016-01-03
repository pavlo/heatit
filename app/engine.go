package app

import (
	"github.com/codegangsta/cli"
	"github.com/pavlo/heatit/utils"
	"log"
	"strings"

	"bytes"
	"github.com/pavlo/heatit/directives"
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
	log.Printf("%+v", p)
	return nil
}

func (p *Engine) processInserts() (string, error) {
	data, err := utils.ReadTextFile(p.sourceFile)
	if err != nil {
		return EMPTY, nil
	}

	data, err = processRecursiveInserts(data, 0)
	if err != nil {
		return EMPTY, nil
	}

	return data, nil
}

func processRecursiveInserts(data string, indent int) (string, error) {

	var result bytes.Buffer
	lines := strings.Split(data, NEW_LINE)

	for _, line := range lines {
		cleanLine := strings.TrimSpace(line)
		if len(cleanLine) == 0 {
			continue
		}

		if strings.Index(cleanLine, directives.INSERT_DIRECTIVE) == 0 {

			insertion, err := directives.NewInsertDirective(line)

			if err != nil {
				log.Printf("Failed to create InsertDirective for line %s", line)
				return EMPTY, err
			}

			if insertion.SourceType == directives.INSERT_DIRECTIVE_TYPE_FILE {
				content, err := utils.ReadTextFile(insertion.SourceValue)

				if err != nil {
					log.Printf("Failed to read %s file for insertion!", insertion.SourceValue)
					return EMPTY, err
				}

				for _, contentLine := range strings.Split(content, NEW_LINE) {
					contentLine, err = processRecursiveInserts(contentLine, insertion.Indent+indent)
					if err != nil {
						log.Printf("Failed to execute processRecursiveInserts")
						return EMPTY, err
					}
					result.WriteString(contentLine)
				}
			}
		} else {
			result.WriteString(strings.Repeat(SPACE, indent))
			result.WriteString(line)
			result.WriteString(NEW_LINE)
		}
	}

	return result.String(), nil
}
