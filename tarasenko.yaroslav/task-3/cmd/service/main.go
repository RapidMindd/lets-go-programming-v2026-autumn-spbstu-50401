package main

import (
	"flag"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

func main() {
	configPath := flag.String("config", "", "Path to YAML config")
	flag.Parse()

	if *configPath == "" {
		panic("config flag is required")
	}

	type Config struct {
		InputFile  string `yaml:"input-file"`
		OutputFile string `yaml:"output-file"`
	}

	data, err := os.ReadFile(*configPath)
	if err != nil {
		panic("unable to read file")
	}

	var config Config

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		panic("unable to parse config")
	}

	fmt.Println(*configPath)
	fmt.Println(config.InputFile)
	fmt.Println(config.OutputFile)
}
