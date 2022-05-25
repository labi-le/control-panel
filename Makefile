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
	go build -ldflags "-s" -a -v -o $(BUILD_PATH)$(PROJ_NAME) $(MAIN_PATH)

build: clean
	@echo "Building..."
	go build -v -o $(BUILD_PATH)$(PROJ_NAME) $(MAIN_PATH)

install: build-release uninstall
	@echo "Installing..."
	mv $(BUILD_PATH)$(PROJ_NAME) $(INSTALL_PATH)$(PROJ_NAME)

uninstall:
	@echo "Uninstalling..."
	rm $(INSTALL_PATH)$(PROJ_NAME)
	@echo "Remove configuration file..."
	rm -rf ~/.config/$(PROJ_NAME)/

clean:
	@echo "Cleaning..."
	rm -rf $(BUILD_PATH)*

tests:
	go test ./test/ -a