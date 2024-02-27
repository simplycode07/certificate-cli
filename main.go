/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/simplycode07/certificate-cli/cmd"
	"github.com/simplycode07/certificate-cli/image"
)

func main() {
	templatePath, blankCertificatePath := cmd.Execute()
	fmt.Println(templatePath)
	fmt.Println(blankCertificatePath)

	var template interface{}
	dataByte, _ := os.ReadFile(templatePath)
	json.Unmarshal(dataByte, &template)

	fmt.Println(template)

	image.GenerateImage(template, blankCertificatePath)
}
