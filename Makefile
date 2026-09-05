.PHONY: icons all

SHELL := /bin/bash

SRC_IMG = ./raw
BUILD = ./public
BUILD_IMAGES = $(BUILD)/images
resolutions := 16 32 180 192 512
ALL_ICONS := $(foreach resolution, $(resolutions), $(BUILD_IMAGES)/icons-$(resolution).png)

all: icons public/tgo.wasm public/wasm_exec.js public/sw.js

icons: public/images/icons-vector.svg $(ALL_ICONS)

$(BUILD_IMAGES)/icons-%.png: $(SRC_IMG)/icon.svg
	inkscape $(SRC_IMG)/icon.svg -w $* -h $* --export-png=$@

public/images/icons-vector.svg: $(SRC_IMG)/icon.svg
	cp $< $@

public/wasm_exec.js:
	GOROOT=$$(go env GOROOT); \
	if [ -f "$$GOROOT/lib/wasm/wasm_exec.js" ]; then \
		cp "$$GOROOT/lib/wasm/wasm_exec.js" public/wasm_exec.js; \
	elif [ -f "$$GOROOT/misc/wasm/wasm_exec.js" ]; then \
		cp "$$GOROOT/misc/wasm/wasm_exec.js" public/wasm_exec.js; \
	fi

public/tgo.wasm: wasm/tgo.go wasm/main.go
	GOOS=js GOARCH=wasm go build -trimpath -ldflags="-s -w" -o public/tgo.wasm ./wasm

public/sw.js: $(ALL_ICONS) public/index.html public/tgo.wasm public/wasm_exec.js
	npm run build