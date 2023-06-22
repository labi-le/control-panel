PACKAGE = control-panel

MAIN_PATH = cmd/main.go
BUILD_PATH = build/package/

.DEFAULT_GOAL := build
INSTALL_PATH = /usr/local/bin/

run:
	go run -v $(MAIN_PATH)

build-release: clean
	@echo "Building..."
	goreleaser release --clean

build: clean
	@echo "Building..."
	goreleaser release --skip-publish --snapshot --clean

clean:
	@echo "Cleaning..."
	rm -rf $(BUILD_PATH)*

tests:
	go test ./test/ -a