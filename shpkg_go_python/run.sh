#!/bin/bash

echo "🚀 Compiling SHPKG package manager (Go)..."
go build -o shpkg shpkg.go

if [ $? -ne 0 ]; then
    echo "❌ Go build failed. Please check shpkg.go"
    exit 1
fi

echo "✅ Build complete."

echo -e "\n📦 Installing 'hello' package..."
./shpkg install ./shpkg-packages/hello

echo -e "\n⌛ Sleeping for 2 seconds before uninstall..."
sleep 2

echo -e "\n🗑️ Uninstalling 'hello' package..."
./shpkg uninstall ./shpkg-packages/hello

echo -e "\n✅ All done!"
