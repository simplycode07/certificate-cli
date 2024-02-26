package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	// "path/filepath"

	"github.com/spf13/cobra"
)

var (
	templatePath string
	BlankCertificatePath string
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

	return templatePath, BlankCertificatePath
}

func init() {
	rootCmd.Flags().StringVarP(&templatePath, "template", "t", "", "path to template.json")
	rootCmd.Flags().StringVarP(&BlankCertificatePath, "certificate", "c", "", "path to blank certificate")

	err := rootCmd.MarkFlagRequired("template")
	if err != nil {
		fmt.Println(err)
	}

	if err := rootCmd.MarkFlagRequired("certificate"); err != nil {
		fmt.Println(err)
	}
}
