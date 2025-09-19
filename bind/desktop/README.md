# V2Ray Cross-Platform Library

This directory contains cross-platform bindings for V2Ray Core, providing a simple C interface for starting, stopping, and managing V2Ray instances on Windows and macOS.

## Features

The DLL provides the following functions:

- **StartV2Ray(configData)** - Start V2Ray with JSON configuration string
- **StartV2RayWithConfig(configPath)** - Start V2Ray with configuration file path
- **StopV2Ray()** - Stop the running V2Ray instance
- **GetV2RayStatus()** - Get current V2Ray status ("running" or "stopped")
- **IsV2RayRunning()** - Check if V2Ray is running (returns 1 or 0)
- **GetV2RayVersion()** - Get V2Ray version string
- **RestartV2Ray(configData)** - Restart V2Ray with new configuration
- **RestartV2RayWithConfig(configPath)** - Restart V2Ray with new configuration file
- **TestV2RayConfig(configData)** - Test if JSON configuration is valid
- **TestV2RayConfigFile(configPath)** - Test if configuration file is valid
- **FreeString(str)** - Free memory allocated for C strings

## Building the DLL

### Prerequisites

1. Go 1.19 or later
2. CGO enabled (requires GCC on Windows, e.g., from MinGW-w64 or TDM-GCC)
3. Windows SDK (for C compilation)

### Build Steps

1. Open Command Prompt or PowerShell in this directory
2. Run the build script:
   ```cmd
   build.bat
   ```

Or manually build with:
```cmd
set CGO_ENABLED=1
set GOOS=windows
set GOARCH=amd64
go build -buildmode=c-shared -o v2ray.dll main.go
```

This will generate:
- `v2ray.dll` - The Windows DLL
- `v2ray.h` - Auto-generated C header (use our custom v2ray.h for better documentation)

## Using the DLL

### From C/C++

1. Include the header file:
   ```c
   #include "v2ray.h"
   ```

2. Link against the DLL or load it dynamically:
   ```c
   // Dynamic loading example
   HMODULE hDll = LoadLibrary(L"v2ray.dll");
   StartV2RayFunc StartV2Ray = (StartV2RayFunc)GetProcAddress(hDll, "StartV2Ray");
   ```

3. Use the functions:
   ```c
   char* result = StartV2Ray(configJson);
   printf("Result: %s\n", result);
   FreeString(result);  // Important: free the returned string
   ```

### Example Configuration

Here's a minimal V2Ray configuration for a SOCKS proxy:

```json
{
  "inbounds": [{
    "port": 1080,
    "protocol": "socks",
    "settings": {
      "udp": true
    }
  }],
  "outbounds": [{
    "protocol": "freedom"
  }]
}
```

### Memory Management

**Important**: All string functions return dynamically allocated memory. You must call `FreeString()` on every returned string to avoid memory leaks:

```c
char* status = GetV2RayStatus();
printf("Status: %s\n", status);
FreeString(status);  // Don't forget this!
```

## Example Usage

See `example.c` for a complete example showing how to:
1. Load the DLL
2. Get function pointers
3. Start V2Ray with configuration
4. Check status
5. Stop V2Ray
6. Proper cleanup

Compile the example:
```cmd
gcc -o example.exe example.c
```

## Error Handling

All functions return descriptive error messages as strings. Check the returned string to determine if the operation was successful:

```c
char* result = StartV2Ray(config);
if (strstr(result, "successfully") != NULL) {
    printf("Success: %s\n", result);
} else {
    printf("Error: %s\n", result);
}
FreeString(result);
```

## Thread Safety

The DLL is thread-safe. All functions use internal locking to ensure safe concurrent access.

## Supported Platforms

- Windows 10/11 (x64)
- Windows Server 2016/2019/2022 (x64)

For 32-bit support, change `GOARCH=386` in the build script.

## Troubleshooting

### Build Issues

1. **CGO not enabled**: Ensure `CGO_ENABLED=1` is set
2. **GCC not found**: Install MinGW-w64 or TDM-GCC and add to PATH
3. **Go modules**: Run `go mod tidy` if there are dependency issues

### Runtime Issues

1. **DLL not found**: Ensure `v2ray.dll` is in the same directory as your executable or in PATH
2. **Function not found**: Check that you're using the correct function names (case-sensitive)
3. **Configuration errors**: Use `TestV2RayConfig()` to validate your configuration before starting

### Common Configuration Issues

1. **Port conflicts**: Ensure the port specified in inbounds is not already in use
2. **Invalid JSON**: Validate your JSON configuration syntax
3. **Missing protocols**: Ensure all required protocols are properly configured

## Integration Examples

### C# (.NET)

```csharp
[DllImport("v2ray.dll", CallingConvention = CallingConvention.Cdecl)]
public static extern IntPtr StartV2Ray(string config);

[DllImport("v2ray.dll", CallingConvention = CallingConvention.Cdecl)]
public static extern void FreeString(IntPtr str);

// Usage
IntPtr result = StartV2Ray(configJson);
string message = Marshal.PtrToStringAnsi(result);
FreeString(result);
```

### Python (ctypes)

```python
import ctypes

# Load DLL
v2ray = ctypes.CDLL('./v2ray.dll')

# Define function signatures
v2ray.StartV2Ray.argtypes = [ctypes.c_char_p]
v2ray.StartV2Ray.restype = ctypes.c_char_p
v2ray.FreeString.argtypes = [ctypes.c_char_p]

# Usage
config = b'{"inbounds":[{"port":1080,"protocol":"socks"}],"outbounds":[{"protocol":"freedom"}]}'
result = v2ray.StartV2Ray(config)
print(result.decode('utf-8'))
v2ray.FreeString(result)
```

## License

This project follows the same license as V2Ray Core.
