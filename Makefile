GO ?= go
GOPATH := $(CURDIR)/_vendor:$(GOPATH)

all: docker

clean:
	rm spellchecker
	rm -f static/jquery.spellchecker.css
	rm -f static/bundle.js
	rm -rf node_modules

docker: spellchecker static/bundle.js static/jquery.spellchecker.css
	docker build .

spellchecker: rest.go
	$(GO) build

node_modules: package.json
	npm install

static/jquery.spellchecker.css: node_modules
	cp node_modules/jquery-spellchecker/src/css/jquery.spellchecker.css $@

static/bundle.js: spellcheck.js node_modules
	browserify spellcheck.js > $@
