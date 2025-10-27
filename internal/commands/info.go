package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Version is set at build time via ldflags
var Version = "dev"

// NewVersionCmd returns the version command
func NewVersionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Display the current version of envdoc",
		Long:  `Display the current version of the envdoc tool.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("envdoc version %s\n", Version)
		},
	}
}

// NewDocumentationCmd returns the documentation command
func NewDocumentationCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "documentation",
		Short: "Open the documentation for envdoc",
		Long:  `Opens the documentation for envdoc in the default web browser.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Opening documentation: https://github.com/MayR-Labs/envdoc-go")
			fmt.Println("Please visit the URL above to view the documentation.")
		},
	}
}

// NewLicenseCmd returns the license command
func NewLicenseCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "license",
		Short: "Display license information",
		Long:  `Displays the license information for the envdoc tool.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("MIT License")
			fmt.Println()
			fmt.Println("Copyright (c) 2025 MayR Labs")
			fmt.Println()
			fmt.Println("Permission is hereby granted, free of charge, to any person obtaining a copy")
			fmt.Println("of this software and associated documentation files (the \"Software\"), to deal")
			fmt.Println("in the Software without restriction, including without limitation the rights")
			fmt.Println("to use, copy, modify, merge, publish, distribute, sublicense, and/or sell")
			fmt.Println("copies of the Software, and to permit persons to whom the Software is")
			fmt.Println("furnished to do so, subject to the following conditions:")
			fmt.Println()
			fmt.Println("The above copyright notice and this permission notice shall be included in all")
			fmt.Println("copies or substantial portions of the Software.")
			fmt.Println()
			fmt.Println("THE SOFTWARE IS PROVIDED \"AS IS\", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR")
			fmt.Println("IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,")
			fmt.Println("FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE")
			fmt.Println("AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER")
			fmt.Println("LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,")
			fmt.Println("OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE")
			fmt.Println("SOFTWARE.")
		},
	}
}

// NewChangelogCmd returns the changelog command
func NewChangelogCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "changelog",
		Short: "Display the changelog",
		Long:  `Displays the changelog for the envdoc tool.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Changelog: https://github.com/MayR-Labs/envdoc-go/blob/main/CHANGELOG.md")
			fmt.Println("Please visit the URL above to view the full changelog.")
		},
	}
}

// NewAuthorsCmd returns the authors command
func NewAuthorsCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "authors",
		Short: "Display the authors",
		Long:  `Displays the authors of the envdoc tool.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Built with ❤️  by MayR Labs")
			fmt.Println("https://github.com/MayR-Labs")
		},
	}
}
