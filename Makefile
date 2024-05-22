
.PHONY : tools mock docs

all: example

mod:
	go mod tidy

# maxos brew install FiloSottile/musl-cross/musl-cross
example: mod
	mkdir -p build
	go build -o example/example ./example/*.go
