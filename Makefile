.PHONY: build run clean generate genres tags

all: run clean

build:
	go build -o ./cmd/nuc/nuc ./cmd/nuc/main.go

run: build
	./cmd/nuc/nuc

./cmd/nug/nug:
	go build -o ./cmd/nug/nug ./cmd/nug/main.go

generate: ./cmd/nug/nug
	./cmd/nug/nug -type=all | gofmt > ./types_generated.go

clean:
	$(RM) ./cmd/nuc/nuc
	$(RM) ./cmd/nug/nug