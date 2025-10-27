package commands

import (
	"fmt"
	"os"

	"github.com/MayR-Labs/envdoc-go/internal/parser"
	"github.com/MayR-Labs/envdoc-go/internal/utils"
	"github.com/spf13/cobra"
)

// NewClearValuesCmd returns the clear-values command
func NewClearValuesCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "clear-values [file]",
		Short: "Clear all values from an environment file",
		Long: `Clears all values from the specified environment file, leaving only the keys.
This is a dangerous operation and requires PIN confirmation.`,
		Args: cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			var inputFile string
			var err error

			// Get input file
			if len(args) > 0 {
				inputFile = args[0]
			} else {
				inputFile, err = utils.PromptForEnvFile("Select the .env file to clear values:")
				if err != nil {
					fmt.Printf("Error: %v\n", err)
					os.Exit(1)
				}
			}

			// Check if input file exists
			if !utils.FileExists(inputFile) {
				fmt.Printf("Error: File '%s' does not exist\n", inputFile)
				os.Exit(1)
			}

			// First warning
			fmt.Printf("\n⚠️  WARNING: This operation will CLEAR ALL VALUES from '%s'\n", inputFile)
			fmt.Println("⚠️  This action is IRREVERSIBLE and will remove all sensitive data!")
			fmt.Println("⚠️  Make sure you have a backup before proceeding.\n")

			// First PIN confirmation
			confirmed, err := utils.ConfirmWithPin("To proceed with clearing all values, please confirm with PIN")
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				os.Exit(1)
			}
			if !confirmed {
				fmt.Println("Operation cancelled.")
				return
			}

			// Second warning - final confirmation
			fmt.Printf("\n⚠️  FINAL WARNING: You are about to PERMANENTLY CLEAR all values in '%s'\n", inputFile)
			fmt.Println("⚠️  This is your LAST CHANCE to cancel this operation!\n")

			// Second confirmation (yes/no)
			finalConfirm, err := utils.PromptForConfirmation("Are you ABSOLUTELY SURE you want to continue?")
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				os.Exit(1)
			}
			if !finalConfirm {
				fmt.Println("Operation cancelled.")
				return
			}

			// Parse input file
			envVars, err := parser.ParseEnvFile(inputFile)
			if err != nil {
				fmt.Printf("Error parsing file: %v\n", err)
				os.Exit(1)
			}

			// Clear all values
			clearedVars := make([]parser.EnvVar, len(envVars))
			for i, envVar := range envVars {
				clearedVars[i] = parser.EnvVar{
					Key:        envVar.Key,
					Value:      "",
					Comment:    envVar.Comment,
					BlankAfter: envVar.BlankAfter,
				}
			}

			// Write back to file
			if err := parser.WriteEnvFile(inputFile, clearedVars); err != nil {
				fmt.Printf("Error writing file: %v\n", err)
				os.Exit(1)
			}

			fmt.Printf("✓ All values cleared from: %s\n", inputFile)
			fmt.Printf("✓ %d keys retained with empty values\n", len(clearedVars))
		},
	}
}
