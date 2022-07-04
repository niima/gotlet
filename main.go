package main

import (
	"flag"
	"fmt"
	"gotlet/pkg/colors"
	"io/ioutil"
	"os"
)

func main() {
	flag.Usage = func() {
		Output := flag.CommandLine.Output()
		fmt.Fprintf(Output, "Usage of Simple commandline templating utility %s: \n", os.Args[0])
		flag.PrintDefaults()
	}

	templateFile := flag.String("template", "", "[Required] Path to the template file.")
	envPrefix := flag.String("envprefix", "", "Prefix of the OS environment variables you want to use during templating, if none are provided, all of the environment variables will be imported")
	variablesFile := flag.String("varsfile", "", "Path to the file containing the variables in YAML format, if none are provided, template rendering will fall back to environment variables")
	outputFile := flag.String("output", "", "Path to the the file in which the rendered template will be stored in, defaults to stdout if no file name is provided.")
	stdOut := flag.Bool("stdout", false, "Output the results in stdout.")

	flag.Parse()

	// If no argument is provided to the CLI interface, show help and exit with code 0.
	if len(os.Args[1:]) <= 0 {
		flag.Usage()
		os.Exit(0)
	}

	// Exit with error code 2 (invalid usage of shell command)
	if *templateFile == "" {
		fmt.Println(colors.Red, "\n Error: A template file path is required, use the --template flag", colors.Reset)
		os.Exit(2)
	}

	if *envPrefix == "" {
		fmt.Println(colors.Yellow, "\nWarning: No environment variable prefix has been specified, loading all environment variables.\n"+
			"\t This is not recommended, use the --envprefix flag to specify a prefix, to read more refer to the documentation.", colors.Reset)
	}

	data := &Model{
		Variables: make(map[string]any),
	}

	err := getVariables(data, *variablesFile)
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

	if *outputFile != "" {
		err = ioutil.WriteFile(*outputFile, result, 0644)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	// If no file name is provided or stdout flag is set, print the results in stdout.
	if *outputFile == "" || *stdOut {
		fmt.Println(string(result))
	}
}
