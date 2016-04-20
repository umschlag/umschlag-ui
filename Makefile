DIST := dist
SRC := src
WEBPACK := node_modules/.bin/webpack
ESLINT := node_modules/.bin/eslint

all: deps build

deps:
	npm install

build: dir
	$(WEBPACK) --optimize-minimize

watch: dir
	$(WEBPACK) -w

dir:
	mkdir -p $(DIST)

lint:
	$(ESLINT) $(SRC)/

clean:
	rm -rf $(DIST)

.PHONY: all deps build watch dir lint clean
