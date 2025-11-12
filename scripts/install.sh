#!/usr/bin/env bash
set -euo pipefail
ROOT_DIR="$(cd "$(dirname "$0")/.." && pwd)"
cd "$ROOT_DIR"

echo "Building Go binary..."
if command -v go >/dev/null 2>&1; then
    go build -o opset ./go
else
    echo "go not found in PATH; please install Go to build the binary."
fi

echo "Compiling Java helpers (if javac available)..."
if command -v javac >/dev/null 2>&1; then
    javac java/*.java || true
else
    echo "javac not found; skipping Java compile."
fi

echo "Install step complete."
