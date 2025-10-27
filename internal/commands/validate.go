package commands

import (
	"fmt"
	"os"
	"strings"

	"github.com/MayR-Labs/envdoc-go/internal/parser"
	"github.com/MayR-Labs/envdoc-go/internal/utils"
	"github.com/MayR-Labs/envdoc-go/internal/validator"
	"github.com/spf13/cobra"
)

// NewValidateCmd returns the validate command
func NewValidateCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "validate [file] [schema-file]",
		Short: "Validate a file against a JSON schema",
		Long: `Validates the specified file against the provided JSON schema file. 
A report is generated detailing any discrepancies found during validation.`,
		Args: cobra.MaximumNArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			var inputFile, schemaFile string
			var err error

			// Get input file
			if len(args) > 0 {
				inputFile = args[0]
			} else {
				inputFile, err = utils.PromptForEnvFile("Select the .env file to validate:")
				if err != nil {
					fmt.Printf("Error: %v\n", err)
					os.Exit(1)
				}
			}

			// Get schema file
			if len(args) > 1 {
				schemaFile = args[1]
			} else {
				schemaFile, err = utils.PromptForAnyFile("Select the schema file:")
				if err != nil {
					fmt.Printf("Error: %v\n", err)
					os.Exit(1)
				}
			}

			// Check if files exist
			if !utils.FileExists(inputFile) {
				fmt.Printf("Error: File '%s' does not exist\n", inputFile)
				os.Exit(1)
			}
			if !utils.FileExists(schemaFile) {
				fmt.Printf("Error: Schema file '%s' does not exist\n", schemaFile)
				os.Exit(1)
			}

			// Parse input file
			envVars, err := parser.ParseEnvFile(inputFile)
			if err != nil {
				fmt.Printf("Error parsing file: %v\n", err)
				os.Exit(1)
			}

			// Read schema
			schemaJSON, err := utils.ReadFromFile(schemaFile)
			if err != nil {
				fmt.Printf("Error reading schema: %v\n", err)
				os.Exit(1)
			}

			// Validate
			errors, err := validator.ValidateAgainstSchema(envVars, schemaJSON)
			if err != nil {
				fmt.Printf("Error validating: %v\n", err)
				os.Exit(1)
			}

			// Generate report
			report := generateValidationReport(inputFile, schemaFile, errors)

			// Show options
			handleReportOutput(report, "envdoc-validate")
		},
	}
}

// NewDoctorCmd returns the doctor command
func NewDoctorCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "doctor",
		Short: "Audit all .env files in the current directory",
		Long: `Audits and compares every .env file (.env, .env.*) except encrypted files 
in the current working directory. A comprehensive report is generated.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Find all .env files
			files, err := findEnvFiles()
			if err != nil {
				fmt.Printf("Error finding .env files: %v\n", err)
				os.Exit(1)
			}

			if len(files) == 0 {
				fmt.Println("No .env files found in the current directory")
				return
			}

			fmt.Printf("Found %d .env file(s):\n", len(files))
			for _, file := range files {
				fmt.Printf("  - %s\n", file)
			}
			fmt.Println()

			// Parse all files
			allEnvVars := make(map[string][]parser.EnvVar)
			for _, file := range files {
				envVars, err := parser.ParseEnvFile(file)
				if err != nil {
					fmt.Printf("Warning: Could not parse '%s': %v\n", file, err)
					continue
				}
				allEnvVars[file] = envVars
			}

			// Generate comprehensive report
			report := generateDoctorReport(allEnvVars)

			// Show options
			handleReportOutput(report, "envdoc-doctor")
		},
	}
}

// NewEngineerCmd returns the engineer command
func NewEngineerCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "engineer",
		Short: "Sync and arrange all .env files in the current directory",
		Long: `Synchronizes and arranges every .env file (.env, .env.*) except encrypted files 
in the current working directory.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Find all .env files
			files, err := findEnvFiles()
			if err != nil {
				fmt.Printf("Error finding .env files: %v\n", err)
				os.Exit(1)
			}

			if len(files) == 0 {
				fmt.Println("No .env files found in the current directory")
				return
			}

			fmt.Printf("Found %d .env file(s):\n", len(files))
			for _, file := range files {
				fmt.Printf("  - %s\n", file)
			}
			fmt.Println()

			// Parse all files
			allEnvVars := make(map[string][]parser.EnvVar)
			for _, file := range files {
				envVars, err := parser.ParseEnvFile(file)
				if err != nil {
					fmt.Printf("Warning: Could not parse '%s': %v\n", file, err)
					continue
				}
				allEnvVars[file] = envVars
			}

			// Collect all unique keys
			allKeys := make(map[string]bool)
			for _, envVars := range allEnvVars {
				for _, envVar := range envVars {
					allKeys[envVar.Key] = true
				}
			}

			// Show preview
			fmt.Println("Engineering Preview:")
			fmt.Println("===================")
			for file, envVars := range allEnvVars {
				fileKeys := parser.GetEnvKeys(envVars)
				var allKeysList []string
				for key := range allKeys {
					allKeysList = append(allKeysList, key)
				}
				missing := parser.FindMissingKeys(allKeysList, fileKeys)
				fmt.Printf("\n%s:\n", file)
				fmt.Printf("  - Current keys: %d\n", len(envVars))
				fmt.Printf("  - Keys to add: %d\n", len(missing))
				fmt.Printf("  - Will be arranged: Yes\n")
			}
			fmt.Println()

			// Confirm with PIN
			confirmed, err := utils.ConfirmWithPin("This will synchronize and arrange all .env files.")
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				os.Exit(1)
			}
			if !confirmed {
				fmt.Println("Operation cancelled.")
				return
			}

			// Synchronize and arrange
			for file, envVars := range allEnvVars {
				fileKeys := make(map[string]bool)
				for _, envVar := range envVars {
					fileKeys[envVar.Key] = true
				}

				// Add missing keys
				for key := range allKeys {
					if !fileKeys[key] {
						envVars = append(envVars, parser.EnvVar{
							Key:   key,
							Value: "",
						})
					}
				}

				// Sort and arrange
				envVars = parser.ArrangeByPrefix(envVars)

				// Write back
				if err := parser.WriteEnvFile(file, envVars); err != nil {
					fmt.Printf("Error writing '%s': %v\n", file, err)
					continue
				}
				fmt.Printf("✓ Engineered: %s\n", file)
			}

			fmt.Println("\n✓ All files engineered successfully")
		},
	}
}

