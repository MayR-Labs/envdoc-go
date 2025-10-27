package commands

import (
	"fmt"
	"os"

	"github.com/MayR-Labs/envdoc-go/internal/parser"
	"github.com/MayR-Labs/envdoc-go/internal/utils"
	"github.com/MayR-Labs/envdoc-go/internal/validator"
	"github.com/spf13/cobra"
)

// NewCreateExampleCmd returns the create-example command
func NewCreateExampleCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "create-example [file] [output]",
		Short: "Generate an example file from environment variables",
		Long: `Generates an example file based on the environment variable keys found in the specified file. 
The values in the example file are set to empty strings.`,
		Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			inputFile := args[0]
			outputFile := ".env.example"
			if len(args) > 1 {
				outputFile = args[1]
			}

			// Check if input file exists
			if !utils.FileExists(inputFile) {
				fmt.Printf("Error: File '%s' does not exist\n", inputFile)
				os.Exit(1)
			}

			// Check if output file exists
			if utils.FileExists(outputFile) {
				confirmed, err := utils.PromptForConfirmation(fmt.Sprintf("File '%s' already exists. Overwrite?", outputFile))
				if err != nil || !confirmed {
					fmt.Println("Operation cancelled.")
					return
				}
			}

			// Parse input file
			envVars, err := parser.ParseEnvFile(inputFile)
			if err != nil {
				fmt.Printf("Error parsing file: %v\n", err)
				os.Exit(1)
			}

			// Create example with empty values
			exampleVars := make([]parser.EnvVar, len(envVars))
			for i, envVar := range envVars {
				exampleVars[i] = parser.EnvVar{
					Key:     envVar.Key,
					Value:   "",
					Comment: envVar.Comment,
				}
			}

			// Write output file
			if err := parser.WriteEnvFile(outputFile, exampleVars); err != nil {
				fmt.Printf("Error writing file: %v\n", err)
				os.Exit(1)
			}

			fmt.Printf("✓ Example file created: %s\n", outputFile)
		},
	}
}

// NewCreateSchemaCmd returns the create-schema command
func NewCreateSchemaCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "create-schema [file] [output]",
		Short: "Generate a JSON schema from environment variables",
		Long: `Generates a JSON schema file based on the environment variable keys found in the specified file. 
The schema defines each key as a string type.`,
		Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			inputFile := args[0]
			outputFile := ".env.schema.json"
			if len(args) > 1 {
				outputFile = args[1]
			}

			// Check if input file exists
			if !utils.FileExists(inputFile) {
				fmt.Printf("Error: File '%s' does not exist\n", inputFile)
				os.Exit(1)
			}

			// Check if output file exists
			if utils.FileExists(outputFile) {
				confirmed, err := utils.PromptForConfirmation(fmt.Sprintf("File '%s' already exists. Overwrite?", outputFile))
				if err != nil || !confirmed {
					fmt.Println("Operation cancelled.")
					return
				}
			}

			// Parse input file
			envVars, err := parser.ParseEnvFile(inputFile)
			if err != nil {
				fmt.Printf("Error parsing file: %v\n", err)
				os.Exit(1)
			}

			// Generate schema
			schemaJSON, err := validator.GenerateSchema(envVars)
			if err != nil {
				fmt.Printf("Error generating schema: %v\n", err)
				os.Exit(1)
			}

			// Write output file
			if err := utils.WriteToFile(outputFile, schemaJSON); err != nil {
				fmt.Printf("Error writing file: %v\n", err)
				os.Exit(1)
			}

			fmt.Printf("✓ Schema file created: %s\n", outputFile)
		},
	}
}
