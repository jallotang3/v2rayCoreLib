# V2Ray Mobile Libraries

This directory contains mobile bindings for V2Ray Core, providing native libraries for Android and iOS platforms.

## Features

The mobile library provides the following functionality through the `V2RayService` class:

### Core Functions
- **StartWithConfig(configData)** - Start V2Ray with JSON configuration string
- **StartWithConfigFile(configPath)** - Start V2Ray with configuration file path
- **Stop()** - Stop the running V2Ray instance
- **GetStatus()** - Get current V2Ray status ("running" or "stopped")
- **IsRunning()** - Check if V2Ray is running (returns boolean)
- **GetVersion()** - Get V2Ray version string

### Configuration Management
- **Restart(configData)** - Restart V2Ray with new configuration
- **RestartWithConfigFile(configPath)** - Restart V2Ray with new configuration file
- **TestConfig(configData)** - Test if JSON configuration is valid
- **TestConfigFile(configPath)** - Test if configuration file is valid
- **GetConfigTemplate()** - Get a basic configuration template

### Monitoring & Control
- **QueryStats(tag, name)** - Query statistics for a given tag and name
- **GetSystemStats()** - Get system statistics as JSON string
- **SetLogLevel(level)** - Set log level (0=debug, 1=info, 2=warning, 3=error)

## Building the Libraries

### Prerequisites

1. **Go 1.19 or later**
2. **gomobile tool**:
   ```bash
   go install golang.org/x/mobile/cmd/gomobile@latest
   gomobile init
   ```

### For Android

3. **Android NDK** (for Android builds)
4. **Set ANDROID_NDK_HOME environment variable**

### For iOS

3. **Xcode** (for iOS builds)
4. **Xcode Command Line Tools**

### Build Commands

#### Build All Platforms
```bash
chmod +x build_all.sh
./build_all.sh
```

#### Build Android Only
```bash
chmod +x build_android.sh
./build_android.sh
```

#### Build iOS Only
```bash
chmod +x build_ios.sh
./build_ios.sh
```

## Generated Files

### Android
- `v2ray.aar` - Universal library (ARM64 + AMD64)
- `v2ray-arm64.aar` - ARM64 only
- `v2ray-amd64.aar` - AMD64 only

### iOS
- `V2Ray.xcframework` - Universal XCFramework (device + simulator)
- `V2Ray.framework` - Universal framework
- `V2Ray-arm64.framework` - Device only
- `V2Ray-amd64.framework` - Simulator only

## Integration

### Android Integration

1. **Add to your project**:
   ```bash
   # Copy the AAR file to your project
   cp v2ray.aar /path/to/your/android/project/app/libs/
   ```

2. **Update build.gradle**:
   ```gradle
   android {
       ...
   }
   
   dependencies {
       implementation files('libs/v2ray.aar')
       ...
   }
   ```

3. **Use in Java/Kotlin**:
   ```java
   // Java
   import main.V2RayService;
   
   V2RayService service = Main.newV2RayService();
   String result = service.startWithConfig(configJson);
   ```
   
   ```kotlin
   // Kotlin
   import main.V2RayService
   
   val service = Main.newV2RayService()
   val result = service.startWithConfig(configJson)
   ```

### iOS Integration

1. **Add to Xcode project**:
   - Drag `V2Ray.xcframework` into your Xcode project
   - Add to "Frameworks, Libraries, and Embedded Content"
   - Set to "Embed & Sign"

2. **Use in Swift**:
   ```swift
   import V2Ray
   
   let service = V2RayNewV2RayService()
   let result = service?.startWithConfig(configJson)
   ```

3. **Use in Objective-C**:
   ```objc
   #import <V2Ray/V2Ray.h>
   
   V2RayService *service = V2RayNewV2RayService();
   NSString *result = [service startWithConfig:configJson];
   ```

## Example Usage

### Basic Configuration

