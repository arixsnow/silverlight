#!/bin/bash

# Exit on error
set -e

echo "Setting up development environment..."

# Check if pre-commit is installed
if ! command -v pre-commit &> /dev/null; then
    echo "Installing pre-commit..."
    pip install pre-commit
fi

# Check if golangci-lint is installed
if ! command -v golangci-lint &> /dev/null; then
    echo "Installing golangci-lint..."
    go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
fi

# Install pre-commit hooks
echo "Installing pre-commit hooks..."
pre-commit install

# Install Go tools
echo "Installing required Go tools..."
go install golang.org/x/tools/cmd/goimports@latest
go install github.com/go-critic/go-critic/cmd/gocritic@latest

echo "Development environment setup complete!"
echo "Pre-commit hooks are now installed and will run before each commit."
