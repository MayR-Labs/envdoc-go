package commands

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/MayR-Labs/envdoc-go/internal/crypto"
	"github.com/MayR-Labs/envdoc-go/internal/utils"
	"github.com/spf13/cobra"
)

// NewBase64Cmd returns the base64 command
func NewBase64Cmd() *cobra.Command {
	return &cobra.Command{
		Use:   "base64 [encode|decode] [file]",
		Short: "Encode or decode a file using base64",
		Long:  `Encodes or decodes the specified file using base64 encoding.`,
		Args:  cobra.MaximumNArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			var operation, inputFile string
			var err error

			// Get operation
			if len(args) > 0 {
				operation = args[0]
			} else {
				operation, err = utils.PromptForSelection("Select operation:", []string{"encode", "decode"})
				if err != nil {
					fmt.Printf("Error: %v\n", err)
					os.Exit(1)
				}
			}

			if operation != "encode" && operation != "decode" {
				fmt.Println("Error: Operation must be 'encode' or 'decode'")
				os.Exit(1)
			}

			// Get input file
			if len(args) > 1 {
				inputFile = args[1]
			} else {
				inputFile, err = utils.PromptForAnyFile("Select the file:")
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

			// Determine output filename
			var defaultOutput string
			if operation == "encode" {
				defaultOutput = inputFile + ".b64"
			} else {
				defaultOutput = strings.TrimSuffix(inputFile, ".b64") + ".decoded"
			}

			outputFile, err := utils.PromptForOutputFile("Enter output filename:", defaultOutput)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				os.Exit(1)
			}

			// Process file
			var output string
			if operation == "encode" {
				output = crypto.EncodeBase64([]byte(data))
			} else {
				decoded, err := crypto.DecodeBase64(data)
				if err != nil {
					fmt.Printf("Error decoding: %v\n", err)
					os.Exit(1)
				}
				output = string(decoded)
			}

			// Write output file
			if err := utils.WriteToFile(outputFile, output); err != nil {
				fmt.Printf("Error writing file: %v\n", err)
				os.Exit(1)
			}

			fmt.Printf("✓ File %sd: %s\n", operation, outputFile)
		},
	}
}

// NewHashCmd returns the hash command
func NewHashCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "hash [file]",
		Short: "Generate SHA256 hash of a file",
		Long:  `Generates a SHA256 hash of the specified file's contents and displays it.`,
		Args:  cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			var inputFile string
			var err error

			// Get input file
			if len(args) > 0 {
				inputFile = args[0]
			} else {
				inputFile, err = utils.PromptForAnyFile("Select the file to hash:")
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

			// Generate hash
			hash := crypto.HashSHA256([]byte(data))

			fmt.Printf("\nSHA256 Hash: %s\n\n", hash)

			// Prompt to copyToClipboard
			copyToClipboard, err := utils.PromptForConfirmation("Copy hash to clipboard?")
			if err == nil && copyToClipboard {
				if err := utils.CopyToClipboard(hash); err != nil {
					fmt.Printf("Error copying to clipboard: %v\n", err)
				} else {
					fmt.Println("✓ Hash copied to clipboard")
				}
			}
		},
	}
}

// NewEncryptCmd returns the encrypt command
func NewEncryptCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "encrypt [file]",
		Short: "Encrypt a file using AES-256",
		Long:  `Encrypts the specified file using AES-256-CBC encryption with PBKDF2 key derivation.`,
		Args:  cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			var inputFile string
			var err error

			// Get input file
			if len(args) > 0 {
				inputFile = args[0]
			} else {
				inputFile, err = utils.PromptForAnyFile("Select the file to encrypt:")
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

			// Get password
			password, err := utils.PromptForPassword("Enter encryption password:")
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				os.Exit(1)
			}

			if password == "" {
				fmt.Println("Error: Password cannot be empty")
				os.Exit(1)
			}

			// Confirm password
			confirmPassword, err := utils.PromptForPassword("Confirm password:")
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				os.Exit(1)
			}

			if password != confirmPassword {
				fmt.Println("Error: Passwords do not match")
				os.Exit(1)
			}

			// Read input file
			data, err := utils.ReadFromFile(inputFile)
			if err != nil {
				fmt.Printf("Error reading file: %v\n", err)
				os.Exit(1)
			}

			// Encrypt
			encrypted, err := crypto.Encrypt([]byte(data), password)
			if err != nil {
				fmt.Printf("Error encrypting: %v\n", err)
				os.Exit(1)
			}

			// Get output filename
			defaultOutput := inputFile + ".encrypted"
			outputFile, err := utils.PromptForOutputFile("Enter output filename:", defaultOutput)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				os.Exit(1)
			}

			// Write output file
			if err := utils.WriteToFile(outputFile, encrypted); err != nil {
				fmt.Printf("Error writing file: %v\n", err)
				os.Exit(1)
			}

			fmt.Printf("✓ File encrypted: %s\n", outputFile)
		},
	}
}

// NewDecryptCmd returns the decrypt command
func NewDecryptCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "decrypt [file]",
		Short: "Decrypt an encrypted file",
		Long:  `Decrypts a file that was encrypted using the encrypt command.`,
		Args:  cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			var inputFile string
			var err error

			// Get input file
			if len(args) > 0 {
				inputFile = args[0]
			} else {
				inputFile, err = utils.PromptForAnyFile("Select the file to decrypt:")
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

			// Get password
			password, err := utils.PromptForPassword("Enter decryption password:")
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				os.Exit(1)
			}

			// Read input file
			data, err := utils.ReadFromFile(inputFile)
			if err != nil {
				fmt.Printf("Error reading file: %v\n", err)
				os.Exit(1)
			}

			// Decrypt
			decrypted, err := crypto.Decrypt(data, password)
			if err != nil {
				fmt.Printf("Error decrypting: %v\n", err)
				os.Exit(1)
			}

			// Get output filename
			defaultOutput := strings.TrimSuffix(inputFile, ".encrypted")
			if defaultOutput == inputFile {
				defaultOutput = filepath.Base(inputFile) + ".decrypted"
			}

			outputFile, err := utils.PromptForOutputFile("Enter output filename:", defaultOutput)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				os.Exit(1)
			}

			// Write output file
			if err := utils.WriteToFile(outputFile, string(decrypted)); err != nil {
				fmt.Printf("Error writing file: %v\n", err)
				os.Exit(1)
			}

			fmt.Printf("✓ File decrypted: %s\n", outputFile)
		},
	}
}
