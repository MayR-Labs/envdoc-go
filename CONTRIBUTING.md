# Contributing to envdoc

Thank you for your interest in contributing to envdoc! We welcome contributions from the community.

## Table of Contents

- [Code of Conduct](#code-of-conduct)
- [Getting Started](#getting-started)
- [Development Setup](#development-setup)
- [Making Changes](#making-changes)
- [Code Style](#code-style)
- [Testing](#testing)
- [Submitting Changes](#submitting-changes)
- [Reporting Bugs](#reporting-bugs)
- [Requesting Features](#requesting-features)

## Code of Conduct

This project adheres to a code of conduct. By participating, you are expected to uphold this code. Please be respectful and considerate to others.

## Getting Started

1. Fork the repository on GitHub
2. Clone your fork locally
3. Set up the development environment
4. Create a new branch for your changes
5. Make your changes
6. Test your changes
7. Submit a pull request

## Development Setup

### Prerequisites

- Go 1.25 or higher
- Git

### Setting Up Your Environment

```bash
# Clone your fork
git clone https://github.com/YOUR_USERNAME/envdoc-go.git
cd envdoc-go

# Add the upstream repository
git remote add upstream https://github.com/MayR-Labs/envdoc-go.git

# Install dependencies
go mod download

# Build the project
go build -o envdoc .

# Run the binary
./envdoc --help
```

## Making Changes

### Creating a Branch

Always create a new branch for your changes:

```bash
git checkout -b feature/your-feature-name
# or
git checkout -b fix/your-bug-fix
```

Branch naming conventions:
- `feature/` - for new features
- `fix/` - for bug fixes
- `docs/` - for documentation changes
- `refactor/` - for code refactoring

### Making Commits

Write clear and descriptive commit messages:

```bash
git commit -m "feat: add new feature X"
git commit -m "fix: resolve issue with Y"
git commit -m "docs: update README with installation instructions"
```

Commit message format:
- `feat:` - new feature
- `fix:` - bug fix
- `docs:` - documentation changes
- `refactor:` - code refactoring
- `test:` - adding or updating tests
- `chore:` - maintenance tasks

## Code Style

### Go Style Guide

Follow the [Effective Go](https://golang.org/doc/effective_go) guidelines and these additional rules:

1. **Format your code**: Run `gofmt` before committing
   ```bash
   gofmt -w .
   ```

2. **Use meaningful names**: Variables, functions, and types should have descriptive names

3. **Keep functions small**: Each function should do one thing well

4. **Add comments**: Add comments for exported functions, types, and packages
   ```go
   // ParseEnvFile parses a .env file and returns a list of environment variables
   func ParseEnvFile(filename string) ([]EnvVar, error) {
       // implementation
   }
   ```

5. **Handle errors**: Always check and handle errors appropriately
   ```go
   if err != nil {
       return fmt.Errorf("failed to parse file: %w", err)
   }
   ```

6. **Use constants**: Define constants for magic values
   ```go
   const (
       defaultTimeout = 30 * time.Second
       maxRetries     = 3
   )
   ```

## Testing

### Running Tests

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run tests for a specific package
go test ./internal/parser

# Run tests in verbose mode
go test -v ./...
```

### Writing Tests

1. Place test files next to the code they test (e.g., `parser.go` → `parser_test.go`)

2. Use table-driven tests when appropriate:
   ```go
   func TestParseEnvFile(t *testing.T) {
       tests := []struct {
           name    string
           input   string
           want    []EnvVar
           wantErr bool
       }{
           {
               name:  "valid file",
               input: "KEY=value",
               want:  []EnvVar{{Key: "KEY", Value: "value"}},
           },
           // more test cases...
       }
       
       for _, tt := range tests {
           t.Run(tt.name, func(t *testing.T) {
               got, err := ParseEnvFile(tt.input)
               if (err != nil) != tt.wantErr {
                   t.Errorf("ParseEnvFile() error = %v, wantErr %v", err, tt.wantErr)
               }
               // assertions...
           })
       }
   }
   ```

## Submitting Changes

### Before Submitting

1. **Format your code**: `gofmt -w .`
2. **Run tests**: `go test ./...`
3. **Build the project**: `go build -o envdoc .`
4. **Update documentation**: If you changed functionality, update README.md
5. **Add changelog entry**: Update CHANGELOG.md with your changes

### Creating a Pull Request

1. Push your branch to your fork:
   ```bash
   git push origin feature/your-feature-name
   ```

2. Go to the original repository on GitHub

3. Click "New Pull Request"

4. Select your branch

5. Fill out the PR template:
   - **Title**: Clear, concise description of changes
   - **Description**: Detailed explanation of what and why
   - **Related Issues**: Link any related issues
   - **Testing**: Describe how you tested your changes
   - **Screenshots**: Add screenshots if applicable

6. Submit the pull request

### PR Review Process

1. A maintainer will review your PR
2. Address any feedback or requested changes
3. Once approved, your PR will be merged

## Reporting Bugs

### Before Reporting

1. Check if the bug has already been reported in [Issues](https://github.com/MayR-Labs/envdoc-go/issues)
2. Make sure you're using the latest version
3. Try to reproduce the bug

### Bug Report Template

When reporting a bug, include:

1. **Title**: Clear, descriptive title
2. **Description**: What happened vs. what you expected
3. **Steps to Reproduce**:
   ```
   1. Run command X
   2. With arguments Y
   3. See error Z
   ```
4. **Environment**:
   - OS: (e.g., Ubuntu 22.04, macOS 13, Windows 11)
   - Go version: (run `go version`)
   - envdoc version: (run `envdoc version`)
5. **Additional Context**: Any other relevant information

## Requesting Features

### Feature Request Template

1. **Title**: Clear description of the feature
2. **Problem**: What problem does this solve?
3. **Proposed Solution**: How should it work?
4. **Alternatives**: Other solutions you've considered
5. **Additional Context**: Any other relevant information

## Questions?

If you have questions, feel free to:
- Open an issue with the `question` label
- Reach out to the maintainers

## License

By contributing to envdoc, you agree that your contributions will be licensed under the MIT License.

---

Thank you for contributing to envdoc! 🎉
