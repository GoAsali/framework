build:
	mkdir -p bin/languages
	cp -r ./languages ./bin/languages
	go build -o bin/main main.go

run:
	go run main.go