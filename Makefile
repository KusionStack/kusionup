# Default variable
GOIMPORTS ?= goimports
GOCILINT ?= golangci-lint
APP_NAME ?= kusionup
COVER_FILE ?= cover.out

default:
	@go run cmd/main.go

build-all: build-darwin build-darwin-arm64 build-linux build-windows

build-darwin: gen-version
	-rm -rf ./build/darwin/bin
	# Build kusion
	GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 \
		go build -o ./build/darwin/bin/$(APP_NAME) \
		./cmd/main.go

build-darwin-arm64: gen-version
	-rm -rf ./build/darwin-arm64/bin
	# Build kusion
	GOOS=darwin GOARCH=arm64 CGO_ENABLED=0 \
		go build -o ./build/darwin-arm64/bin/$(APP_NAME) \
		./cmd/main.go

build-linux: gen-version
	-rm -rf ./build/linux/bin
	# Build kusion
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 \
		go build -o ./build/linux/bin/$(APP_NAME) \
		./cmd/main.go

build-windows: gen-version
	-rm -rf ./build/windows/bin
	# Build kusion
	GOOS=windows GOARCH=amd64 CGO_ENABLED=0 \
		go build -o ./build/windows/bin/$(APP_NAME).exe \
		./cmd/main.go

# If you encounter an error like "panic: permission denied" on MacOS,
# please visit https://github.com/eisenxp/macos-golink-wrapper to find the solution.
test:
	go test -gcflags=all=-l -timeout=10m -v `go list ./pkg/... ./cmd/...` ${TEST_FLAGS}

test-cover:
	go test -gcflags=all=-l -timeout=10m -v `go list ./pkg/... ./cmd/...` -coverprofile $(COVER_FILE) ${TEST_FLAGS}

test-html:
	go tool cover -html=$(COVER_FILE)

lint:
	@$(GOCILINT) run --no-config --disable=errcheck ./...

gen-version: # Update version
	cd pkg/version/scripts && go run gen.go

# Install git-chglog before execution:
#   go install github.com/git-chglog/git-chglog/cmd/git-chglog@latest
build-changelog:
	mkdir -p ./build
	git-chglog -o ./build/CHANGELOG.md
