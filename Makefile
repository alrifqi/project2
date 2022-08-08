BINARY_NAME=app

build:
	GOARCH=amd64 GOOS=darwin go build -o ./dist/${BINARY_NAME} app.go

run:
	@./dist/app

build_and_run: build run

clean:
	go clean
	rm ${BINARY_NAME} -darwin

run-dev:
	air

lint:
	golangci-lint run --enable-all

run-tests:
	go test -v -tags dynamic `go list ./... | grep -i 'controller\|repository\|usecase'` -coverprofile=coverage.out && go tool cover -html=coverage.out -o coverage.html