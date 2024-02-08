build:
	@templ generate 
	@go build -C ./cmd/app -o ../../tmp/main

generate:
	@templ generate
