.PHONY: all build generate tags genre
all: run

run:
	go run ./cmd/nuc/main.go

generate: tags genre

tags:
	go run ./cmd/tagc/main.go | gofmt > ./tags.go

genre:
	go run ./cmd/genrec/main.go | gofmt > ./genres.go