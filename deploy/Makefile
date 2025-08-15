APP_NAME ?= you-mon-so-fat

# Docker Hub
DOCKER_USER ?= tommy6300167
IMAGE ?= $(DOCKER_USER)/$(APP_NAME)
TAG ?= latest

PLATFORMS ?= linux/amd64,linux/arm64

# ========== Go ==========

run:
	go run main.go

build: frontend
	go build -o main .

frontend:
	@echo "✅ Frontend files ready in web/ directory"
	@ls -la web/ || true

test:
	go test ./...

clean:
	rm -f main

dev:
	@echo "🚀 Starting development server..."
	go run main.go

docker-build:
	docker build -t $(IMAGE):$(TAG) .

docker-run:
	@if [ -f .env ]; then EF="--env-file .env"; else EF=""; fi; \
	docker run --rm -p 8080:8080 $$EF $(IMAGE):$(TAG)

# ========== Docker Hub ==========
docker-login:
	docker login

docker-push:
	docker push $(IMAGE):$(TAG)

docker-publish:
	# 初始化 buildx
	docker buildx create --use >/dev/null 2>&1 || true
	docker buildx build \
		--platform $(PLATFORMS) \
		-t $(IMAGE):$(TAG) \
		--push .

# ========== utils ==========

docker-stop-all:
	- docker stop $$(docker ps -q) 2>/dev/null || true
	- docker rm $$(docker ps -aq) 2>/dev/null || true

.PHONY: run build frontend test clean dev docker-build docker-run docker-login docker-push docker-publish docker-stop-all