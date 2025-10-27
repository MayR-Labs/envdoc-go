package commands

import (
	"fmt"
	"os"

	"github.com/MayR-Labs/envdoc-go/internal/parser"
	"github.com/MayR-Labs/envdoc-go/internal/utils"
	"github.com/spf13/cobra"
)

// NewSyncCmd returns the sync command
func NewSyncCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "sync [file1] [file2] [fileN...]",
		Short: "Synchronize keys across multiple files",
		Long: `Synchronizes environment variable keys across multiple specified files. 
Missing keys in each file are added with empty values.`,
		Run: func(cmd *cobra.Command, args []string) {
			var files []string
			var err error

			// Get files
			if len(args) >= 2 {
				files = args
			} else {
				files, err = utils.PromptForMultipleEnvFiles("Select the .env files to sync:")
				if err != nil {
					fmt.Printf("Error: %v\n", err)
					os.Exit(1)
				}
				if len(files) < 2 {
					fmt.Println("Error: At least 2 files are required for synchronization")
					os.Exit(1)
				}
			}

			// Check if all files exist
			for _, file := range files {
				if !utils.FileExists(file) {
					fmt.Printf("Error: File '%s' does not exist\n", file)
					os.Exit(1)
				}
			}

			// Parse all files
			allEnvVars := make(map[string][]parser.EnvVar)
			for _, file := range files {
				envVars, err := parser.ParseEnvFile(file)
				if err != nil {
					fmt.Printf("Error parsing file '%s': %v\n", file, err)
					os.Exit(1)
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

			// Show preview of changes
			fmt.Println("\nSynchronization Preview:")
			fmt.Println("========================")
			for file, envVars := range allEnvVars {
				fileKeys := parser.GetEnvKeys(envVars)
				var allKeysList []string
				for key := range allKeys {
					allKeysList = append(allKeysList, key)
				}
				missing := parser.FindMissingKeys(allKeysList, fileKeys)
				fmt.Printf("\n%s: %d keys to add\n", file, len(missing))
				for _, key := range missing {
					fmt.Printf("  + %s\n", key)
				}
			}
			fmt.Println()

			// Confirm with PIN
			confirmed, err := utils.ConfirmWithPin("This will modify the specified files.")
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				os.Exit(1)
			}
			if !confirmed {
				fmt.Println("Operation cancelled.")
				return
			}

			// Synchronize files
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

				// Sort and write
				envVars = parser.ArrangeByPrefix(envVars)
				if err := parser.WriteEnvFile(file, envVars); err != nil {
					fmt.Printf("Error writing file '%s': %v\n", file, err)
					os.Exit(1)
				}
			}

			fmt.Println("✓ Files synchronized successfully")
		},
	}
}
