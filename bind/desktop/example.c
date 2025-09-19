#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <windows.h>

// Function pointer types for DLL functions
typedef char* (*StartV2RayFunc)(char*);
typedef char* (*StartV2RayWithConfigFunc)(char*);
typedef char* (*StopV2RayFunc)();
typedef char* (*GetV2RayStatusFunc)();
typedef int (*IsV2RayRunningFunc)();
typedef char* (*GetV2RayVersionFunc)();
typedef char* (*RestartV2RayFunc)(char*);
typedef char* (*TestV2RayConfigFunc)(char*);
typedef void (*FreeStringFunc)(char*);

int main() {
    // Load the DLL
    HMODULE hDll = LoadLibrary(L"v2ray.dll");
    if (!hDll) {
        printf("Failed to load v2ray.dll\n");
        return 1;
    }

    // Get function pointers
    StartV2RayFunc StartV2Ray = (StartV2RayFunc)GetProcAddress(hDll, "StartV2Ray");
    StopV2RayFunc StopV2Ray = (StopV2RayFunc)GetProcAddress(hDll, "StopV2Ray");
    GetV2RayStatusFunc GetV2RayStatus = (GetV2RayStatusFunc)GetProcAddress(hDll, "GetV2RayStatus");
    IsV2RayRunningFunc IsV2RayRunning = (IsV2RayRunningFunc)GetProcAddress(hDll, "IsV2RayRunning");
    GetV2RayVersionFunc GetV2RayVersion = (GetV2RayVersionFunc)GetProcAddress(hDll, "GetV2RayVersion");
    TestV2RayConfigFunc TestV2RayConfig = (TestV2RayConfigFunc)GetProcAddress(hDll, "TestV2RayConfig");
    FreeStringFunc FreeString = (FreeStringFunc)GetProcAddress(hDll, "FreeString");

    if (!StartV2Ray || !StopV2Ray || !GetV2RayStatus || !IsV2RayRunning || 
        !GetV2RayVersion || !TestV2RayConfig || !FreeString) {
        printf("Failed to get function pointers from DLL\n");
        FreeLibrary(hDll);
        return 1;
    }

    printf("V2Ray DLL Example\n");
    printf("=================\n\n");

    // Get version
    char* version = GetV2RayVersion();
    printf("V2Ray Version: %s\n\n", version);
    FreeString(version);

    // Check initial status
    char* status = GetV2RayStatus();
    printf("Initial Status: %s\n", status);
    FreeString(status);

    // Example configuration (minimal SOCKS proxy)
    const char* config = "{"
        "\"inbounds\": [{"
            "\"port\": 1080,"
            "\"protocol\": \"socks\","
            "\"settings\": {"
                "\"udp\": true"
            "}"
        "}],"
        "\"outbounds\": [{"
            "\"protocol\": \"freedom\""
        "}]"
    "}";

    // Test configuration
    printf("\nTesting configuration...\n");
    char* testResult = TestV2RayConfig((char*)config);
    printf("Test result: %s\n", testResult);
    FreeString(testResult);

    // Start V2Ray
    printf("\nStarting V2Ray...\n");
    char* startResult = StartV2Ray((char*)config);
    printf("Start result: %s\n", startResult);
    FreeString(startResult);

    // Check status
    status = GetV2RayStatus();
    printf("Status after start: %s\n", status);
    FreeString(status);

    // Check if running
    int running = IsV2RayRunning();
    printf("Is running: %s\n", running ? "Yes" : "No");

    // Wait for user input
    printf("\nPress Enter to stop V2Ray...");
    getchar();

    // Stop V2Ray
    printf("Stopping V2Ray...\n");
    char* stopResult = StopV2Ray();
    printf("Stop result: %s\n", stopResult);
    FreeString(stopResult);

    // Check final status
    status = GetV2RayStatus();
    printf("Final status: %s\n", status);
    FreeString(status);

    // Cleanup
    FreeLibrary(hDll);
    printf("\nExample completed.\n");
    return 0;
}
