.PHONY: icons

SHELL := /bin/bash

SRC_IMG = ./raw
BUILD = ./public
BUILD_IMAGES = $(BUILD)/images
resolutions := 16 32 180 192 512
ALL_ICONS := $(foreach resolution, $(resolutions), $(BUILD_IMAGES)/icons-$(resolution).png)

all: icons public/tgo.js public/sw.js

icons: public/images/icons-vector.svg $(ALL_ICONS)

$(BUILD_IMAGES)/icons-%.png: $(SRC_IMG)/icon.svg
	inkscape $(SRC_IMG)/icon.svg -w $* -h $* --export-png=$@

public/images/icons-vector.svg: $(SRC_IMG)/icon.svg
	cp $< $@

public/tgo.js: gopherjs/tgo.go
	gopherjs build -m -o public/tgo.js gopherjs/tgo.go

public/sw.js: $(ALL_ICONS) public/index.html public/tgo.js
	npm run build