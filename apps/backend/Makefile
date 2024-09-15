# Check to see if we can use ash, in Alpine images, or default to BASH.
SHELL_PATH = /bin/ash
SHELL = $(if $(wildcard $(SHELL_PATH)),/bin/ash,/bin/bash)

# ==============================================================================
# Install Tooling and Dependencies
#
#	Run these commands to install everything needed.
#	$ make dev-brew
#	$ make dev-docker
#	$ make dev-gotooling

# ==============================================================================
# Define dependencies

GOLANG          := golang:1.23
ALPINE          := alpine:3.20
POSTGRES        := postgres:16.4
KIND            := kindest/node:v1.31.0
GRAFANA         := grafana/grafana:11.1.0
PROMETHEUS      := prom/prometheus:v2.54.0

BASE_IMAGE_NAME := iamNilotpal/openpulse
NAMESPACE       := openpulse-main-backend
APP_NAME       	:= openpulse-main-backend
KIND_CLUSTER    := openpulse-main-backend-cluster
APP_IMAGE     	:= $(BASE_IMAGE_NAME)/$(APP_NAME):$(VERSION)
VERSION       	:= "1.0.0-$(shell git rev-parse --short HEAD)"

# ==============================================================================
# Install dependencies

dev-gotooling:
	go install golang.org/x/vuln/cmd/govulncheck@latest
	go install honnef.co/go/tools/cmd/staticcheck@latest

dev-brew:
	brew update
	brew list kind || brew install kind
	brew list kubectl || brew install kubectl
	brew list kustomize || brew install kustomize

dev-docker:
	docker pull $(GOLANG) & \
	docker pull $(ALPINE) & \
	docker pull $(KIND) & \
	docker pull $(POSTGRES) & \
	docker pull $(GRAFANA) & \
	docker pull $(PROMETHEUS) & \
	wait;

# ==============================================================================
# Linting and Vulnarability check
lint:
	CGO_ENABLED=0 go vet ./...
	staticcheck -checks=all ./...

vuln-check:
	govulncheck ./...

# ==============================================================================
# Modules support

deps-reset:
	git checkout -- go.mod
	go mod tidy

tidy:
	go mod tidy

deps-list:
	go list -m -u -mod=readonly all

deps-upgrade:
	go get -u -v ./...
	go mod tidy

deps-cleancache:
	go clean -modcache

list:
	go list -mod=mod all

# ==============================================================================
# Test the application locally

run:
	go run apps/api/main.go

ready:
	curl -il http://localhost:3000/v1/readiness

live:
	curl -il http://localhost:3000/v1/liveness