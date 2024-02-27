package cmd

import (
	"fmt"
	"os"
	"path/filepath"

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
		// checking if the extension of templatePath is ".json"
		if filepath.Ext(templatePath) != ".json" {
			fmt.Println("template files can only be json and not", filepath.Ext(templatePath))
		}
		
		settings, err := os.ReadFile(templatePath)

		if err != nil {
			fmt.Println("could not read file")
			os.Exit(1)
		}else{
			fmt.Println(settings)
		}
	},
}

func Execute() (string, string) {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}

	return templatePath, blankCertificatePath
}

func init() {
	rootCmd.Flags().StringVarP(&templatePath, "template", "t", "", "path to template.json")
	rootCmd.Flags().StringVarP(&blankCertificatePath, "certificate", "c", "", "path to blank certificate")

	err := rootCmd.MarkFlagRequired("template")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := rootCmd.MarkFlagRequired("certificate"); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
