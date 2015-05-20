GO ?= go
GOPATH := $(CURDIR)/_vendor:$(GOPATH)

all: spellchecker

docker: spellchecker static/bundle.js static/jquery.spellchecker.css
	docker build .

spellchecker:
	$(GO) build

static/jquery.spellchecker.css:
	cp node_modules/jquery-spellchecker/src/css/jquery.spellchecker.css $@

static/bundle.js: spellcheck.js
	browserify spellcheck.js > $@
