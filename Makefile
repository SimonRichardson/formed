all: build

install:
	go get github.com/Masterminds/glide
	go get github.com/mjibson/esc
	glide install
	$(MAKE) clean all

dist/formed:
	go build -o dist/formed github.com/SimonRichardson/formed/cmd/formed

pkg/templates/static.go:
	esc -o="pkg/templates/static.go" -pkg="templates" views

.PHONY: build
build: pkg/templates/static.go dist/formed

.PHONY: clean
clean: FORCE
	rm -f pkg/templates/static.go
	rm -rf dist/formed

FORCE:
