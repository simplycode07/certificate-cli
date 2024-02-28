/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"fmt"
	"os"
	"encoding/csv"
	"encoding/json"

	"github.com/simplycode07/certificate-cli/cmd"
	"github.com/simplycode07/certificate-cli/image"
)

func main() {
	templatePath, blankCertificatePath, nameListPath := cmd.Execute()
	fmt.Println(templatePath)
	fmt.Println(blankCertificatePath)
	fmt.Println(nameListPath)

	// read settings from template file
	var template interface{}
	dataByte, _ := os.ReadFile(templatePath)
	json.Unmarshal(dataByte, &template)

	// fmt.Println(template)
	nameBytes, _ := os.Open(nameListPath)
	names, err := csv.NewReader(nameBytes).ReadAll()
	nameBytes.Close()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("NAMES: ")
	for _, val := range names {
		for _, name := range val{
			fmt.Printf("%s, ", name)
			image.GenerateImage(template, blankCertificatePath, name)
		}
	}
	fmt.Println()
}
