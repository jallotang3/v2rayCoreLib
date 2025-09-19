#ifndef V2RAY_H
#define V2RAY_H

#ifdef __cplusplus
extern "C" {
#endif

// Function declarations for V2Ray DLL

/**
 * Start V2Ray with JSON configuration string
 * @param configData JSON configuration as C string
 * @return Result message (use FreeString to free memory)
 */
char* StartV2Ray(char* configData);

/**
 * Start V2Ray with configuration file path
 * @param configPath Path to JSON configuration file
 * @return Result message (use FreeString to free memory)
 */
char* StartV2RayWithConfig(char* configPath);

/**
 * Stop the running V2Ray instance
 * @return Result message (use FreeString to free memory)
 */
char* StopV2Ray();

/**
 * Get current V2Ray status
 * @return Status string: "running" or "stopped" (use FreeString to free memory)
 */
char* GetV2RayStatus();

/**
 * Check if V2Ray is running
 * @return 1 if running, 0 if stopped
 */
int IsV2RayRunning();

/**
 * Get V2Ray version
 * @return Version string (use FreeString to free memory)
 */
char* GetV2RayVersion();

/**
 * Restart V2Ray with new JSON configuration
 * @param configData JSON configuration as C string
 * @return Result message (use FreeString to free memory)
 */
char* RestartV2Ray(char* configData);

/**
 * Restart V2Ray with new configuration file
 * @param configPath Path to JSON configuration file
 * @return Result message (use FreeString to free memory)
 */
char* RestartV2RayWithConfig(char* configPath);

/**
 * Test if JSON configuration is valid
 * @param configData JSON configuration as C string
 * @return Result message (use FreeString to free memory)
 */
char* TestV2RayConfig(char* configData);

/**
 * Test if configuration file is valid
 * @param configPath Path to JSON configuration file
 * @return Result message (use FreeString to free memory)
 */
char* TestV2RayConfigFile(char* configPath);

/**
 * Free memory allocated for C strings returned by other functions
 * @param str String to free
 */
void FreeString(char* str);

#ifdef __cplusplus
}
#endif

#endif // V2RAY_H
