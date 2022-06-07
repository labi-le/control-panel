.PHONY: run
.DEFAULT_GOAL := run

PROJ_NAME = control-panel

MAIN_PATH = cmd/main.go
BUILD_PATH = build/package/

INSTALL_PATH = /usr/local/bin/

export CGO_ENABLED = 1

run:
	go run $(MAIN_PATH)

build-release: clean
	@echo "Building..."
	goreleaser release --rm-dist

build: clean
	@echo "Building..."
	goreleaser release --skip-publish --snapshot --rm-dist

clean:
	@echo "Cleaning..."
	rm -rf $(BUILD_PATH)*

tests:
	go test ./test/ -a