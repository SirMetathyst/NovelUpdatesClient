.PHONY: build run clean generate genres tags

all: run clean

build:
	go build -o ./cmd/nuc/nuc ./cmd/nuc/main.go

run: build
	./cmd/nuc/nuc

clean:
	$(RM) ./cmd/nuc/nuc
	$(RM) ./cmd/nug/nug

generate: genres tags

./cmd/nug/nug:
	go build -o ./cmd/nug/nug ./cmd/nug/main.go

genres: ./cmd/nug/nug
	./cmd/nug/nug -type=genres | gofmt > ./genres_generated.go

tags: ./cmd/nug/nug
	./cmd/nug/nug -type=tags | gofmt > ./tags_generated.go