package envdoc

import (
	"fmt"
	"os"

	"github.com/MayR-Labs/envdoc-go/internal/commands"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "envdoc",
	Short: "A powerful CLI tool for managing environment files",
	Long: `envdoc is a comprehensive command-line tool for managing, validating, 
and transforming environment variable files. It helps you maintain consistency 
across different environments and provides utilities for encryption, validation, 
and documentation generation.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	// Add all commands
	rootCmd.AddCommand(commands.NewCreateExampleCmd())
	rootCmd.AddCommand(commands.NewCreateSchemaCmd())
	rootCmd.AddCommand(commands.NewArrangeCmd())
	rootCmd.AddCommand(commands.NewAuditCmd())
	rootCmd.AddCommand(commands.NewCompareCmd())
	rootCmd.AddCommand(commands.NewSyncCmd())
	rootCmd.AddCommand(commands.NewBase64Cmd())
	rootCmd.AddCommand(commands.NewHashCmd())
	rootCmd.AddCommand(commands.NewEncryptCmd())
	rootCmd.AddCommand(commands.NewDecryptCmd())
	rootCmd.AddCommand(commands.NewToCmd())
	rootCmd.AddCommand(commands.NewFromCmd())
	rootCmd.AddCommand(commands.NewValidateCmd())
	rootCmd.AddCommand(commands.NewDoctorCmd())
	rootCmd.AddCommand(commands.NewEngineerCmd())
	rootCmd.AddCommand(commands.NewVersionCmd())
	rootCmd.AddCommand(commands.NewDocumentationCmd())
	rootCmd.AddCommand(commands.NewLicenseCmd())
	rootCmd.AddCommand(commands.NewChangelogCmd())
	rootCmd.AddCommand(commands.NewAuthorsCmd())
}
