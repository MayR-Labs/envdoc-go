package commands

import (
	"fmt"
	"os"

	"github.com/MayR-Labs/envdoc-go/internal/utils"
)

// handleReportOutput prompts the user for what to do with a generated report
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
