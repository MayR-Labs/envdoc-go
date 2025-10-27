package commands

import (
	"fmt"
	"os"

	"github.com/MayR-Labs/envdoc-go/internal/parser"
	"github.com/MayR-Labs/envdoc-go/internal/utils"
	"github.com/spf13/cobra"
)

// NewArrangeCmd returns the arrange command
func NewArrangeCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "arrange [file]",
		Short: "Arrange and group environment variables",
		Long: `Arrange and group environment variable keys in the specified file. 
Grouping means similar prefixes will be clustered together.`,
		Args: cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			var inputFile string
			var err error

			// Get input file
			if len(args) > 0 {
				inputFile = args[0]
			} else {
				inputFile, err = utils.PromptForEnvFile("Select the .env file to arrange:")
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

			// Parse input file
			envVars, err := parser.ParseEnvFile(inputFile)
			if err != nil {
				fmt.Printf("Error parsing file: %v\n", err)
				os.Exit(1)
			}

			// Arrange by prefix
			arrangedVars := parser.ArrangeByPrefix(envVars)

			// Confirm action with PIN
			confirmed, err := utils.ConfirmWithPin(fmt.Sprintf("This will rearrange keys in '%s'.", inputFile))
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				os.Exit(1)
			}
			if !confirmed {
				fmt.Println("Operation cancelled.")
				return
			}

			// Write back to file
			if err := parser.WriteEnvFile(inputFile, arrangedVars); err != nil {
				fmt.Printf("Error writing file: %v\n", err)
				os.Exit(1)
			}

			fmt.Printf("✓ File arranged: %s\n", inputFile)
		},
	}
}
