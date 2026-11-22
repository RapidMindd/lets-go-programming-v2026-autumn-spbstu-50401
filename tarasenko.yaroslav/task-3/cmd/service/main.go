package main

import (
	"bytes"
	"encoding/xml"
	"flag"
	"fmt"
	"os"

	"golang.org/x/net/html/charset"
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
		panic(err)
	}

	var config Config

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		panic(err)
	}

	type Valute struct {
		NumCode  int    `xml:"NumCode"`
		CharCode string `xml:"CharCode"`
		Value    string `xml:"Value"`
	}

	type ValCurs struct {
		XMLName xml.Name `xml:"ValCurs"`
		Valutes []Valute `xml:"Valute"`
	}

	inputData, err := os.ReadFile(config.InputFile)
	if err != nil {
		panic(err)
	}

	decoder := xml.NewDecoder(bytes.NewReader(inputData))
	decoder.CharsetReader = charset.NewReaderLabel

	var parsedData ValCurs

	err = decoder.Decode(&parsedData)
	if err != nil {
		panic(err)
	}

	fmt.Println(*configPath)
	fmt.Println(config.InputFile)
	fmt.Println(config.OutputFile)
	fmt.Println(len(parsedData.Valutes))
}
