#!/usr/bin/env bash
set -euo pipefail
ROOT_DIR="$(cd "$(dirname "$0")/.." && pwd)"

echo "OpSet demo: workspace root is $ROOT_DIR"

echo "Sample data files:"
ls -la "$ROOT_DIR/data" || true

echo "To build and run the CLI:"
echo "  ./scripts/install.sh"
echo "  ./opset -config data/servers.json"
