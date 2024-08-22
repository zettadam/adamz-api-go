BINARY_NAME=adamz-api
ARCH=amd64
TARGET_DIR=./target
MAIN_FILE=./main.go
.DEFAULT_GOAL := run

build:
	GOARCH=${ARCH} GOOS=linux go build -o ${TARGET_DIR}/${BINARY_NAME}-linux-${ARCH} ${MAIN_FILE}

run: build
	${TARGET_DIR}/${BINARY_NAME}-linux-${ARCH}

clean:
	go clean
	rm -rf ${TARGET_DIR}

test:
	go test ./...

test_coverage:
	go test ./... -coverprofile=coverage.out

dep:
	go mod download

tidy:
	go mod tidy

