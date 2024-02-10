FOLDERS ?= ./cmd/... ./pkg/... ./internal/...
PROJECT=$(shell basename ${CURRENT_DIR})
CURRENT_DIR=$(shell pwd)
PKG_LIST := $(shell go list ./...)
UNAME := $(shell uname)

install-dep:
	go install github.com/segmentio/golines@v0.5.0
	go install github.com/go-courier/husky/cmd/husky@v1.7.2
	go install github.com/golang/mock/mockgen@v1.6.0
	go install github.com/swaggo/swag/cmd/swag@v1.7.2
ifeq ($(UNAME), Linux)
	curl -L https://packagecloud.io/golang-migrate/migrate/gpgkey | apt-key add - \
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





golines:
	golines -w --max-len=80 --tab-len=4 ./cmd ./internal ./pkg 

all: lint test

set-env:
	export $(egrep -v '^#' .env | xargs)
# Run lint
lint: golines fmt
	golangci-lint run ${FOLDERS}
# Run tests
test:
	go test -v ${FOLDERS} -coverprofile=coverage.out

test-coverage:
	go test -cover ./...

run: 
	go run ./cmd/${PROJECT}/main.go

fmt:
	gofmt -l -w .

unit-tests: set-env ## Run unit-tests
	go test -mod-readonly -v -cover -short ${PKG_LIST}

race: set-env ## Run data race detector
	go test -mod=readonly -race -short ${PKG_LIST}

msan: set-env ## Run memory sanitizer. If this test fails, you need to write the following command: export CC=clang (if you have installed clang)
	env CC=clang env CXX=clang++ go test -mod=readonly -msan -short ${PKG_LIST}

gen-proto:
	./scripts/gen-proto.sh ${CURRENT_DIR}


ci-lint:
	@golangci-lint run --deadline 5m --verbose --build-tags=${UNAME} --paths=${FOLD}




git:
	git add .
	git commit -m "$m"
	git push origin $(b)

git-pull:
	git pull origin master
	
.PHONY: migration-up
migration-up:
	@echo "Migrations Up"
	sleep 2
	docker compose run --rm migrate -path=migrations/ -database='postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DATABASE}?sslmode=disable' up

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