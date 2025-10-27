# envdoc

[![Release](https://img.shields.io/github/v/release/MayR-Labs/envdoc-go)](https://github.com/MayR-Labs/envdoc-go/releases)
[![License](https://img.shields.io/github/license/MayR-Labs/envdoc-go)](LICENSE)
[![Go Version](https://img.shields.io/github/go-mod/go-version/MayR-Labs/envdoc-go)](go.mod)

A powerful CLI tool for managing, validating, and transforming environment variable files.

## 🌟 Features

- **📝 Documentation Generation**: Create example files and JSON schemas from .env files
- **🔍 Auditing**: Find duplicate keys and missing variables across multiple files
- **🔄 Synchronization**: Keep environment files in sync across different environments
- **🔐 Security**: Encrypt/decrypt files with AES-256 and generate SHA256 hashes
- **🔀 Conversion**: Convert between .env, JSON, and YAML formats
- **✅ Validation**: Validate .env files against JSON schemas
- **🎨 Interactive**: User-friendly prompts and PIN-based confirmations
- **📊 Reports**: Generate comprehensive markdown reports with table of contents
- **🎯 Cross-platform**: Works on Linux, macOS, and Windows

## 📦 Installation

### Quick Install (Linux/macOS)

Using curl:
```bash
curl -sSL https://raw.githubusercontent.com/MayR-Labs/envdoc-go/main/install.sh | bash
```

Or using wget:
```bash
wget -qO- https://raw.githubusercontent.com/MayR-Labs/envdoc-go/main/install.sh | bash
```

### Manual Installation

#### Linux

```bash
# AMD64
wget https://github.com/MayR-Labs/envdoc-go/releases/latest/download/envdoc-linux-amd64
chmod +x envdoc-linux-amd64
sudo mv envdoc-linux-amd64 /usr/local/bin/envdoc

# ARM64
wget https://github.com/MayR-Labs/envdoc-go/releases/latest/download/envdoc-linux-arm64
chmod +x envdoc-linux-arm64
sudo mv envdoc-linux-arm64 /usr/local/bin/envdoc
```

#### macOS

```bash
# Intel
wget https://github.com/MayR-Labs/envdoc-go/releases/latest/download/envdoc-darwin-amd64
chmod +x envdoc-darwin-amd64
sudo mv envdoc-darwin-amd64 /usr/local/bin/envdoc

# Apple Silicon (M1/M2/M3)
wget https://github.com/MayR-Labs/envdoc-go/releases/latest/download/envdoc-darwin-arm64
chmod +x envdoc-darwin-arm64
sudo mv envdoc-darwin-arm64 /usr/local/bin/envdoc
```

#### Windows

Download the latest `envdoc-windows-amd64.exe` from the [releases page](https://github.com/MayR-Labs/envdoc-go/releases/latest) and add it to your PATH.

### Build from Source

```bash
git clone https://github.com/MayR-Labs/envdoc-go.git
cd envdoc-go
go build -o envdoc .
sudo mv envdoc /usr/local/bin/
```

### Verify Installation

```bash
envdoc --version
```

## 🚀 Usage

### Quick Start

```bash
# Create an example file from your .env
envdoc create-example .env

# Generate a JSON schema
envdoc create-schema .env

# Audit a single file
envdoc audit .env

# Compare multiple files
envdoc compare .env .env.staging .env.production

# Sync files
envdoc sync .env .env.staging .env.production
```

### Commands

#### 📚 Documentation & Schema Generation

##### Create Example File
```bash
envdoc create-example [file] [output]
```
Generates an example file with empty values based on keys in the source file.

##### Create JSON Schema
```bash
envdoc create-schema [file] [output]
```
Generates a JSON schema defining all environment variables.

-----------------------------------------------------------------------

#### 📑 File Management

##### Arrange
```bash
envdoc arrange [file]
```
Sorts and groups environment variables alphabetically.

##### Sync
```bash
envdoc sync [file1] [file2] [fileN...]
```
Synchronizes keys across multiple files, adding missing keys with empty values.

-----------------------------------------------------------------------

#### 🔍 Auditing & Comparison

##### Audit
```bash
envdoc audit [file]
```
Generates a report of duplicate keys in a file.

##### Compare
```bash
envdoc compare [file1] [file2] [fileN...]
```
Generates a comparison report showing missing keys across files.

##### Doctor
```bash
envdoc doctor
```
Audits all .env files in the current directory.

#### Engineer
```bash
envdoc engineer
```
Synchronizes and arranges all .env files in the current directory.

-----------------------------------------------------------------------

#### 📝 Validation

##### Validate
```bash
envdoc validate [file] [schema-file]
```
Validates a .env file against a JSON schema.

-----------------------------------------------------------------------

#### 🔄 Conversion

##### Convert To
```bash
envdoc to [json|yaml] [file]
```
Converts .env file to JSON or YAML format.

##### Convert From
```bash
envdoc from [file]
```
Converts JSON or YAML file to .env format.

-----------------------------------------------------------------------

#### 🔐 Security

##### Encrypt
```bash
envdoc encrypt [file]
```
Encrypts a file using AES-256-CBC with PBKDF2 key derivation.

##### Decrypt
```bash
envdoc decrypt [file]
```
Decrypts an encrypted file.

##### Hash
```bash
envdoc hash [file]
```
Generates and displays SHA256 hash of a file.

##### Base64
```bash
envdoc base64 [encode|decode] [file]
```
Encodes or decodes a file using base64.

-----------------------------------------------------------------------

#### ℹ️ Information

```bash
envdoc version        # Show version
envdoc documentation  # Open documentation
envdoc license        # Show license
envdoc changelog      # Show changelog
envdoc authors        # Show authors
```

-----------------------------------------------------------------------

## 📖 Examples

### Example 1: Project Setup

```bash
# Create a template for new developers
envdoc create-example .env.production .env.example

# Create a schema for validation
envdoc create-schema .env.production .env.schema.json

# Validate staging environment
envdoc validate .env.staging .env.schema.json
```

### Example 2: Multi-Environment Management

```bash
# Compare environments
envdoc compare .env.development .env.staging .env.production

# Sync missing keys
envdoc sync .env.development .env.staging .env.production

# Arrange all files
envdoc arrange .env.development
envdoc arrange .env.staging
envdoc arrange .env.production
```

### Example 3: Security

```bash
# Encrypt production secrets
envdoc encrypt .env.production

# Generate hash for verification
envdoc hash .env.production.encrypted

# Later, decrypt when needed
envdoc decrypt .env.production.encrypted
```

### Example 4: CI/CD Integration

```bash
# In your CI/CD pipeline
envdoc validate .env $SCHEMA_FILE || exit 1
envdoc compare .env .env.example || exit 1
```

## 🎯 Use Cases

### For Developers
- Quickly create .env.example files for new team members
- Validate local environment against production schema
- Keep track of required environment variables

### For DevOps
- Audit environment configurations across multiple deployments
- Ensure consistency between staging and production
- Generate documentation for environment variables

### For Security Teams
- Encrypt sensitive configuration files
- Verify file integrity with hash generation
- Track changes in environment configurations

## 🤝 Contributing

We welcome contributions! Please follow these steps:

1. **Fork the repository**
   ```bash
   git clone https://github.com/YOUR_USERNAME/envdoc-go.git
   cd envdoc-go
   ```

2. **Create a feature branch**
   ```bash
   git checkout -b feature/amazing-feature
   ```

3. **Make your changes**
   - Write clean, documented code
   - Follow Go best practices
   - Add tests for new features

4. **Test your changes**
   ```bash
   go test ./...
   go build -o envdoc .
   ./envdoc --help
   ```

5. **Commit your changes**
   ```bash
   git add .
   git commit -m "Add amazing feature"
   ```

6. **Push to your fork**
   ```bash
   git push origin feature/amazing-feature
   ```

7. **Open a Pull Request**
   - Go to the original repository
   - Click "New Pull Request"
   - Select your branch
   - Describe your changes

### Code Style

- Follow [Effective Go](https://golang.org/doc/effective_go) guidelines
- Use `gofmt` to format your code
- Write clear commit messages
- Add comments for complex logic

### Development Setup

```bash
# Clone the repository
git clone https://github.com/MayR-Labs/envdoc-go.git
cd envdoc-go

# Install dependencies
go mod download

# Build the project
go build -o envdoc .

# Run tests
go test ./...

# Run linter (if installed)
golangci-lint run
```

## 🐛 Bug Reports & Feature Requests

### Reporting Bugs

If you find a bug, please create an issue with:

- **Clear title**: Briefly describe the problem
- **Description**: Detailed explanation of the issue
- **Steps to reproduce**: How to reproduce the bug
- **Expected behavior**: What should happen
- **Actual behavior**: What actually happens
- **Environment**: OS, Go version, envdoc version
- **Screenshots**: If applicable

Example:
```
Title: envdoc sync fails with special characters in keys

Description: When syncing files with keys containing special characters like '@' or '$',
the command crashes with a parse error.

Steps to reproduce:
1. Create .env with KEY@TEST=value
2. Run: envdoc sync .env .env.test
3. Observe error

Expected: Files should sync successfully
Actual: Parse error: invalid character '@'

Environment:
- OS: Ubuntu 22.04
- Go: 1.21
- envdoc: v0.1.0
```

### Feature Requests

To request a feature:

1. Check if it already exists in [issues](https://github.com/MayR-Labs/envdoc-go/issues)
2. Create a new issue with the label `enhancement`
3. Describe:
   - The problem you're trying to solve
   - Your proposed solution
   - Alternative solutions you've considered
   - Any additional context

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 👥 Authors

Built with ❤️ by [MayR Labs](https://github.com/MayR-Labs)

## 🙏 Acknowledgments

- [Cobra](https://github.com/spf13/cobra) - CLI framework
- [Survey](https://github.com/AlecAivazis/survey) - Interactive prompts
- [godotenv](https://github.com/joho/godotenv) - .env file parsing
- All our [contributors](https://github.com/MayR-Labs/envdoc-go/graphs/contributors)

## 📚 Documentation

For more detailed documentation, visit our [documentation page](https://github.com/MayR-Labs/envdoc-go/wiki).

## 🔗 Links

- [GitHub Repository](https://github.com/MayR-Labs/envdoc-go)
- [Issue Tracker](https://github.com/MayR-Labs/envdoc-go/issues)
- [Releases](https://github.com/MayR-Labs/envdoc-go/releases)
- [Changelog](CHANGELOG.md)

## ⭐ Show Your Support

If you find envdoc helpful, please give it a star on GitHub!

---

**Made with ❤️ by MayR Labs** | [GitHub @MayR-Labs](https://github.com/MayR-Labs) | [Website (mayrlabs.com)](https://mayrlabs.com)
