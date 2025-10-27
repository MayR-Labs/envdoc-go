package utils

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/AlecAivazis/survey/v2"
	"github.com/atotto/clipboard"
)

// ConfirmWithPin prompts the user to confirm an action with a 6-digit PIN
func ConfirmWithPin(message string) (bool, error) {
	pin := generatePin()
	fmt.Printf("\n%s\n", message)
	fmt.Printf("Please enter the PIN to confirm: %s\n\n", pin)

	var input string
	prompt := &survey.Input{
		Message: "Enter PIN:",
	}
	if err := survey.AskOne(prompt, &input); err != nil {
		return false, err
	}

	return input == pin, nil
}

// generatePin generates a random 6-digit PIN
func generatePin() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%06d", rand.Intn(1000000))
}

// PromptForFile prompts the user to enter a file path
func PromptForFile(message, defaultValue string) (string, error) {
	var result string
	prompt := &survey.Input{
		Message: message,
		Default: defaultValue,
	}
	if err := survey.AskOne(prompt, &result); err != nil {
		return "", err
	}
	return result, nil
}

// PromptForPassword prompts the user to enter a password
func PromptForPassword(message string) (string, error) {
	var password string
	prompt := &survey.Password{
		Message: message,
	}
	if err := survey.AskOne(prompt, &password); err != nil {
		return "", err
	}
	return password, nil
}

// PromptForSelection prompts the user to select from a list
func PromptForSelection(message string, options []string) (string, error) {
	var selected string
	prompt := &survey.Select{
		Message: message,
		Options: options,
	}
	if err := survey.AskOne(prompt, &selected); err != nil {
		return "", err
	}
	return selected, nil
}

// FileExists checks if a file exists
func FileExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

// CopyToClipboard copies text to the clipboard
func CopyToClipboard(text string) error {
	return clipboard.WriteAll(text)
}

// WriteToFile writes content to a file
func WriteToFile(filename, content string) error {
	return os.WriteFile(filename, []byte(content), 0644)
}

// ReadFromFile reads content from a file
func ReadFromFile(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// GetTimestamp returns a formatted timestamp
func GetTimestamp() string {
	return time.Now().Format("20060102-150405")
}

// PromptForConfirmation prompts the user for a yes/no confirmation
func PromptForConfirmation(message string) (bool, error) {
	var confirmed bool
	prompt := &survey.Confirm{
		Message: message,
	}
	if err := survey.AskOne(prompt, &confirmed); err != nil {
		return false, err
	}
	return confirmed, nil
}

// PromptForMultiSelect prompts the user to select multiple items from a list
func PromptForMultiSelect(message string, options []string) ([]string, error) {
	var selected []string
	prompt := &survey.MultiSelect{
		Message: message,
		Options: options,
	}
	if err := survey.AskOne(prompt, &selected); err != nil {
		return nil, err
	}
	return selected, nil
}

// FindEnvFiles finds all .env files in the current directory
// Excludes backup, json, yaml, hashed, b64, encrypted, and enc files
func FindEnvFiles() ([]string, error) {
	const envPrefix = ".env."
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
		// Include .env and .env.* files
		if name == ".env" || (len(name) > len(envPrefix) && name[:len(envPrefix)] == envPrefix) {
			// Exclude specific suffixes
			if !hasExcludedSuffix(name) {
				files = append(files, name)
			}
		}
	}

	return files, nil
}

// hasExcludedSuffix checks if a filename has an excluded suffix
func hasExcludedSuffix(name string) bool {
	excludedSuffixes := []string{".bak", ".json", ".yaml", ".yml", ".hashed", ".b64", ".encrypted", ".enc"}
	for _, suffix := range excludedSuffixes {
		if strings.HasSuffix(name, suffix) {
			return true
		}
	}
	return false
}

// FindAllFiles finds all files (not directories) in the current directory
func FindAllFiles() ([]string, error) {
	var files []string
	entries, err := os.ReadDir(".")
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			files = append(files, entry.Name())
		}
	}

	return files, nil
}

// PromptForEnvFile prompts the user to select an .env file with a "Custom" option
func PromptForEnvFile(message string) (string, error) {
	files, err := FindEnvFiles()
	if err != nil {
		return "", err
	}

	if len(files) == 0 {
		// No .env files found, ask user to enter custom path
		return PromptForFile(message, "")
	}

	// Add "Custom" option
	options := append(files, "Custom (enter path)")
	selected, err := PromptForSelection(message, options)
	if err != nil {
		return "", err
	}

	if selected == "Custom (enter path)" {
		return PromptForFile("Enter file path:", "")
	}

	return selected, nil
}

// PromptForMultipleEnvFiles prompts the user to select multiple .env files with a "Custom" option
func PromptForMultipleEnvFiles(message string) ([]string, error) {
	files, err := FindEnvFiles()
	if err != nil {
		return nil, err
	}

	if len(files) == 0 {
		// No .env files found, ask user to enter custom paths
		var result []string
		for {
			file, err := PromptForFile("Enter file path (or leave empty to finish):", "")
			if err != nil {
				return nil, err
			}
			if file == "" {
				break
			}
			result = append(result, file)
		}
		return result, nil
	}

	// Add "Custom" option
	options := append(files, "Custom (enter paths)")
	selected, err := PromptForMultiSelect(message, options)
	if err != nil {
		return nil, err
	}

	var result []string
	hasCustom := false
	for _, s := range selected {
		if s == "Custom (enter paths)" {
			hasCustom = true
		} else {
			result = append(result, s)
		}
	}

	// If user selected "Custom", prompt for additional paths
	if hasCustom {
		for {
			file, err := PromptForFile("Enter file path (or leave empty to finish):", "")
			if err != nil {
				return nil, err
			}
			if file == "" {
				break
			}
			result = append(result, file)
		}
	}

	return result, nil
}

// PromptForAnyFile prompts the user to select any file with a "Custom" option
func PromptForAnyFile(message string) (string, error) {
	files, err := FindAllFiles()
	if err != nil {
		return "", err
	}

	if len(files) == 0 {
		// No files found, ask user to enter custom path
		return PromptForFile(message, "")
	}

	// Add "Custom" option
	options := append(files, "Custom (enter path)")
	selected, err := PromptForSelection(message, options)
	if err != nil {
		return "", err
	}

	if selected == "Custom (enter path)" {
		return PromptForFile("Enter file path:", "")
	}

	return selected, nil
}

// PromptForOutputFile prompts for an output file and handles overwrite confirmation
func PromptForOutputFile(message, defaultValue string) (string, error) {
	for {
		filename, err := PromptForFile(message, defaultValue)
		if err != nil {
			return "", err
		}

		// Check if file exists
		if FileExists(filename) {
			options := []string{"Overwrite", "Change filename"}
			choice, err := PromptForSelection(fmt.Sprintf("File '%s' already exists. What would you like to do?", filename), options)
			if err != nil {
				return "", err
			}

			if choice == "Overwrite" {
				return filename, nil
			}
			// If "Change filename", loop continues
		} else {
			return filename, nil
		}
	}
}
