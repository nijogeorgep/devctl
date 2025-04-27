APP_NAME := devctl
CMD_DIR := ./cmd/devctl
BIN := ./bin/$(APP_NAME)
VERSION := $(shell git describe --tags --always --dirty)
GIT_SHA := $(shell git rev-parse --short HEAD)
BUILD_DATE := $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")

LDFLAGS := -ldflags "-X 'main.version=$(VERSION)' -X 'main.gitSha=$(GIT_SHA)' -X 'main.buildDate=$(BUILD_DATE)'"

.PHONY: all build run install test clean tidy docker release

all: build

build:
	@echo "🔧 Building $(APP_NAME)..."
	@go build $(LDFLAGS) -o $(BIN) $(CMD_DIR)
	@echo "✅ Build complete."

run:
	@echo "🚀 Running $(APP_NAME)..."
	@$(BIN) $(CMD)

install:
	@echo "📦 Installing $(APP_NAME) to /usr/local/bin..."
	@sudo mv $(BIN) /usr/local/bin/$(APP_NAME)
	@echo "✅ Installed as /usr/local/bin/$(APP_NAME)"

test:
	@echo "🧪 Running tests..."
	@go test ./...

clean:
	@echo "🧹 Cleaning up..."
	@rm -f $(BIN)

tidy:
	@echo "📦 Tidying Go modules..."
	@go mod tidy

docker:
	@echo "🐳 Building Docker image..."
	@docker build -t $(APP_NAME):$(VERSION) .

release:
	@if [ -z "$(TYPE)" ]; then \
		echo "❗ Please specify TYPE=patch, TYPE=minor, or TYPE=major"; \
		exit 1; \
	fi
	@echo "🚀 Bumping $(TYPE) version..."
	@./scripts/bump_version.sh $(TYPE)
	@echo "🔧 Building after release bump..."
	@make build