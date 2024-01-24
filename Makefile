gorun:
	go build -o bin/mnrnaters ./cmd

run: gorun
	./bin/mnrnaters

build_cross: gorun
	GOOS=windows GOARCH=amd64 go build -o bin/mnrnaters.exe ./cmd
	GOOS=darwin GOARCH=amd64 go build -o bin/mnrnaters ./cmd
	GOOS=linux GOARCH=amd64 go build -o bin/mnrnaters-linux ./cmd