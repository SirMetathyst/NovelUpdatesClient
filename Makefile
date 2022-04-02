.PHONY: all build generate tags genre
all: run

run:
	go run ./cmd/nuc/main.go

./cmd/nug/nug:
	go build -o ./cmd/nug/nug ./cmd/nug/main.go

generate: ./cmd/nug tags genre

tags: ./cmd/nug/nug
	./cmd/nug/nug -type=tags | gofmt > ./tags_generated.go

genre: ./cmd/nug/nug
	./cmd/nug/nug -type=genres | gofmt > ./genres_generated.go