GOLANG_VERSION = ${shell cat .go-version}

CSPELL_VERSION = latest
GOCYCLO_VERSION = latest
GOIMPORTS_VERSION = latest
GOLANGCI_LINT_VERSION = latest
GOVERALLS_VERSION = latest

TARGET = hidori/go-tools:latest

.PHONY: tool/install
tool/install:
	go install github.com/fzipp/gocyclo/cmd/gocyclo@${GOCYCLO_VERSION}
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@${GOLANGCI_LINT_VERSION}
	go install github.com/mattn/goveralls@${GOVERALLS_VERSION}
	go install golang.org/x/tools/cmd/goimports@${GOIMPORTS_VERSION}

.PHONY: mod/download
mod/download:
	go mod download

.PHONY: mod/tidy
mod/tidy:
	go mod tidy

.PHONY: spell
spell:
	docker run -t --rm \
		-v ${shell pwd}:${shell pwd} \
		-w ${shell pwd} \
		ghcr.io/streetsidesoftware/cspell:${CSPELL_VERSION} "**"

.PHONY: cyclo
cyclo:
	gocyclo -top 30 .

.PHONY: lint
lint:
	golangci-lint run ./...

.PHONY: test
test:
	go test -race -cover -covermode atomic -coverprofile=profile.cov ./...

.PHONY: build
build:
	docker build -f ./Dockerfile -t ${TARGET} \
		--build-arg GOLANG_VERSION=${GOLANG_VERSION} \
		.

.PHONY: run
run:
	docker run -it --rm \
		-v ${shell pwd}:${shell pwd} \
		-w ${shell pwd} \
		${TARGET}
