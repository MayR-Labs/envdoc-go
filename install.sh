#!/bin/bash

# envdoc installation script
# This script installs the envdoc CLI tool

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# GitHub repository
REPO="MayR-Labs/envdoc-go"
INSTALL_DIR="/usr/local/bin"

echo "╔═══════════════════════════════════════════════════════════╗"
echo "║                  envdoc Installer                          ║"
echo "╚═══════════════════════════════════════════════════════════╝"
echo ""

# Detect OS and architecture
OS=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)

case $OS in
    linux*)
        OS="linux"
        ;;
    darwin*)
        OS="darwin"
        ;;
    mingw*|msys*|cygwin*)
        OS="windows"
        ;;
    *)
        echo -e "${RED}✗ Unsupported operating system: $OS${NC}"
        exit 1
        ;;
esac

case $ARCH in
    x86_64|amd64)
        ARCH="amd64"
        ;;
    arm64|aarch64)
        ARCH="arm64"
        ;;
    *)
        echo -e "${RED}✗ Unsupported architecture: $ARCH${NC}"
        exit 1
        ;;
esac

echo -e "${GREEN}Detected OS:${NC} $OS"
echo -e "${GREEN}Detected Architecture:${NC} $ARCH"
echo ""

# Get latest release version
echo -e "${YELLOW}Fetching latest release...${NC}"
LATEST_VERSION=$(curl -s "https://api.github.com/repos/$REPO/releases/latest" | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/')

if [ -z "$LATEST_VERSION" ]; then
    echo -e "${RED}✗ Failed to fetch latest version${NC}"
    exit 1
fi

echo -e "${GREEN}Latest version:${NC} $LATEST_VERSION"
echo ""

# Construct download URL
BINARY_NAME="envdoc-${OS}-${ARCH}"
if [ "$OS" = "windows" ]; then
    BINARY_NAME="${BINARY_NAME}.exe"
fi

DOWNLOAD_URL="https://github.com/$REPO/releases/download/$LATEST_VERSION/$BINARY_NAME"

echo -e "${YELLOW}Downloading envdoc...${NC}"
echo "URL: $DOWNLOAD_URL"

# Create temporary directory
TMP_DIR=$(mktemp -d)
cd "$TMP_DIR"

# Download binary
if ! curl -L -o envdoc "$DOWNLOAD_URL"; then
    echo -e "${RED}✗ Failed to download envdoc${NC}"
    rm -rf "$TMP_DIR"
    exit 1
fi

# Make binary executable
chmod +x envdoc

echo ""
echo -e "${GREEN}✓ Download complete${NC}"
echo ""

# Check if we need sudo
if [ -w "$INSTALL_DIR" ]; then
    SUDO=""
else
    SUDO="sudo"
    echo -e "${YELLOW}Installing to $INSTALL_DIR requires sudo privileges${NC}"
fi

# Install binary
echo -e "${YELLOW}Installing envdoc to $INSTALL_DIR...${NC}"
if ! $SUDO mv envdoc "$INSTALL_DIR/envdoc"; then
    echo -e "${RED}✗ Failed to install envdoc${NC}"
    rm -rf "$TMP_DIR"
    exit 1
fi

# Clean up
cd - > /dev/null
rm -rf "$TMP_DIR"

echo ""
echo -e "${GREEN}✓ envdoc installed successfully!${NC}"
echo ""
echo "Verify installation by running:"
echo "  envdoc --version"
echo ""
echo "Get started with:"
echo "  envdoc --help"
echo ""
echo -e "${GREEN}Happy environment management! 🚀${NC}"