```json
{
  "log": {
    "loglevel": "info"
  },
  "inbounds": [
    {
      "port": 1080,
      "protocol": "socks",
      "settings": {
        "udp": true
      }
    }
  ],
  "outbounds": [
    {
      "protocol": "freedom"
    }
  ]
}
```

### Android Example

```java
public class V2RayManager {
    private V2RayService service;
    
    public V2RayManager() {
        service = Main.newV2RayService();
    }
    
    public boolean startV2Ray(String config) {
        String result = service.startWithConfig(config);
        return result.contains("successfully");
    }
    
    public boolean stopV2Ray() {
        String result = service.stop();
        return result.contains("successfully");
    }
    
    public boolean isRunning() {
        return service.isRunning();
    }
    
    public String getStatus() {
        return service.getStatus();
    }
}
```

### iOS Example

```swift
class V2RayManager {
    private let service: V2RayService?
    
    init() {
        service = V2RayNewV2RayService()
    }
    
    func startV2Ray(config: String) -> Bool {
        guard let result = service?.startWithConfig(config) else { return false }
        return result.contains("successfully")
    }
    
    func stopV2Ray() -> Bool {
        guard let result = service?.stop() else { return false }
        return result.contains("successfully")
    }
    
    func isRunning() -> Bool {
        return service?.isRunning() ?? false
    }
    
    func getStatus() -> String {
        return service?.getStatus() ?? "unknown"
    }
}
```

## Error Handling

All functions return descriptive error messages as strings. Check the returned string to determine if the operation was successful:

```java
// Android/Java
String result = service.startWithConfig(config);
if (result.contains("successfully")) {
    Log.d("V2Ray", "Started successfully");
} else {
    Log.e("V2Ray", "Start failed: " + result);
}
```

```swift
// iOS/Swift
let result = service?.startWithConfig(config) ?? ""
if result.contains("successfully") {
    print("Started successfully")
} else {
    print("Start failed: \(result)")
}
```

## Thread Safety

The mobile library is thread-safe. All functions use internal locking to ensure safe concurrent access from multiple threads.

## Supported Platforms

### Android
- **Minimum SDK**: API 21 (Android 5.0)
- **Architectures**: ARM64, AMD64
- **Target SDK**: Latest

### iOS
- **Minimum Version**: iOS 12.0
- **Architectures**: ARM64 (device), AMD64 (simulator)
- **Xcode**: 12.0+

## Troubleshooting

### Build Issues

1. **gomobile not found**:
   ```bash
   go install golang.org/x/mobile/cmd/gomobile@latest
   gomobile init
   ```

2. **Android NDK not found**:
   ```bash
   export ANDROID_NDK_HOME=/path/to/android-ndk
   ```

3. **iOS build fails**:
   - Ensure Xcode Command Line Tools are installed
   - Check Xcode version compatibility

### Runtime Issues

1. **Library not found**: Ensure the library is properly added to your project
2. **Configuration errors**: Use `TestConfig()` to validate configuration
3. **Permission issues**: Ensure your app has network permissions

### Common Configuration Issues

1. **Port conflicts**: Ensure the port specified in inbounds is not already in use
2. **Invalid JSON**: Validate your JSON configuration syntax
3. **Missing protocols**: Ensure all required protocols are properly configured

## Performance Considerations

1. **Memory Usage**: V2Ray instances consume memory; stop unused instances
2. **Battery Life**: Network proxying can impact battery life on mobile devices
3. **Network Usage**: Monitor data usage when using V2Ray on mobile networks

## Security Notes

1. **Permissions**: Request only necessary network permissions
2. **Configuration**: Validate all configuration data before use
3. **Storage**: Store sensitive configuration data securely
4. **Updates**: Keep V2Ray library updated for security patches

## License

This project follows the same license as V2Ray Core.

## Support

For issues and questions:
1. Check the troubleshooting section above
2. Review V2Ray Core documentation
3. Check existing issues in the repository
4. Create a new issue with detailed information
