BINARY_NAME=adamz-api
ARCH=amd64
TARGET_DIR=./target
.DEFAULT_GOAL := run

build:
	GOARCH=${ARCH} GOOS=darwin go build -o ${TARGET_DIR}/${BINARY_NAME}-darwin-${ARCH} main.go
	GOARCH=${ARCH} GOOS=linux go build -o ${TARGET_DIR}/${BINARY_NAME}-linux-${ARCH} main.go
	GOARCH=${ARCH} GOOS=windows go build -o ${TARGET_DIR}/${BINARY_NAME}-windows-${ARCH} main.go

run: build
	${TARGET_DIR}/${BINARY_NAME}-linux-${ARCH}

clean:
	go clean
	rm ${TARGET_DIR}/${BINARY_NAME}-darwin-${ARCH}
	rm ${TARGET_DIR}/${BINARY_NAME}-linux-${ARCH}
	rm ${TARGET_DIR}/${BINARY_NAME}-windows-${ARCH}

test:
	go test ./...

test_coverage:
	go test ./... -coverprofile=coverage.out

dep:
	go mod download

tidy:
	go mod tidy

vet:
	go vet

lint:
	golangci-lint run --enable-all

