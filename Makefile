SHELL_PATH = /bin/ash
SHELL = $(if $(wildcard $(SHELL_PATH)),/bin/ash,/bin/bash)

GOLANG          := golang:1.23.1
ALPINE          := alpine:3.21
POSTGRES        := postgres:17.3

SLOT_APP := slot
BASE_IMAGE_NAME := localhost/charlieroth
VERSION := 0.0.1
SLOT_IMAGE := ${BASE_IMAGE_NAME}/${SLOT_APP}:${VERSION}

dev-gotooling:
	go install github.com/divan/expvarmon@latest && \
	go install github.com/rakyll/hey@latest && \
	go install honnef.co/go/tools/cmd/staticcheck@latest && \
	go install golang.org/x/vuln/cmd/govulncheck@latest && \
	go install golang.org/x/tools/cmd/goimports@latest

dev-brew:
	brew update
	brew list pgcli || brew install pgcli
	brew list watch || brew install watch

dev-docker:
	docker pull $(GOLANG) & \
	docker pull $(ALPINE) & \
	docker pull $(POSTGRES) & \
	wait;

build:
	docker build \
		-f zarf/docker/dockerfile.slot \
		-t $(SLOT_IMAGE) \
		--build-arg BUILD_REF=$(VERSION) \
		--build-arg BUILD_DATE=$(date -u +"%Y-%m-%dT%H:%M:%SZ") \
		.

compose-up:
	cd ./zarf/compose/ && docker compose -f docker-compose.yaml -p compose up -d

compose-build-up: build compose-up

compose-down:
	cd ./zarf/compose/ && docker compose -f docker-compose.yaml -p compose down

compose-logs:
	cd ./zarf/compose/ && docker compose -f docker-compose.yaml logs

pgcli:
	pgcli postgresql://postgres:postgres@localhost

test-r:
	CGO_ENABLED=1 go test -race -count=1 ./...

test-only:
	CGO_ENABLED=0 go test -count=1 ./...

lint:
	CGO_ENABLED=0 go vet ./...
	staticcheck -checks=all ./...

vuln-check:
	govulncheck ./...

test: test-only lint vuln-check

test-race: test-race-only lint vuln-check

deps-reset:
	git checkout -- go.mod
	go mod tidy
	go mod vendor

tidy:
	go mod tidy
	go mod vendor

deps-list:
	go list -m -u -mod=readonly all

deps-upgrade:
	go get -u -v ./...
	go mod tidy
	go mod vendor

deps-cleancache:
	go clean -modcache

list:
	go list -mod=mod all

run:
	go run api/services/slot/main.go

help:
	@echo "Usage: make <target>"
	@echo ""
	@echo "Targets:"
	@echo "  build: Build the slot service"
	@echo "  compose-up: Start the compose environment"
	@echo "  compose-build-up: Build and start the compose environment"
	@echo "  compose-down: Stop the compose environment"
	@echo "  compose-logs: Show the logs of the compose environment"
	@echo "  pgcli: Connect to the postgres database"
	@echo "  test: Run the tests"
	@echo "  test-race: Run the tests with race detection"
	@echo "  tidy: Tidy the go modules"
	@echo "  deps-reset: Reset the go modules"
	@echo "  deps-list: List the go modules"
	@echo "  deps-upgrade: Upgrade the go modules"
	@echo "  deps-cleancache: Clean the go module cache"
	@echo "  list: List the go modules"
	@echo "  run: Run the slot service"