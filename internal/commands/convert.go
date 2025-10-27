package commands

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/MayR-Labs/envdoc-go/internal/parser"
	"github.com/MayR-Labs/envdoc-go/internal/utils"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

// NewToCmd returns the to command
func NewToCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "to [json|yaml] [file]",
		Short: "Convert .env file to JSON or YAML",
		Long:  `Converts the specified .env file to the desired format (JSON or YAML).`,
		Args:  cobra.MaximumNArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			var format, inputFile string
			var err error

			// Get format
			if len(args) > 0 {
				format = args[0]
			} else {
				format, err = utils.PromptForSelection("Select output format:", []string{"json", "yaml"})
				if err != nil {
					fmt.Printf("Error: %v\n", err)
					os.Exit(1)
				}
			}

			if format != "json" && format != "yaml" {
				fmt.Println("Error: Format must be 'json' or 'yaml'")
				os.Exit(1)
			}

			// Get input file
			if len(args) > 1 {
				inputFile = args[1]
			} else {
				inputFile, err = utils.PromptForEnvFile("Select the .env file to convert:")
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

			// Convert to target format
			var output string
			var ext string
			if format == "json" {
				// Convert to map for JSON (order doesn't matter for JSON display)
				envMap := make(map[string]string)
				for _, envVar := range envVars {
					envMap[envVar.Key] = envVar.Value
				}
				jsonData, err := json.MarshalIndent(envMap, "", "  ")
				if err != nil {
					fmt.Printf("Error converting to JSON: %v\n", err)
					os.Exit(1)
				}
				output = string(jsonData)
				ext = ".json"
			} else {
				// For YAML, arrange by prefix and generate manually to preserve order
				arrangedVars := parser.ArrangeByPrefix(envVars)
				output = convertToYAMLWithBlankLines(arrangedVars)
				ext = ".yaml"
			}

			// Get output filename
			baseName := strings.TrimSuffix(filepath.Base(inputFile), filepath.Ext(inputFile))
			defaultOutput := baseName + ext
			outputFile, err := utils.PromptForOutputFile("Enter output filename:", defaultOutput)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				os.Exit(1)
			}

			// Write output file
			if err := utils.WriteToFile(outputFile, output); err != nil {
				fmt.Printf("Error writing file: %v\n", err)
				os.Exit(1)
			}

			fmt.Printf("✓ File converted to %s: %s\n", format, outputFile)
		},
	}
}

// NewFromCmd returns the from command
func NewFromCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "from [file]",
		Short: "Convert JSON or YAML file to .env",
		Long:  `Converts the specified JSON or YAML file to .env format.`,
		Args:  cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			var inputFile string
			var err error

			// Get input file
			if len(args) > 0 {
				inputFile = args[0]
			} else {
				inputFile, err = utils.PromptForAnyFile("Select the file to convert (JSON or YAML):")
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

			// Read input file
			data, err := utils.ReadFromFile(inputFile)
			if err != nil {
				fmt.Printf("Error reading file: %v\n", err)
				os.Exit(1)
			}

			// Determine format and parse
			ext := strings.ToLower(filepath.Ext(inputFile))
			envMap := make(map[string]string)

			switch ext {
			case ".json":
				if err := json.Unmarshal([]byte(data), &envMap); err != nil {
					fmt.Printf("Error parsing JSON: %v\n", err)
					os.Exit(1)
				}
			case ".yaml", ".yml":
				if err := yaml.Unmarshal([]byte(data), &envMap); err != nil {
					fmt.Printf("Error parsing YAML: %v\n", err)
					os.Exit(1)
				}
			default:
				fmt.Printf("Error: Unsupported file format '%s'. Must be .json, .yaml, or .yml\n", ext)
				os.Exit(1)
			}

			// Convert to EnvVar slice
			var envVars []parser.EnvVar
			for key, value := range envMap {
				envVars = append(envVars, parser.EnvVar{
					Key:   key,
					Value: value,
				})
			}

			// Sort
			envVars = parser.ArrangeByPrefix(envVars)

			// Get output filename
			baseName := strings.TrimSuffix(filepath.Base(inputFile), filepath.Ext(inputFile))
			defaultOutput := baseName + ".env"
			outputFile, err := utils.PromptForOutputFile("Enter output filename:", defaultOutput)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				os.Exit(1)
			}

			// Write output file
			if err := parser.WriteEnvFile(outputFile, envVars); err != nil {
				fmt.Printf("Error writing file: %v\n", err)
				os.Exit(1)
			}

			fmt.Printf("✓ File converted to .env: %s\n", outputFile)
		},
	}
}

// convertToYAMLWithBlankLines converts environment variables to YAML format with blank lines between different prefixes
func convertToYAMLWithBlankLines(envVars []parser.EnvVar) string {
	var sb strings.Builder

	for i, envVar := range envVars {
		// Escape special characters in value if needed
		value := envVar.Value
		if strings.ContainsAny(value, ":#{}[],&*!|>'\"%@`") || strings.HasPrefix(value, " ") || strings.HasSuffix(value, " ") {
			// Quote the value if it contains special YAML characters
			value = fmt.Sprintf("\"%s\"", strings.ReplaceAll(value, "\"", "\\\""))
		}

		sb.WriteString(fmt.Sprintf("%s: %s\n", envVar.Key, value))

		// Add blank line if this variable is marked for a blank line after it
		if envVar.BlankAfter && i < len(envVars)-1 {
			sb.WriteString("\n")
		}
	}

	return sb.String()
}
