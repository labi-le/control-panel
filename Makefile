.PHONY: build
.DEFAULT_GOAL := build

build-release:
	@echo "Building..."
	go build -ldflags "-s" -a -v -o build/package/control-panel-release cmd/main.go

build:
	@echo "Building..."
	go build -v -o build/package/control-panel-debug cmd/main.go

install: build-release uninstall
	@echo "Installing..."
	mv build/package/control-panel-release /usr/local/bin/control-panel

uninstall:
	@echo "Uninstalling..."
	rm /usr/local/bin/control-panel
	@echo "Remove configuration file..."
	rm -rf ~/.config/control-panel/

clean:
	@echo "Cleaning..."
	rm -rf build/package/*
