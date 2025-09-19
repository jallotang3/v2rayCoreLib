# V2Ray Cross-Platform Library Makefile

# Variables
DLL_NAME = v2ray.dll
DYLIB_NAME = v2ray.dylib
SOURCE = main.go
EXAMPLE_WIN = example.c
EXAMPLE_MAC = example_macos.c
EXAMPLE_WIN_EXE = example.exe
EXAMPLE_MAC_EXE = example_macos

# Detect OS
UNAME_S := $(shell uname -s)
ifeq ($(UNAME_S),Darwin)
    DEFAULT_TARGET = $(DYLIB_NAME)
else
    DEFAULT_TARGET = $(DLL_NAME)
endif

# Default target
all: $(DEFAULT_TARGET)

# Build Windows DLL
$(DLL_NAME): $(SOURCE)
	@echo "Building V2Ray Windows DLL..."
	@set CGO_ENABLED=1 && set GOOS=windows && set GOARCH=amd64 && go build -buildmode=c-shared -o $(DLL_NAME) $(SOURCE)
	@echo "Build complete: $(DLL_NAME)"

# Build macOS dylib
$(DYLIB_NAME): $(SOURCE)
	@echo "Building V2Ray macOS Dynamic Library..."
	@./build_macos.sh

# Build for 32-bit Windows
dll32:
	@echo "Building V2Ray Windows DLL (32-bit)..."
	@set CGO_ENABLED=1 && set GOOS=windows && set GOARCH=386 && go build -buildmode=c-shared -o v2ray32.dll $(SOURCE)
	@echo "Build complete: v2ray32.dll"

# Build Windows DLL (explicit)
windows: $(DLL_NAME)

# Build macOS dylib (explicit)
macos: $(DYLIB_NAME)

# Build example (OS-specific)
ifeq ($(UNAME_S),Darwin)
example: $(DYLIB_NAME)
	@echo "Building macOS example application..."
	gcc -o $(EXAMPLE_MAC_EXE) $(EXAMPLE_MAC) -ldl
	@echo "Example built: $(EXAMPLE_MAC_EXE)"
else
example: $(DLL_NAME)
	@echo "Building Windows example application..."
	gcc -o $(EXAMPLE_WIN_EXE) $(EXAMPLE_WIN)
	@echo "Example built: $(EXAMPLE_WIN_EXE)"
endif

# Build Windows example
example_windows: $(DLL_NAME)
	@echo "Building Windows example application..."
	gcc -o $(EXAMPLE_WIN_EXE) $(EXAMPLE_WIN)
	@echo "Example built: $(EXAMPLE_WIN_EXE)"

# Build macOS example
example_macos: $(DYLIB_NAME)
	@echo "Building macOS example application..."
	gcc -o $(EXAMPLE_MAC_EXE) $(EXAMPLE_MAC) -ldl
	@echo "Example built: $(EXAMPLE_MAC_EXE)"

# Test (OS-specific)
ifeq ($(UNAME_S),Darwin)
test: $(DYLIB_NAME)
	@echo "Testing V2Ray macOS dylib..."
	@./test_macos.sh
else
test: $(DLL_NAME)
	@echo "Testing V2Ray Windows DLL..."
	@if exist $(EXAMPLE_WIN_EXE) ($(EXAMPLE_WIN_EXE)) else (echo "Example not built. Run 'make example' first.")
endif

# Test Windows DLL
test_windows: $(DLL_NAME)
	@echo "Testing V2Ray Windows DLL..."
	@./test.bat

# Test macOS dylib
test_macos: $(DYLIB_NAME)
	@echo "Testing V2Ray macOS dylib..."
	@./test_macos.sh

# Clean build artifacts (OS-specific)
ifeq ($(UNAME_S),Darwin)
clean:
	@echo "Cleaning build artifacts..."
	@rm -f $(DYLIB_NAME) v2ray_*.dylib v2ray.h $(EXAMPLE_MAC_EXE)
	@echo "Clean complete"
else
clean:
	@echo "Cleaning build artifacts..."
	@if exist $(DLL_NAME) del $(DLL_NAME)
	@if exist v2ray32.dll del v2ray32.dll
	@if exist v2ray.h del v2ray.h
	@if exist $(EXAMPLE_WIN_EXE) del $(EXAMPLE_WIN_EXE)
	@echo "Clean complete"
endif

# Install dependencies (run go mod tidy)
deps:
	@echo "Installing dependencies..."
	go mod tidy
	@echo "Dependencies updated"

# Show help
help:
	@echo "V2Ray Windows DLL Build System"
	@echo ""
	@echo "Available targets:"
	@echo "  all      - Build the DLL (default)"
	@echo "  dll32    - Build 32-bit DLL"
	@echo "  example  - Build example application"
	@echo "  test     - Test the DLL with example"
	@echo "  clean    - Clean build artifacts"
	@echo "  deps     - Update Go dependencies"
	@echo "  help     - Show this help"
	@echo ""
	@echo "Files:"
	@echo "  $(DLL_NAME)     - Main Windows DLL"
	@echo "  v2ray32.dll - 32-bit Windows DLL"
	@echo "  $(EXAMPLE_EXE)  - Example application"
	@echo ""
	@echo "Usage:"
	@echo "  make          # Build DLL"
	@echo "  make example  # Build DLL and example"
	@echo "  make test     # Build and test"

# Phony targets
.PHONY: all dll32 example test clean deps help
