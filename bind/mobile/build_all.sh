#!/bin/bash

echo "Building V2Ray Mobile Libraries for All Platforms..."
echo "=================================================="

# Check if gomobile is installed
if ! command -v gomobile &> /dev/null; then
    echo "gomobile not found. Installing..."
    go install golang.org/x/mobile/cmd/gomobile@latest
    gomobile init
fi

# Set environment variables
export CGO_ENABLED=1

echo "Environment:"
echo "  CGO_ENABLED = $CGO_ENABLED"
echo ""

# Build Android libraries
echo "Building Android libraries..."
echo "=============================="
./build_android.sh

if [ $? -ne 0 ]; then
    echo "Android build failed!"
    exit 1
fi

echo ""
echo ""

# Build iOS frameworks
echo "Building iOS frameworks..."
echo "=========================="
./build_ios.sh

if [ $? -ne 0 ]; then
    echo "iOS build failed!"
    exit 1
fi

echo ""
echo ""
echo "All mobile builds completed successfully!"
echo "========================================"
echo ""
echo "Android files:"
ls -lh *.aar 2>/dev/null || echo "  No Android files found"
echo ""
echo "iOS files:"
ls -ld *.framework *.xcframework 2>/dev/null || echo "  No iOS files found"
echo ""
echo "Summary:"
echo "  Android: v2ray.aar (universal), v2ray-arm64.aar, v2ray-amd64.aar"
echo "  iOS: V2Ray.xcframework (universal), V2Ray.framework, V2Ray-*.framework"
echo ""
echo "Integration guides:"
echo "  Android: Add .aar files to your Android project's libs folder"
echo "  iOS: Add .xcframework to your Xcode project"
