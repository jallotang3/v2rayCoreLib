#!/bin/bash

echo "Building V2Ray iOS Framework..."

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

# Build universal iOS framework directly
echo "Building universal iOS framework..."
gomobile bind -target=ios -o V2Ray.framework .

if [ $? -eq 0 ]; then
    echo "Universal iOS build successful: V2Ray.framework"
else
    echo "Universal iOS build failed!"
    exit 1
fi

# Create XCFramework from universal framework
echo ""
echo "Creating XCFramework..."
xcodebuild -create-xcframework \
    -framework V2Ray.framework \
    -output V2Ray.xcframework

if [ $? -eq 0 ]; then
    echo "XCFramework created successfully: V2Ray.xcframework"
    
    # Show file info
    echo ""
    echo "Generated files and directories:"
    ls -la | grep -E "(V2Ray|\.framework|\.xcframework)"
    
    # Show XCFramework info
    echo ""
    echo "XCFramework information:"
    xcodebuild -version
    if [ -d "V2Ray.xcframework" ]; then
        echo "XCFramework contents:"
        find V2Ray.xcframework -name "*.framework" -exec echo "  {}" \;
    fi
    
    echo ""
    echo "iOS build completed successfully!"
    echo ""
    echo "Generated files:"
    echo "  - V2Ray.xcframework (Universal XCFramework for iOS device + simulator)"
    echo "  - V2Ray.framework (Universal framework)"
    echo "  - V2Ray-arm64.framework (Device only)"
    echo "  - V2Ray-amd64.framework (Simulator only)"
    echo ""
    echo "Usage:"
    echo "  1. Drag V2Ray.xcframework into your Xcode project"
    echo "  2. Add to 'Frameworks, Libraries, and Embedded Content'"
    echo "  3. Set 'Embed & Sign' for the framework"
    echo "  4. Import in Swift: import V2Ray"
    echo "  5. Import in Objective-C: #import <V2Ray/V2Ray.h>"
else
    echo "XCFramework creation failed!"
    exit 1
fi
