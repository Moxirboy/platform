-include .env
export

# Define variables
POSTGRES_USER ?= postgres
POSTGRES_PASSWORD ?= postgres
POSTGRES_HOST ?= db
POSTGRES_PORT ?= 5432
POSTGRES_DATABASE ?= postgres
#docker compose run --rm migrate -path=./migrations/ -database='postgresql://postgres:postgres@db:5432/postgres?sslmode=disable' force 3 up

FOLDERS ?= ./cmd/... ./pkg/... ./internal/...
PROJECT := $(shell basename $(CURRENT_DIR))
CURRENT_DIR := $(shell pwd)
PKG_LIST := $(shell go list ./...)
UNAME := $(shell uname)
DOCKER_COMPOSE_FILE := docker-compose.yaml

# Define default Docker service
DOCKER_SERVICES ?= app

# Define arguments
arg := $(filter-out $@,$(MAKECMDGOALS))

.PHONY: start
start:
    @echo "Start Containers"
    docker compose -f ${DOCKER_COMPOSE_FILE} up -d ${DOCKER_SERVICES}
    sleep 2
    docker compose -f ${DOCKER_COMPOSE_FILE} ps

.PHONY: stop
stop:
    @echo "Stop Containers"
    docker compose -f ${DOCKER_COMPOSE_FILE} stop ${DOCKER_SERVICES}
    sleep 2
    docker compose -f ${DOCKER_COMPOSE_FILE} ps

.PHONY: rm
rm: stop
    @echo "Remove Containers"
    docker compose -f ${DOCKER_COMPOSE_FILE} rm -v -f ${DOCKER_SERVICES}

.PHONY: install-dep
install-dep:
    go install github.com/segmentio/golines@v0.5.0
    go install github.com/go-courier/husky/cmd/husky@v1.7.2
    go install github.com/golang/mock/mockgen@v1.6.0
    go install github.com/swaggo/swag/cmd/swag@v1.7.2
ifeq ($(UNAME), Linux)
    curl -sSfL https://packagecloud.io/golang-migrate/migrate/gpgkey | apt-key add - \
    echo "deb https://packagecloud.io/golang-migrate/migrate/ubuntu/ $(lsb_release -sc) main" > /etc/apt/sources.list.d/migrate.list \
    apt-get update \
    apt-get install -y migrate \
    \
    apt-get install -y protobuf-compiler
endif
ifeq ($(UNAME), Darwin)
    brew install golang-migrate
    brew install protobuf
endif
    husky init

.PHONY: golines
golines:
    golines -w --max-len=80 --tab-len=4 ./cmd ./internal ./pkg 

.PHONY: all
all: lint test

.PHONY: set-env
set-env:
    export $(egrep -v '^#' .env | xargs)

.PHONY: lint
lint: golines fmt
    golangci-lint run ${FOLDERS}

.PHONY: test
test:
    go test -v ${FOLDERS} -coverprofile=coverage.out

.PHONY: test-coverage
test-coverage:
    go test -cover ./...

.PHONY: run
run: 
    go run ./cmd/${PROJECT}/main.go

.PHONY: fmt
fmt:
    gofmt -l -w .

.PHONY: unit-tests
unit-tests: set-env
    go test -mod-readonly -v -cover -short ${PKG_LIST}

.PHONY: race
race: set-env
    go test -mod=readonly -race -short ${PKG_LIST}

.PHONY: msan
msan: set-env
    env CC=clang env CXX=clang++ go test -mod=readonly -msan -short ${PKG_LIST}

.PHONY: proto-gen
proto-gen:
    ./scripts/gen-proto.sh  ${CURRENT_DIR}

.PHONY: ci-lint
ci-lint:
    @golangci-lint run --deadline 5m --verbose --build-tags=${UNAME} --paths=${FOLD}

.PHONY: git
git:
    git add .
    git commit -m "$m"
    git push origin $(b)

.PHONY: git-pull
git-pull:
    git pull origin master

.PHONY: migration-up
migration-up:
    @echo "Migrations Up"
    sleep 2
    docker compose run --rm migrate -path=./migrations/ -database='postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DATABASE}?sslmode=disable' force 3 up

.PHONY: migration-generate
migration-generate:
    @echo "Generation migration file $(name)"
    sleep 2
    docker compose run --rm migrate create -ext sql -dir ./migrations -seq $(name)

.PHONY: mod-download
mod-download:
    @echo "Go mod download"
    sleep 2
    docker compose exec app go mod download

.PHONY: go-get
go-get:
    @echo "Go get ${arg}"
    sleep 2
    docker compose exec app go get -d ${arg}

.PHONY: mod-tidy
mod-tidy:
    @echo "Go mod tidy"
    sleep 2
    docker compose exec app go mod tidy
