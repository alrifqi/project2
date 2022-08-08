BINARY_NAME=app

build:
	GOARCH=amd64 GOOS=darwin go build -o ${BINARY_NAME} -darwin *.go


run:
	./${BINARY_NAME}

build_and_run: build run

clean:
	go clean
	rm ${BINARY_NAME} -darwin

run-dev:
	go run *.go

lint:
	golangci-lint run --enable-all