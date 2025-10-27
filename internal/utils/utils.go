package utils

import (
	"fmt"
	"math/rand"
	"os"
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
