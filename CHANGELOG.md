# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [0.1.0] - 2025-01-XX

### Added
- Initial release of envdoc
- `create-example` command to generate example files from .env files
- `create-schema` command to generate JSON schemas from .env files
- `arrange` command to sort and group environment variables
- `audit` command to find duplicate keys in .env files
- `compare` command to compare keys across multiple .env files
- `sync` command to synchronize keys across multiple .env files
- `base64` command for encoding/decoding files
- `hash` command to generate SHA256 hashes
- `encrypt` command for AES-256 encryption with PBKDF2
- `decrypt` command for decrypting encrypted files
- `to` command to convert .env to JSON/YAML
- `from` command to convert JSON/YAML to .env
- `validate` command to validate .env files against JSON schemas
- `doctor` command to audit all .env files in current directory
- `engineer` command to sync and arrange all .env files
- `version`, `documentation`, `license`, `changelog`, and `authors` commands
- Interactive prompts using survey
- PIN-based confirmation for destructive operations
- Comprehensive markdown reports with table of contents
- Clipboard integration for copying reports and hashes
- Installation script (install.sh) for easy setup
- GitHub Actions workflow for automated releases
- Cross-platform support (Linux, macOS, Windows)
- Multi-architecture support (AMD64, ARM64)

[unreleased]: https://github.com/MayR-Labs/envdoc-go/compare/v0.1.0...HEAD
[0.1.0]: https://github.com/MayR-Labs/envdoc-go/releases/tag/v0.1.0
