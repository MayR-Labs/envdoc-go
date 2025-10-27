package commands

import (
	"fmt"
	"os"
	"strings"

	"github.com/MayR-Labs/envdoc-go/internal/parser"
	"github.com/MayR-Labs/envdoc-go/internal/utils"
	"github.com/spf13/cobra"
)

// NewAuditCmd returns the audit command
func NewAuditCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "audit [file]",
		Short: "Generate a report of missing and duplicated keys",
		Long: `Generates an extensive markdown report of missing environment keys 
and duplicated keys in the specified file.`,
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			inputFile := args[0]

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

			// Find duplicates
			duplicates := parser.FindDuplicates(envVars)

			// Generate report
			report := generateAuditReport(inputFile, duplicates, len(envVars))

			// Show options
			handleReportOutput(report, "envdoc-audit")
		},
	}
}

// NewCompareCmd returns the compare command
func NewCompareCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "compare [file1] [file2] [fileN...]",
		Short: "Compare keys across multiple files",
		Long: `Generates an extensive markdown report of keys that are missing 
across multiple specified files.`,
		Args: cobra.MinimumNArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			files := args

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

			// Generate comparison report
			report := generateComparisonReport(allEnvVars)

			// Show options
			handleReportOutput(report, "envdoc-compare")
		},
	}
}

func generateAuditReport(filename string, duplicates []string, totalKeys int) string {
	var sb strings.Builder

	sb.WriteString("# Environment Variables Audit Report\n\n")
	sb.WriteString("## Table of Contents\n")
	sb.WriteString("- [Overview](#overview)\n")
	sb.WriteString("- [Duplicate Keys](#duplicate-keys)\n\n")

	sb.WriteString("## Overview\n\n")
	sb.WriteString(fmt.Sprintf("**File:** `%s`\n\n", filename))
	sb.WriteString(fmt.Sprintf("**Total Keys:** %d\n\n", totalKeys))
	sb.WriteString(fmt.Sprintf("**Duplicate Keys:** %d\n\n", len(duplicates)))

	sb.WriteString("## Duplicate Keys\n\n")
	if len(duplicates) == 0 {
		sb.WriteString("✓ No duplicate keys found.\n\n")
	} else {
		sb.WriteString("| Key |\n")
		sb.WriteString("|-----|\n")
		for _, key := range duplicates {
			sb.WriteString(fmt.Sprintf("| `%s` |\n", key))
		}
		sb.WriteString("\n")
	}

	return sb.String()
}

func generateComparisonReport(allEnvVars map[string][]parser.EnvVar) string {
	var sb strings.Builder

	sb.WriteString("# Environment Variables Comparison Report\n\n")
	sb.WriteString("## Table of Contents\n")
	sb.WriteString("- [Overview](#overview)\n")
	sb.WriteString("- [Files Analyzed](#files-analyzed)\n")
	sb.WriteString("- [Missing Keys](#missing-keys)\n\n")

	sb.WriteString("## Overview\n\n")
	sb.WriteString(fmt.Sprintf("**Files Compared:** %d\n\n", len(allEnvVars)))

	sb.WriteString("## Files Analyzed\n\n")
	for file, envVars := range allEnvVars {
		sb.WriteString(fmt.Sprintf("- `%s` (%d keys)\n", file, len(envVars)))
	}
	sb.WriteString("\n")

	// Collect all keys from all files
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

func handleReportOutput(report, prefix string) {
	options := []string{"Show on CLI", "Copy report content", "Save to file"}
	selected, err := utils.PromptForSelection("What would you like to do with the report?", options)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	switch selected {
	case "Show on CLI":
		fmt.Println("\n" + report)
	case "Copy report content":
		if err := utils.CopyToClipboard(report); err != nil {
			fmt.Printf("Error copying to clipboard: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("✓ Report copied to clipboard")
	case "Save to file":
		defaultFilename := fmt.Sprintf("%s-%s.md", prefix, utils.GetTimestamp())
		filename, err := utils.PromptForFile("Enter output filename:", defaultFilename)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
		if err := utils.WriteToFile(filename, report); err != nil {
			fmt.Printf("Error writing file: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("✓ Report saved to: %s\n", filename)
	}
}
