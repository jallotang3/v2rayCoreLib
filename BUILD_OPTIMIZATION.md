# V2Ray Core Library 构建优化说明

## 概述

本次优化主要针对 `.github/workflows/build-libraries.yml` 构建脚本，使用 `cgo` 和 `xgo` 进行跨平台构建，支持更多架构和平台。

## 主要优化内容

### 1. Windows 构建优化
- **多架构支持**: 同时构建 x86 (32位) 和 x64 (64位) 版本
- **CGO 配置**: 正确配置 CGO 环境变量和编译器
- **构建缓存**: 添加 Go 模块和构建缓存，提高构建速度
- **产物验证**: 增强构建产物验证，显示文件大小信息

### 2. macOS 构建优化
- **通用二进制**: 使用 `lipo` 创建支持 Intel 和 Apple Silicon 的通用二进制文件
- **xgo 集成**: 使用 xgo 进行跨平台编译
- **架构分离**: 分别构建 Intel (amd64) 和 ARM64 版本，然后合并

### 3. Linux 构建支持 (新增)
- **多架构支持**: 支持 amd64, 386, arm64, arm 架构
- **共享库构建**: 构建 `.so` 共享库文件
- **交叉编译**: 使用 xgo 进行交叉编译

### 4. 构建配置优化
- **环境变量**: 统一管理构建环境变量
- **缓存策略**: 优化 Go 模块和构建缓存
- **依赖管理**: 改进构建依赖安装和管理

## 构建产物

### Windows
- `v2ray-x64-windows-amd64.dll` - 64位 Windows DLL
- `v2ray-x86-windows-386.dll` - 32位 Windows DLL
- `v2ray.h` - C 头文件

### macOS
- `v2ray.dylib` - 通用二进制文件 (Intel + Apple Silicon)
- `v2ray-intel-darwin-amd64.dylib` - Intel Mac 专用
- `v2ray-arm64-darwin-arm64.dylib` - Apple Silicon 专用
- `v2ray.h` - C 头文件

### Linux
- `v2ray-x64-linux-amd64.so` - 64位 Linux 共享库
- `v2ray-x86-linux-386.so` - 32位 Linux 共享库
- `v2ray-arm64-linux-arm64.so` - ARM64 Linux 共享库
- `v2ray-arm-linux-arm.so` - ARM Linux 共享库
- `v2ray.h` - C 头文件

## 技术特性

### CGO 配置
- 启用 CGO 进行 C 语言互操作
- 正确配置跨平台编译器
- 支持 Windows MinGW, macOS Clang, Linux GCC

### xgo 集成
- 使用 xgo 进行跨平台交叉编译
- 支持多目标架构同时构建
- 自动处理不同平台的构建环境

### 构建优化
- 使用 `-trimpath` 移除构建路径信息
- 使用 `-ldflags="-w -s"` 减小二进制文件大小
- 启用构建缓存提高构建速度

## 使用方法

### 本地构建
```bash
# 安装 xgo
go install src.techknowlogick.com/xgo@latest

# Windows 构建
cd bind/desktop
xgo --targets=windows/amd64,windows/386 -buildmode=c-shared -out=v2ray ./

# macOS 构建
xgo --targets=darwin/amd64,darwin/arm64 -buildmode=c-shared -out=v2ray ./

# Linux 构建
xgo --targets=linux/amd64,linux/386,linux/arm64,linux/arm -buildmode=c-shared -out=v2ray ./
```

### CI/CD 构建
构建脚本会在以下情况下自动触发：
- 推送到 master, main, v* 分支
- 创建 Pull Request
- 发布 Release
- 手动触发 (workflow_dispatch)

## 构建环境要求

### Windows
- Go 1.21+
- MinGW-w64 GCC
- xgo

### macOS
- Go 1.21+
- Xcode Command Line Tools
- xgo

### Linux
- Go 1.21+
- GCC multilib
- xgo

## 注意事项

1. **CGO 依赖**: 构建需要 C 编译器，确保构建环境已安装相应的编译工具链
2. **交叉编译**: xgo 会自动处理交叉编译环境，但某些 C 库可能需要特殊配置
3. **文件大小**: 使用 `-ldflags="-w -s"` 可以显著减小二进制文件大小
4. **兼容性**: 通用二进制文件兼容性更好，但文件大小较大

## 故障排除

### 常见问题
1. **CGO 编译失败**: 检查 C 编译器是否正确安装
2. **xgo 安装失败**: 确保 Go 版本 >= 1.18
3. **交叉编译失败**: 检查目标平台的构建环境配置

### 调试方法
1. 查看构建日志中的详细错误信息
2. 检查环境变量配置
3. 验证构建依赖是否正确安装
