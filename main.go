/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"fmt"
	"github.com/simplycode07/certificate-cli/cmd"
	"github.com/simplycode07/certificate-cli/image"
)

type template struct {
	// make every var json type
	event struct {
		name string
		date string
	}

	title struct {
		font string
		fontsize int

		align_x string
		align_y string
		offset_x int
		offset_y int
	}

	serial struct {
		font string
		fontsize int

		align_x string
		align_y string

		offset_x int
		offset_y int
	}
}

func main() {
	templatePath, blankCertificatePath := cmd.Execute()
	fmt.Println(templatePath)
	fmt.Println(blankCertificatePath)

	image.GenerateImage(templatePath, blankCertificatePath)
}
