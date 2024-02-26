package cmd

import (
	"fmt"
	"os"
	// "path/filepath"

	"github.com/spf13/cobra"
)

var (
	templatePath string
	blankCertificatePath string
)


var rootCmd = &cobra.Command{
	Use:   "certcli",
	Short: "Generate Certificates with Name and Serial",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("template: %s\ncertificate: %s\n", templatePath, blankCertificatePath)
		settings, err := os.ReadFile(templatePath)
		if err != nil {
			fmt.Println("could not read file")
		}else{
			fmt.Println(settings)
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&templatePath, "template", "t", "", "path to template.json")
	rootCmd.Flags().StringVarP(&blankCertificatePath, "certificate", "c", "", "path to blank certificate")

	err := rootCmd.MarkFlagRequired("template")
	if err != nil {
		fmt.Println(err)
	}

	if err := rootCmd.MarkFlagRequired("certificate"); err != nil {
		fmt.Println(err)
	}
}


