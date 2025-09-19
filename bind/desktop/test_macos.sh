#!/bin/bash

echo "V2Ray macOS Dynamic Library Test Script"
echo "======================================="

# Check if dylib exists
if [ ! -f "v2ray.dylib" ]; then
    echo "Error: v2ray.dylib not found!"
    echo "Please run ./build_macos.sh first to build the dylib."
    exit 1
fi

echo "Dynamic library found: v2ray.dylib"
echo ""

# Show dylib info
echo "Library information:"
file v2ray.dylib
lipo -info v2ray.dylib
echo ""

# Check if example exists
if [ ! -f "example_macos" ]; then
    echo "Building example application..."
    gcc -o example_macos example_macos.c -ldl
    if [ $? -ne 0 ]; then
        echo "Error: Failed to build example_macos"
        echo "Make sure GCC is installed"
        exit 1
    fi
    echo "Example built successfully!"
    echo ""
fi

# Run the example
echo "Running dylib test..."
echo ""
./example_macos

echo ""
echo "Test completed!"
