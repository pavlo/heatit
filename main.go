package main

import (
	"github.com/codegangsta/cli"
	"os"

	"github.com/pavlo/heatit/commands"
)

const (
	VERSION					= "0.0.1"
	DIRECTIVE_INDICATOR 	= "@"
	DIRECTIVE_SEPARATOR   	= ":"
	IMPORT_DIRECTIVE 		= DIRECTIVE_INDICATOR + "import"
	INSERT_DIRECTIVE 		= DIRECTIVE_INDICATOR + "insert"
	PARAM_DIRECTIVE 		= DIRECTIVE_INDICATOR + "param"
	NEW_LINE				= "\n"

	SPACE					= " "
	EMPTY					= ""
)

func main() {
	app := cli.NewApp()

	app.Name = "heatit"
	app.HelpName = app.Name
	app.Version = VERSION

	app.Usage = "A command line tool that simplifies HEAT templates authoring and processing"
	app.Flags = appFlags()
	app.Commands = appCommands()

	app.Run(os.Args)
}

func appCommands() []cli.Command {
	return []cli.Command{

		cli.Command {
			Name: "process",
			Usage: "Processess a YAML template",
			Action: commands.PerformTheProcessCommand,
			Flags: []cli.Flag {
				cli.StringFlag{
					Name: "source, s",
					Value: "heat.yaml",
					Usage: "Source HEAT template to process",
				},
				cli.StringFlag{
					Name: "destination, d",
					Value: "result.yaml",
					Usage: "Destination file where the resulting YAML will be saved",
				},
				cli.StringFlag{
					Name: "params, p",
					Value: "",
					Usage: "A flat YAML file (k/v) to take parameters from",
				},
			},
		},
		/*
		cli.Command {
			Name: "new",
			Usage: "Creates a new Heatit project in the current directory",
			Action: nil,
		},
		*/
		/*
		cli.Command {
			Name: "generate",
			Usage: "Generates an asset",
			Action: nil,
		},
		*/
	}
}

func appFlags() []cli.Flag {
	return []cli.Flag{
		cli.BoolFlag{
			Name: "verbose",
			Usage: "show more output",
		},
	}
}

//func _main() {
//
//	var sourceFile 	= flag.String("source", "heat.yaml", "Path to the source YAML file")
//	var paramsFile 	= flag.String("params-file", EMPTY, "A flat (key/value) YAML file with parameters to substitute @param:XXX directives with")
//	var outFile 	= flag.String("out", "preheat-result.yaml", "Filename to save the resulting YAML to")
//	flag.Parse()
//
//	data := parseYaml(*sourceFile)
//	processDirectives(data)
//
//	bytes, err := yaml.Marshal(&data)
//	if err != nil {
//		panic(err)
//	}
//
//	params := make(map[interface{}]interface{})
//	if *paramsFile != EMPTY {
//		params = parseYaml(*paramsFile)
//	}
//	bytes = processParams(string(bytes), params)
//
//	err = writeTextFile(*outFile, bytes)
//	if err != nil {
//		panic(err)
//	}
//
//	fmt.Print(string(bytes))
//}


/*
func processDirectives(data map[interface{}]interface{}) error {
	for k, v := range data {
		switch v.(type) {
		case string:
			if strings.Index(v.(string), IMPORT_DIRECTIVE) == 0 {
				filename := extractArgumentFromDirective(v.(string))
				yaml := parseYaml(filename)
				data[k] = processDirectives(yaml)
			} else if strings.Index(v.(string), INSERT_DIRECTIVE) == 0 {
				filename := extractArgumentFromDirective(v.(string))
				content, err := readTextFile(filename)
				if err != nil {
					panic(err)
				}
				data[k] = processInserts(content)
			}
		case map[interface{}]interface{}:
			vc := v.(map[interface{}]interface{})
			processDirectives(vc)
		}
	}
	return nil
}
*/


/*
func processParams(data string, params map[interface{}]interface{}) []byte {
	var result bytes.Buffer

	reg, _ := regexp.Compile( fmt.Sprintf("(%s%s[a-z-]*)", PARAM_DIRECTIVE, DIRECTIVE_SEPARATOR) )
	lines := strings.Split(data, NEW_LINE)

	for _, line := range lines {
		matches := reg.FindAllString(line, -1)
		if matches != nil {
			for _, m := range matches {
				paramName := extractArgumentFromDirective(m)
				value := params[paramName]
				if value == nil {
					log.Fatalf("Got a param: '%s', but had no value to replace it with.", m)
				}
				line = strings.Replace(line, m, value.(string), -1)
			}
		}
		result.WriteString(line)
		result.WriteString(NEW_LINE)
	}

	return result.Bytes()
}
*/


/*
func processInserts(content string) string {
	var result bytes.Buffer

	lines := strings.Split(content, NEW_LINE)
	for _, line := range lines {

		cleanLine := strings.TrimSpace(line)
		if strings.Index(cleanLine, INSERT_DIRECTIVE) == 0 {
			fn := extractArgumentFromDirective(cleanLine)

			indentCount := strings.Index(line, DIRECTIVE_INDICATOR)
			indent := strings.Repeat(SPACE, indentCount)

			data, err := readTextFile(fn)
			if err != nil {
				panic(err)
			}

			dataLines := strings.Split(data, NEW_LINE)
			var buffer bytes.Buffer
			for _, dataLine := range dataLines {
				buffer.WriteString(indent)
				buffer.WriteString(dataLine)
				buffer.WriteString(NEW_LINE)
			}
			result.WriteString(buffer.String())
		} else {
			result.WriteString(line)
			result.WriteString(NEW_LINE)
		}
	}

	return result.String()
}
*/


/*func extractArgumentFromDirective(directive string) string {
	segments := strings.Split(directive, DIRECTIVE_SEPARATOR)
	return strings.TrimSpace(segments[1])
}*/

/*func readTextFile(filename string) (string, error) {
	b, e := ioutil.ReadFile(filename)
	return string(b), e
}*/

/*func writeTextFile(filename string, content []byte) error {
	err := ioutil.WriteFile(filename, content, 0644)
	if err != nil {
		return err
	}
	return nil
}*/
