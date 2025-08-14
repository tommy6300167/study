#!/usr/bin/env bash
set -euo pipefail

APP_NAME="${APP_NAME:-you-mon-so-fat}"
DOCKER_USER="${DOCKER_USER:-tommy6300167}"
TAG="${TAG:-latest}"
IMAGE="${DOCKER_USER}/${APP_NAME}:${TAG}"

echo "🚀 本地自動化部署開始..."

make build

make docker-build

echo "🔐 Docker Hub 登入"
docker login || true

echo "📦 推送到 Docker Hub: ${IMAGE}"
docker tag "${APP_NAME}:latest" "${IMAGE}"
docker push "${IMAGE}"

echo "▶️ 以 .env（若存在）啟動本地容器"
make docker-run || true

echo "✅ 部署完成（單架構已推送到 ${IMAGE}）"