#!/bin/bash

echo "Building V2Ray macOS Dynamic Library..."

# Set environment variables for macOS
export CGO_ENABLED=1
export GOOS=darwin
export GOARCH=amd64

echo "Environment:"
echo "  CGO_ENABLED = $CGO_ENABLED"
echo "  GOOS = $GOOS"
echo "  GOARCH = $GOARCH"
echo ""

# Build for Intel Macs (x86_64)
echo "Building for Intel Macs (x86_64)..."
go build -buildmode=c-shared -o v2ray_amd64.dylib main.go

if [ $? -eq 0 ]; then
    echo "Intel build successful: v2ray_amd64.dylib"
else
    echo "Intel build failed!"
    exit 1
fi

# Build for Apple Silicon Macs (arm64)
echo ""
echo "Building for Apple Silicon Macs (arm64)..."
export GOARCH=arm64
go build -buildmode=c-shared -o v2ray_arm64.dylib main.go

if [ $? -eq 0 ]; then
    echo "ARM64 build successful: v2ray_arm64.dylib"
else
    echo "ARM64 build failed!"
    exit 1
fi

# Create universal binary using lipo
echo ""
echo "Creating universal binary..."
lipo -create -output v2ray.dylib v2ray_amd64.dylib v2ray_arm64.dylib

if [ $? -eq 0 ]; then
    echo "Universal binary created: v2ray.dylib"
    
    # Show file info
    echo ""
    echo "File information:"
    file v2ray.dylib
    lipo -info v2ray.dylib
    
    # Show file sizes
    echo ""
    echo "File sizes:"
    ls -lh v2ray*.dylib
    
    echo ""
    echo "Build completed successfully!"
    echo ""
    echo "Generated files:"
    echo "  - v2ray.dylib (Universal binary for Intel + Apple Silicon)"
    echo "  - v2ray_amd64.dylib (Intel Macs only)"
    echo "  - v2ray_arm64.dylib (Apple Silicon Macs only)"
    echo "  - v2ray.h (C header file)"
    echo ""
    echo "Usage:"
    echo "  Use v2ray.dylib for maximum compatibility"
    echo "  Use architecture-specific files for smaller size"
else
    echo "Failed to create universal binary!"
    exit 1
fi
