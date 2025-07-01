#!/bin/bash
# Local build script for fastGit
# Created for testing version injection and local development

set -e

# Get version from git tag or use default
VERSION=$(git describe --tags --always --dirty 2>/dev/null || echo "dev-$(date +%Y%m%d)")
MAIN_PACKAGE="./cmd/fastgit"
OUTPUT_NAME="fastGit"

echo "Building fastGit version: $VERSION"

# Build with version injection
go build -ldflags="-s -w -X github.com/KevinYouu/fastGit/internal/version.Version=$VERSION" \
    -o "$OUTPUT_NAME" "$MAIN_PACKAGE"

echo "Build completed: $OUTPUT_NAME"
echo "Run './$OUTPUT_NAME version' to verify version injection"
