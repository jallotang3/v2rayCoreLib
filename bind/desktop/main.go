package main

import (
	"C"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"sync"
	"unsafe"

	core "github.com/v2fly/v2ray-core/v5"
	"github.com/v2fly/v2ray-core/v5/infra/conf/serial"
	v4 "github.com/v2fly/v2ray-core/v5/infra/conf/v4"
	_ "github.com/v2fly/v2ray-core/v5/main/distro/all"
)

var (
	instance     *core.Instance
	instanceLock sync.Mutex
	isRunning    bool
)

// StartV2Ray starts V2Ray with the given configuration
//
//export StartV2Ray
func StartV2Ray(configData *C.char) *C.char {
	instanceLock.Lock()
	defer instanceLock.Unlock()

	if isRunning {
		return C.CString("V2Ray is already running")
	}

	configStr := C.GoString(configData)
	if configStr == "" {
		return C.CString("Configuration is empty")
	}

	// Parse JSON configuration using v4 config
	var jsonConfig v4.Config
	if err := json.Unmarshal([]byte(configStr), &jsonConfig); err != nil {
		return C.CString(fmt.Sprintf("Failed to parse configuration: %v", err))
	}

	// Convert to protobuf config
	config, err := jsonConfig.Build()
	if err != nil {
		return C.CString(fmt.Sprintf("Failed to build configuration: %v", err))
	}

	// Create V2Ray instance
	server, err := core.New(config)
	if err != nil {
		return C.CString(fmt.Sprintf("Failed to create V2Ray instance: %v", err))
	}

	// Start the server
	if err := server.Start(); err != nil {
		return C.CString(fmt.Sprintf("Failed to start V2Ray: %v", err))
	}

	instance = server
	isRunning = true
	return C.CString("V2Ray started successfully")
}

// StartV2RayWithConfig starts V2Ray with the given configuration file path
//
//export StartV2RayWithConfig
func StartV2RayWithConfig(configPath *C.char) *C.char {
	instanceLock.Lock()
	defer instanceLock.Unlock()

	if isRunning {
		return C.CString("V2Ray is already running")
	}

	configPathStr := C.GoString(configPath)
	if configPathStr == "" {
		return C.CString("Configuration path is empty")
	}

	// Open and read the configuration file
	file, err := os.Open(configPathStr)
	if err != nil {
		return C.CString(fmt.Sprintf("Failed to open configuration file: %v", err))
	}
	defer file.Close()

	// Load configuration from file
	config, err := serial.LoadJSONConfig(file)
	if err != nil {
		return C.CString(fmt.Sprintf("Failed to load configuration: %v", err))
	}

	// Create V2Ray instance
	server, err := core.New(config)
	if err != nil {
		return C.CString(fmt.Sprintf("Failed to create V2Ray instance: %v", err))
	}

	// Start the server
	if err := server.Start(); err != nil {
		return C.CString(fmt.Sprintf("Failed to start V2Ray: %v", err))
	}

	instance = server
	isRunning = true
	return C.CString("V2Ray started successfully")
}

// StopV2Ray stops the running V2Ray instance
//
//export StopV2Ray
func StopV2Ray() *C.char {
	instanceLock.Lock()
	defer instanceLock.Unlock()

	if !isRunning || instance == nil {
		return C.CString("V2Ray is not running")
	}

	if err := instance.Close(); err != nil {
		return C.CString(fmt.Sprintf("Failed to stop V2Ray: %v", err))
	}

	instance = nil
	isRunning = false
	return C.CString("V2Ray stopped successfully")
}

// GetV2RayStatus returns the current status of V2Ray
//
//export GetV2RayStatus
func GetV2RayStatus() *C.char {
	instanceLock.Lock()
	defer instanceLock.Unlock()

	if isRunning && instance != nil {
		return C.CString("running")
	}
	return C.CString("stopped")
}

// IsV2RayRunning returns 1 if V2Ray is running, 0 otherwise
//
//export IsV2RayRunning
func IsV2RayRunning() C.int {
	instanceLock.Lock()
	defer instanceLock.Unlock()

	if isRunning && instance != nil {
		return 1
	}
	return 0
}

// GetV2RayVersion returns the V2Ray version
//
//export GetV2RayVersion
func GetV2RayVersion() *C.char {
	return C.CString(core.Version())
}

// FreeString frees the memory allocated for C strings
//
//export FreeString
func FreeString(str *C.char) {
	C.free(unsafe.Pointer(str))
}

// RestartV2Ray restarts V2Ray with new configuration
//
//export RestartV2Ray
func RestartV2Ray(configData *C.char) *C.char {
	// Stop current instance
	stopResult := StopV2Ray()
	stopResultStr := C.GoString(stopResult)
	C.free(unsafe.Pointer(stopResult))

	if stopResultStr != "V2Ray stopped successfully" && stopResultStr != "V2Ray is not running" {
		return C.CString(fmt.Sprintf("Failed to stop V2Ray: %s", stopResultStr))
	}

	// Start with new configuration
	return StartV2Ray(configData)
}

// RestartV2RayWithConfig restarts V2Ray with new configuration file
//
//export RestartV2RayWithConfig
func RestartV2RayWithConfig(configPath *C.char) *C.char {
	// Stop current instance
	stopResult := StopV2Ray()
	stopResultStr := C.GoString(stopResult)
	C.free(unsafe.Pointer(stopResult))

	if stopResultStr != "V2Ray stopped successfully" && stopResultStr != "V2Ray is not running" {
		return C.CString(fmt.Sprintf("Failed to stop V2Ray: %s", stopResultStr))
	}

	// Start with new configuration file
	return StartV2RayWithConfig(configPath)
}

// TestV2RayConfig tests if the given configuration is valid
//
//export TestV2RayConfig
func TestV2RayConfig(configData *C.char) *C.char {
	configStr := C.GoString(configData)
	if configStr == "" {
		return C.CString("Configuration is empty")
	}

	// Parse JSON configuration using v4 config
	var jsonConfig v4.Config
	if err := json.Unmarshal([]byte(configStr), &jsonConfig); err != nil {
		return C.CString(fmt.Sprintf("Invalid JSON configuration: %v", err))
	}

	// Convert to protobuf config
	_, err := jsonConfig.Build()
	if err != nil {
		return C.CString(fmt.Sprintf("Invalid configuration: %v", err))
	}

	return C.CString("Configuration is valid")
}

// TestV2RayConfigFile tests if the given configuration file is valid
//
//export TestV2RayConfigFile
func TestV2RayConfigFile(configPath *C.char) *C.char {
	configPathStr := C.GoString(configPath)
	if configPathStr == "" {
		return C.CString("Configuration path is empty")
	}

	// Open and read the configuration file
	file, err := os.Open(configPathStr)
	if err != nil {
		return C.CString(fmt.Sprintf("Failed to open configuration file: %v", err))
	}
	defer file.Close()

	// Load configuration from file
	_, err = serial.LoadJSONConfig(file)
	if err != nil {
		return C.CString(fmt.Sprintf("Invalid configuration file: %v", err))
	}

	return C.CString("Configuration file is valid")
}

// GetConfigTemplate returns a basic configuration template
//
//export GetConfigTemplate
func GetConfigTemplate() *C.char {
	template := `{
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
}`
	return C.CString(strings.TrimSpace(template))
}

func main() {
	// This is required for building as a shared library
	// The main function will not be called when used as a DLL
}
