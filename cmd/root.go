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
	nameListPath string
)

var rootCmd = &cobra.Command{
	Use:   "certcli",
	Short: "Generate Certificates with Name and Serial",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		// checking if the extension of templatePath is ".json"
		if filepath.Ext(templatePath) != ".json" {
			fmt.Println("only json files are allowed for template and not", filepath.Ext(templatePath))
		}

		// checking extension of list of names
		if filepath.Ext(nameListPath) != ".csv" {
			fmt.Println("only CSV files are allowed for template and not", filepath.Ext(nameListPath))
		}
		
		_, err := os.ReadFile(templatePath)

		if err != nil {
			fmt.Println("could not read template file")
			os.Exit(1)
		}

		_, err = os.ReadFile(nameListPath)
		if err != nil {
			fmt.Println("could not read list of names")
			os.Exit(1)
		}
	},
}

func Execute() (string, string, string) {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}

	return templatePath, blankCertificatePath, nameListPath
}

func init() {
	rootCmd.Flags().StringVarP(&templatePath, "template", "t", "", "path to template file (json file)")
	rootCmd.Flags().StringVarP(&blankCertificatePath, "certificate", "c", "", "path to blank certificate (jepg, png file)")
	rootCmd.Flags().StringVarP(&nameListPath, "names", "n", "", "path to list of names (csv file)")

	if err := rootCmd.MarkFlagRequired("template"); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := rootCmd.MarkFlagRequired("certificate"); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := rootCmd.MarkFlagRequired("names"); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
