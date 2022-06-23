package main

import (
	"flag"
	"fmt"
	"gotlet/pkg/colors"
	"io/ioutil"
	"os"
)

func main() {
	if len(flag.Args()) == 0 {
		fmt.Println("gotlet renders your variables into a template file.")
	}

	templateFile := flag.String("t", "", "Specify the template file path")
	envPrefix := flag.String("p", "", "A prefix for filtering which env variables to include")
	dataFile := flag.String("d", "", "Specify the data file containing the variables in Yaml")
	outputPath := flag.String("o", "result.yaml", "Output file path")
	printOutput := flag.Bool("v", false, "Print out the result in stdout")

	flag.Parse()

	if *templateFile == "" {
		fmt.Println(colors.Red + "- The template file path is required, use -t flag")
		os.Exit(1)
	}

	if *envPrefix == "" {
		fmt.Println(colors.Yellow, "- No env variable prefix specified, loading all env variables.\n"+
			"\tThis is not recommended, use -p flag to specify", colors.Reset)
	}

	data := &Model{
		Variables: make(map[string]any),
	}

	err := getVariables(data, *dataFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	addEnvVariables(data, *envPrefix)

	result, err := renderTemplate(*templateFile, data)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if *outputPath != "" {
		err = ioutil.WriteFile(*outputPath, result, 0644)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	if *printOutput {
		fmt.Println(string(result))
	}
}
