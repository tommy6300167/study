#!/bin/bash
set -e

echo "🚀 本地自動化部署開始..."

make build
make docker-build
make docker-run

echo "✅ 部署完成"