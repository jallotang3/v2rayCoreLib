@echo off
echo V2Ray DLL Test Script
echo =====================

REM Check if DLL exists
if not exist "v2ray.dll" (
    echo Error: v2ray.dll not found!
    echo Please run build.bat first to build the DLL.
    pause
    exit /b 1
)

echo DLL found: v2ray.dll
echo.

REM Check if example exists
if not exist "example.exe" (
    echo Building example application...
    gcc -o example.exe example.c
    if %ERRORLEVEL% NEQ 0 (
        echo Error: Failed to build example.exe
        echo Make sure GCC is installed and in PATH
        pause
        exit /b 1
    )
    echo Example built successfully!
    echo.
)

REM Run the example
echo Running DLL test...
echo.
example.exe

echo.
echo Test completed!
pause
