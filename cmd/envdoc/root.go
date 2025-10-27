package envdoc

import (
	"fmt"
	"os"

	"github.com/MayR-Labs/envdoc-go/internal/commands"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "envdoc",
	Short: "A powerful CLI tool for managing, validating, and transforming environment variable files",
	Long: `envdoc is a comprehensive tool for working with environment variable files.
It provides features for documentation generation, auditing, synchronization,
security, conversion, validation, and more.`,
}

func init() {
	// Documentation commands
	rootCmd.AddCommand(commands.NewCreateExampleCmd())
	rootCmd.AddCommand(commands.NewCreateSchemaCmd())

	// Auditing commands
	rootCmd.AddCommand(commands.NewAuditCmd())
	rootCmd.AddCommand(commands.NewCompareCmd())

	// Synchronization commands
	rootCmd.AddCommand(commands.NewSyncCmd())

	// Security commands
	rootCmd.AddCommand(commands.NewEncryptCmd())
	rootCmd.AddCommand(commands.NewDecryptCmd())
	rootCmd.AddCommand(commands.NewHashCmd())
	rootCmd.AddCommand(commands.NewBase64Cmd())

	// Conversion commands
	rootCmd.AddCommand(commands.NewToCmd())
	rootCmd.AddCommand(commands.NewFromCmd())

	// Validation commands
	rootCmd.AddCommand(commands.NewValidateCmd())
	rootCmd.AddCommand(commands.NewDoctorCmd())
	rootCmd.AddCommand(commands.NewEngineerCmd())

	// Utility commands
	rootCmd.AddCommand(commands.NewArrangeCmd())
	rootCmd.AddCommand(commands.NewClearValuesCmd())

	// Info commands
	rootCmd.AddCommand(commands.NewVersionCmd())
	rootCmd.AddCommand(commands.NewDocumentationCmd())
	rootCmd.AddCommand(commands.NewLicenseCmd())
	rootCmd.AddCommand(commands.NewChangelogCmd())
	rootCmd.AddCommand(commands.NewAuthorsCmd())
}

// Execute runs the root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
