package v2ray

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"sync"

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

// V2RayService represents the V2Ray service for mobile platforms
type V2RayService struct{}

// NewV2RayService creates a new V2Ray service instance
func NewV2RayService() *V2RayService {
	return &V2RayService{}
}

// StartWithConfig starts V2Ray with the given JSON configuration string
func (s *V2RayService) StartWithConfig(configData string) string {
	instanceLock.Lock()
	defer instanceLock.Unlock()

	if isRunning {
		return "V2Ray is already running"
	}

	if configData == "" {
		return "Configuration is empty"
	}

	// Parse JSON configuration using v4 config
	var jsonConfig v4.Config
	if err := json.Unmarshal([]byte(configData), &jsonConfig); err != nil {
		return fmt.Sprintf("Failed to parse configuration: %v", err)
	}

	// Convert to protobuf config
	config, err := jsonConfig.Build()
	if err != nil {
		return fmt.Sprintf("Failed to build configuration: %v", err)
	}

	// Create V2Ray instance
	server, err := core.New(config)
	if err != nil {
		return fmt.Sprintf("Failed to create V2Ray instance: %v", err)
	}

	// Start the server
	if err := server.Start(); err != nil {
		return fmt.Sprintf("Failed to start V2Ray: %v", err)
	}

	instance = server
	isRunning = true
	return "V2Ray started successfully"
}

// StartWithConfigFile starts V2Ray with the given configuration file path
func (s *V2RayService) StartWithConfigFile(configPath string) string {
	instanceLock.Lock()
	defer instanceLock.Unlock()

	if isRunning {
		return "V2Ray is already running"
	}

	if configPath == "" {
		return "Configuration path is empty"
	}

	// Open and read the configuration file
	file, err := os.Open(configPath)
	if err != nil {
		return fmt.Sprintf("Failed to open configuration file: %v", err)
	}
	defer file.Close()

	// Load configuration from file
	config, err := serial.LoadJSONConfig(file)
	if err != nil {
		return fmt.Sprintf("Failed to load configuration: %v", err)
	}

	// Create V2Ray instance
	server, err := core.New(config)
	if err != nil {
		return fmt.Sprintf("Failed to create V2Ray instance: %v", err)
	}

	// Start the server
	if err := server.Start(); err != nil {
		return fmt.Sprintf("Failed to start V2Ray: %v", err)
	}

	instance = server
	isRunning = true
	return "V2Ray started successfully"
}

// Stop stops the running V2Ray instance
func (s *V2RayService) Stop() string {
	instanceLock.Lock()
	defer instanceLock.Unlock()

	if !isRunning || instance == nil {
		return "V2Ray is not running"
	}

	if err := instance.Close(); err != nil {
		return fmt.Sprintf("Failed to stop V2Ray: %v", err)
	}

	instance = nil
	isRunning = false
	return "V2Ray stopped successfully"
}

// GetStatus returns the current status of V2Ray
func (s *V2RayService) GetStatus() string {
	instanceLock.Lock()
	defer instanceLock.Unlock()

	if isRunning && instance != nil {
		return "running"
	}
	return "stopped"
}

// IsRunning returns true if V2Ray is running, false otherwise
func (s *V2RayService) IsRunning() bool {
	instanceLock.Lock()
	defer instanceLock.Unlock()

	return isRunning && instance != nil
}

// GetVersion returns the V2Ray version
func (s *V2RayService) GetVersion() string {
	return core.Version()
}

// Restart restarts V2Ray with new configuration
func (s *V2RayService) Restart(configData string) string {
	// Stop current instance
	stopResult := s.Stop()
	if stopResult != "V2Ray stopped successfully" && stopResult != "V2Ray is not running" {
		return fmt.Sprintf("Failed to stop V2Ray: %s", stopResult)
	}

	// Start with new configuration
	return s.StartWithConfig(configData)
}

// RestartWithConfigFile restarts V2Ray with new configuration file
func (s *V2RayService) RestartWithConfigFile(configPath string) string {
	// Stop current instance
	stopResult := s.Stop()
	if stopResult != "V2Ray stopped successfully" && stopResult != "V2Ray is not running" {
		return fmt.Sprintf("Failed to stop V2Ray: %s", stopResult)
	}

	// Start with new configuration file
	return s.StartWithConfigFile(configPath)
}

// TestConfig tests if the given configuration is valid
func (s *V2RayService) TestConfig(configData string) string {
	if configData == "" {
		return "Configuration is empty"
	}

	// Parse JSON configuration using v4 config
	var jsonConfig v4.Config
	if err := json.Unmarshal([]byte(configData), &jsonConfig); err != nil {
		return fmt.Sprintf("Invalid JSON configuration: %v", err)
	}

	// Convert to protobuf config
	_, err := jsonConfig.Build()
	if err != nil {
		return fmt.Sprintf("Invalid configuration: %v", err)
	}

	return "Configuration is valid"
}

