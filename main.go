/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"path"

	"github.com/golang/freetype/truetype"
	"github.com/simplycode07/certificate-cli/cmd"
	"github.com/simplycode07/certificate-cli/generator"
)


func main() {
	templatePath, blankCertificatePath, nameListPath := cmd.Execute()
	var baseDir string = path.Dir(templatePath)
	baseDir += "/"

	fmt.Println("Base Directory:", baseDir)
	fmt.Println("Template Path:", templatePath)
	fmt.Println("Blank Certificate:", blankCertificatePath)
	fmt.Println("Names of attendees:", nameListPath)

	// read settings from template file
	var template generator.Template
	dataByte, _ := os.ReadFile(templatePath)
	err := json.Unmarshal(dataByte, &template)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// read list of names from the csv file
	nameBytes, err := os.Open(nameListPath)
	names, err := csv.NewReader(nameBytes).ReadAll()
	// _, err = csv.NewReader(nameBytes).ReadAll()
	defer nameBytes.Close()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	certificateByte, err := os.Open(blankCertificatePath)
	defer certificateByte.Close()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// read and parse font file
	fontBytes, err := os.ReadFile(baseDir + template.Serial.Font)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	parsedFont, err := truetype.Parse(fontBytes)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}


	// generator.GenerateImage(template, certificateByte, parsedFont, "name", baseDir)

	generator.Initialize(template, certificateByte, parsedFont, baseDir)

	fmt.Println("NAMES: ")
	for _, val := range names {
		for _, name := range val{
			generator.GenerateImage(template, name)
		}
	}
	fmt.Println()
}
