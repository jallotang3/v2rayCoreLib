#!/bin/bash

echo "Building V2Ray Android Libraries..."

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

# Build for Android ARM64
echo "Building for Android ARM64..."
export ANDROID_API=21
gomobile bind -target=android/arm64 -androidapi=$ANDROID_API -o v2ray-arm64.aar .

if [ $? -eq 0 ]; then
    echo "Android ARM64 build successful: v2ray-arm64.aar"
else
    echo "Android ARM64 build failed!"
    exit 1
fi

# Build for Android AMD64 (x86_64)
echo ""
echo "Building for Android AMD64..."
export ANDROID_API=21
gomobile bind -target=android/amd64 -androidapi=$ANDROID_API -o v2ray-amd64.aar .

if [ $? -eq 0 ]; then
    echo "Android AMD64 build successful: v2ray-amd64.aar"
else
    echo "Android AMD64 build failed!"
    exit 1
fi

# Build universal Android library
echo ""
echo "Building universal Android library..."
export ANDROID_API=21
gomobile bind -target=android -androidapi=$ANDROID_API -o v2ray.aar .

if [ $? -eq 0 ]; then
    echo "Universal Android build successful: v2ray.aar"
    
    # Show file info
    echo ""
    echo "Generated files:"
    ls -lh *.aar
    
    echo ""
    echo "Android build completed successfully!"
    echo ""
    echo "Generated files:"
    echo "  - v2ray.aar (Universal library for ARM64 + AMD64)"
    echo "  - v2ray-arm64.aar (ARM64 only)"
    echo "  - v2ray-amd64.aar (AMD64 only)"
    echo ""
    echo "Usage:"
    echo "  Add v2ray.aar to your Android project's libs folder"
    echo "  Add dependency in build.gradle: implementation files('libs/v2ray.aar')"
else
    echo "Universal Android build failed!"
    exit 1
fi