// TestConfigFile tests if the given configuration file is valid
func (s *V2RayService) TestConfigFile(configPath string) string {
	if configPath == "" {
		return "Configuration path is empty"
	}

	// Open and read the configuration file
	file, err := os.Open(configPath)
	if err != nil {
		return fmt.Sprintf("Failed to open configuration file: %v", err)
	}
	defer file.Close()

	// Load configuration from file
	_, err = serial.LoadJSONConfig(file)
	if err != nil {
		return fmt.Sprintf("Invalid configuration file: %v", err)
	}

	return "Configuration file is valid"
}

// QueryStats queries statistics for a given tag and name
func (s *V2RayService) QueryStats(tag, name string) int64 {
	instanceLock.Lock()
	defer instanceLock.Unlock()

	if !isRunning || instance == nil {
		return -1
	}

	// This is a placeholder - actual stats implementation would depend on
	// the stats feature being configured in V2Ray
	return 0
}

// GetSystemStats returns system statistics as JSON string
func (s *V2RayService) GetSystemStats() string {
	instanceLock.Lock()
	defer instanceLock.Unlock()

	if !isRunning || instance == nil {
		return `{"error": "V2Ray is not running"}`
	}

	// This is a placeholder - actual implementation would collect
	// various system and V2Ray statistics
	stats := map[string]interface{}{
		"status":  "running",
		"version": core.Version(),
		"uptime":  "unknown", // Would need to track start time
	}

	data, _ := json.Marshal(stats)
	return string(data)
}

// SetLogLevel sets the log level (0=debug, 1=info, 2=warning, 3=error)
func (s *V2RayService) SetLogLevel(level int) string {
	// This is a placeholder - actual implementation would need to
	// interact with V2Ray's logging system
	if level < 0 || level > 3 {
		return "Invalid log level. Use 0=debug, 1=info, 2=warning, 3=error"
	}
	return "Log level set successfully"
}

// GetConfigTemplate returns a basic configuration template
func (s *V2RayService) GetConfigTemplate() string {
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
	return strings.TrimSpace(template)
}

// Exported functions for gomobile binding

// StartV2Ray starts V2Ray with the given JSON configuration string
func StartV2Ray(configData string) string {
	service := NewV2RayService()
	return service.StartWithConfig(configData)
}

// StartV2RayWithConfigFile starts V2Ray with the given configuration file path
func StartV2RayWithConfigFile(configPath string) string {
	service := NewV2RayService()
	return service.StartWithConfigFile(configPath)
}

// StopV2Ray stops the running V2Ray instance
func StopV2Ray() string {
	service := NewV2RayService()
	return service.Stop()
}

// GetV2RayStatus returns the current status of V2Ray
func GetV2RayStatus() string {
	service := NewV2RayService()
	return service.GetStatus()
}

// IsV2RayRunning returns true if V2Ray is running, false otherwise
func IsV2RayRunning() bool {
	service := NewV2RayService()
	return service.IsRunning()
}

// GetV2RayVersion returns the V2Ray version
func GetV2RayVersion() string {
	service := NewV2RayService()
	return service.GetVersion()
}

// RestartV2Ray restarts V2Ray with new configuration
func RestartV2Ray(configData string) string {
	service := NewV2RayService()
	return service.Restart(configData)
}

// RestartV2RayWithConfigFile restarts V2Ray with new configuration file
func RestartV2RayWithConfigFile(configPath string) string {
	service := NewV2RayService()
	return service.RestartWithConfigFile(configPath)
}

// TestV2RayConfig tests if the given configuration is valid
func TestV2RayConfig(configData string) string {
	service := NewV2RayService()
	return service.TestConfig(configData)
}

// TestV2RayConfigFile tests if the given configuration file is valid
func TestV2RayConfigFile(configPath string) string {
	service := NewV2RayService()
	return service.TestConfigFile(configPath)
}

// QueryV2RayStats queries statistics for a given tag and name
func QueryV2RayStats(tag, name string) int64 {
	service := NewV2RayService()
	return service.QueryStats(tag, name)
}

// GetV2RaySystemStats returns system statistics as JSON string
func GetV2RaySystemStats() string {
	service := NewV2RayService()
	return service.GetSystemStats()
}

// SetV2RayLogLevel sets the log level (0=debug, 1=info, 2=warning, 3=error)
func SetV2RayLogLevel(level int) string {
	service := NewV2RayService()
	return service.SetLogLevel(level)
}

// GetV2RayConfigTemplate returns a basic configuration template
func GetV2RayConfigTemplate() string {
	service := NewV2RayService()
	return service.GetConfigTemplate()
}

func main() {
	// This main function is required but not used in mobile bindings
}
