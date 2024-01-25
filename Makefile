BINARY_NAME=mnrnaters

gorun:
	@go build -o bin/${BINARY_NAME} ./cmd

run: gorun
	clear
	@./bin/${BINARY_NAME}
	@echo "Backend running..."

build_cross: gorun
	GOOS=windows GOARCH=amd64 go build -o bin/${BINARY_NAME}.exe ./cmd
	GOOS=darwin GOARCH=amd64 go build -o bin/${BINARY_NAME} ./cmd
	GOOS=linux GOARCH=amd64 go build -o bin/${BINARY_NAME}-linux ./cmd

clean:
	go clean
	rm -rf bin/*

install:
	go mod tidy

stop:
	@-pkill -f ${BINARY_NAME}