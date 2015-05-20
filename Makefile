GO ?= go
GOPATH := $(CURDIR)/_vendor:$(GOPATH)

all: spellchecker

docker: spellchecker
	docker build .

spellchecker:
	$(GO) build