func findEnvFiles() ([]string, error) {
	var files []string
	entries, err := os.ReadDir(".")
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		name := entry.Name()
		if (name == ".env" || strings.HasPrefix(name, ".env.")) && !strings.HasSuffix(name, ".encrypted") {
			files = append(files, name)
		}
	}

	return files, nil
}

func generateValidationReport(inputFile, schemaFile string, errors []string) string {
	var sb strings.Builder

	sb.WriteString("# Environment Variables Validation Report\n\n")
	sb.WriteString("## Table of Contents\n")
	sb.WriteString("- [Overview](#overview)\n")
	sb.WriteString("- [Validation Errors](#validation-errors)\n\n")

	sb.WriteString("## Overview\n\n")
	sb.WriteString(fmt.Sprintf("**File:** `%s`\n\n", inputFile))
	sb.WriteString(fmt.Sprintf("**Schema:** `%s`\n\n", schemaFile))
	sb.WriteString(fmt.Sprintf("**Errors Found:** %d\n\n", len(errors)))

	sb.WriteString("## Validation Errors\n\n")
	if len(errors) == 0 {
		sb.WriteString("✓ Validation passed! No errors found.\n\n")
	} else {
		sb.WriteString("| Error |\n")
		sb.WriteString("|-------|\n")
		for _, err := range errors {
			sb.WriteString(fmt.Sprintf("| %s |\n", err))
		}
		sb.WriteString("\n")
	}

	return sb.String()
}

func generateDoctorReport(allEnvVars map[string][]parser.EnvVar) string {
	var sb strings.Builder

	sb.WriteString("# Environment Variables Doctor Report\n\n")
	sb.WriteString("## Table of Contents\n")
	sb.WriteString("- [Overview](#overview)\n")
	sb.WriteString("- [Files Analyzed](#files-analyzed)\n")
	sb.WriteString("- [Duplicates](#duplicates)\n")
	sb.WriteString("- [Missing Keys](#missing-keys)\n\n")

	sb.WriteString("## Overview\n\n")
	sb.WriteString(fmt.Sprintf("**Files Analyzed:** %d\n\n", len(allEnvVars)))

	sb.WriteString("## Files Analyzed\n\n")
	for file, envVars := range allEnvVars {
		duplicates := parser.FindDuplicates(envVars)
		sb.WriteString(fmt.Sprintf("### `%s`\n\n", file))
		sb.WriteString(fmt.Sprintf("- **Total Keys:** %d\n", len(envVars)))
		sb.WriteString(fmt.Sprintf("- **Duplicate Keys:** %d\n\n", len(duplicates)))
	}

	sb.WriteString("## Duplicates\n\n")
	hasDuplicates := false
	for file, envVars := range allEnvVars {
		duplicates := parser.FindDuplicates(envVars)
		if len(duplicates) > 0 {
			hasDuplicates = true
			sb.WriteString(fmt.Sprintf("### `%s`\n\n", file))
			sb.WriteString("| Key |\n")
			sb.WriteString("|-----|\n")
			for _, key := range duplicates {
				sb.WriteString(fmt.Sprintf("| `%s` |\n", key))
			}
			sb.WriteString("\n")
		}
	}
	if !hasDuplicates {
		sb.WriteString("✓ No duplicate keys found in any file.\n\n")
	}

	// Collect all keys
	allKeys := make(map[string]bool)
	for _, envVars := range allEnvVars {
		for _, envVar := range envVars {
			allKeys[envVar.Key] = true
		}
	}

	sb.WriteString("## Missing Keys\n\n")
	for file, envVars := range allEnvVars {
		fileKeys := parser.GetEnvKeys(envVars)
		var allKeysList []string
		for key := range allKeys {
			allKeysList = append(allKeysList, key)
		}
		missing := parser.FindMissingKeys(allKeysList, fileKeys)

		sb.WriteString(fmt.Sprintf("### Missing in `%s`\n\n", file))
		if len(missing) == 0 {
			sb.WriteString("✓ No missing keys.\n\n")
		} else {
			sb.WriteString("| Key |\n")
			sb.WriteString("|-----|\n")
			for _, key := range missing {
				sb.WriteString(fmt.Sprintf("| `%s` |\n", key))
			}
			sb.WriteString("\n")
		}
	}

	return sb.String()
}
