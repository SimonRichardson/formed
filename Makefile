all: dist/formed

install:
	go get github.com/Masterminds/glide
	glide install
	$(MAKE) clean all

dist/formed:
	go build -o dist/formed github.com/SimonRichardson/formed/cmd/formed

.PHONY: build
build: dist/formed

.PHONY: clean
clean: FORCE
	rm -rf dist/formed

FORCE:
