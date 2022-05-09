# Default variable
APPROOT 		?= kusionup
GOSOURCE_PATHS 	?= ./pkg/... ./cmd/...

include go.mk


.PHONY: gen-version
gen-version: ## Update version
	# Delete old version file
	-rm -f ./pkg/version/z_update_version.go
	# Generates new version file
	cd pkg/version/scripts && go run gen.go

.PHONY: clean
clean:  ## Clean build bundles
	-rm -rf ./build

.PHONY: build-all
build-all: build-darwin build-darwin-arm64 build-linux build-windows ## Build all platforms

.PHONY: build-darwin
build-darwin: gen-version ## Build for MacOS
	-rm -rf ./build/darwin
	GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 \
		go build -o ./build/darwin/$(APPROOT) \
		./cmd

.PHONY: build-darwin-arm64
build-darwin-arm64: gen-version ## Build for MacOS-arm64
	-rm -rf ./build/darwin-arm64
	GOOS=darwin GOARCH=arm64 CGO_ENABLED=0 \
		go build -o ./build/darwin-arm64/$(APPROOT) \
		./cmd

.PHONY: build-linux
build-linux: gen-version ## Build for Linux
	-rm -rf ./build/linux
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 \
		go build -o ./build/linux/$(APPROOT) \
		./cmd

.PHONY: build-windows
build-windows: gen-version ## Build for Windows
	-rm -rf ./build/windows
	GOOS=windows GOARCH=amd64 CGO_ENABLED=0 \
		go build -o ./build/windows/$(APPROOT).exe \
		./cmd

# Install git-chglog before execution:
#   go install github.com/git-chglog/git-chglog/cmd/git-chglog@latest
.PHONY: build-changelog
build-changelog: ## Build changelog
	mkdir -p ./build
	git-chglog -o ./build/CHANGELOG.md
