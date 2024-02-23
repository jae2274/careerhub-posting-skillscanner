BINARY_NAME=myapp
CODE_DIR=./careerhub/scanner

include test.env

## build: Build binary
build:
	@echo "Building..."
	@go build -ldflags="-s -w" -o ${BINARY_NAME} ${CODE_DIR}
	@echo "Built!"

## run: builds and runs the application
run: build
	@echo "Starting..."
	@env GRPC_ENDPOINT=${GRPC_ENDPOINT} POST_LOG_URL=${POST_LOG_URL} ./${BINARY_NAME} 
	@echo "Started!"

## clean: runs go clean and deletes binaries
clean:
	@echo "Cleaning..."
	@go clean
	@rm ${BINARY_NAME}
	@echo "Cleaned!"

## start: an alias to run
start: run

## stop: stops the running application
stop:
	@echo "Stopping..."
	@-pkill -SIGTERM -f "./${BINARY_NAME}"
	@echo "Stopped!"

## restart: stops and starts the application
restart: stop start

proto:
	@export PATH="$PATH:$(go env GOPATH)/bin"
	@protoc careerhub/scanner/scanner_grpc/*.proto  --go_out=. --go-grpc_out=. --go-grpc_opt=paths=source_relative --go_opt=paths=source_relative  --proto_path=.

## test: runs all tests
test:	
	@echo "Testing..."
	@env GRPC_ENDPOINT=${GRPC_ENDPOINT} POST_LOG_URL=${POST_LOG_URL} go test -p 1 -timeout 300s ./test/...
	

