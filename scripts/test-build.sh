#!/bin/bash

# V2Ray Core Library 构建测试脚本
# 用于验证本地构建环境是否配置正确

set -e

echo "=== V2Ray Core Library 构建环境测试 ==="
echo

# 检查 Go 版本
echo "1. 检查 Go 版本..."
go version
if [ $? -ne 0 ]; then
    echo "❌ Go 未安装或不在 PATH 中"
    exit 1
fi
echo "✅ Go 版本检查通过"
echo

# 检查 xgo 是否安装
echo "2. 检查 xgo 是否安装..."
if ! command -v xgo &> /dev/null; then
    echo "⚠️  xgo 未安装，正在安装..."
    go install src.techknowlogick.com/xgo@latest
    if [ $? -ne 0 ]; then
        echo "❌ xgo 安装失败"
        exit 1
    fi
    echo "✅ xgo 安装成功"
else
    echo "✅ xgo 已安装"
fi
echo

# 检查 C 编译器
echo "3. 检查 C 编译器..."
case "$(uname -s)" in
    Linux*)
        if command -v gcc &> /dev/null; then
            echo "✅ GCC 编译器可用: $(gcc --version | head -n1)"
        else
            echo "❌ GCC 编译器未安装"
            exit 1
        fi
        ;;
    Darwin*)
        if command -v clang &> /dev/null; then
            echo "✅ Clang 编译器可用: $(clang --version | head -n1)"
        else
            echo "❌ Clang 编译器未安装，请安装 Xcode Command Line Tools"
            exit 1
        fi
        ;;
    MINGW*|CYGWIN*|MSYS*)
        if command -v gcc &> /dev/null; then
            echo "✅ MinGW GCC 编译器可用: $(gcc --version | head -n1)"
        else
            echo "❌ MinGW GCC 编译器未安装"
            exit 1
        fi
        ;;
    *)
        echo "⚠️  未知操作系统，跳过 C 编译器检查"
        ;;
esac
echo

# 检查项目结构
echo "4. 检查项目结构..."
if [ ! -f "go.mod" ]; then
    echo "❌ go.mod 文件不存在"
    exit 1
fi

if [ ! -d "bind/desktop" ]; then
    echo "❌ bind/desktop 目录不存在"
    exit 1
fi

if [ ! -f "bind/desktop/main.go" ]; then
    echo "❌ bind/desktop/main.go 文件不存在"
    exit 1
fi

echo "✅ 项目结构检查通过"
echo

# 检查 Go 模块
echo "5. 检查 Go 模块..."
go mod verify
if [ $? -ne 0 ]; then
    echo "❌ Go 模块验证失败"
    exit 1
fi
echo "✅ Go 模块验证通过"
echo

# 测试 CGO 编译
echo "6. 测试 CGO 编译..."
cd bind/desktop
export CGO_ENABLED=1
if go build -buildmode=c-shared -o test-build.so main.go 2>/dev/null; then
    echo "✅ CGO 编译测试通过"
    rm -f test-build.so test-build.h
else
    echo "❌ CGO 编译测试失败"
    exit 1
fi
cd ../..
echo

# 测试 xgo 编译 (仅当前平台)
echo "7. 测试 xgo 编译..."
cd bind/desktop
case "$(uname -s)" in
    Linux*)
        TARGET="linux/amd64"
        ;;
    Darwin*)
        TARGET="darwin/amd64"
        ;;
    MINGW*|CYGWIN*|MSYS*)
        TARGET="windows/amd64"
        ;;
    *)
        echo "⚠️  跳过 xgo 测试 (未知平台)"
        TARGET=""
        ;;
esac

if [ -n "$TARGET" ]; then
    if xgo --targets=$TARGET -buildmode=c-shared -out=test-xgo ./ 2>/dev/null; then
        echo "✅ xgo 编译测试通过"
        rm -f test-xgo-*
    else
        echo "❌ xgo 编译测试失败"
        exit 1
    fi
fi
cd ../..
echo

echo "=== 构建环境测试完成 ==="
echo "✅ 所有测试通过，构建环境配置正确！"
echo
echo "现在可以运行以下命令进行构建："
echo "  cd bind/desktop"
echo "  xgo --targets=linux/amd64,linux/386 -buildmode=c-shared -out=v2ray ./"
echo
