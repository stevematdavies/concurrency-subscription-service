BINARY_NAME=subscriptionapp
DSN="host=localhost port=5432 user=postgres password=postgres dbname=concurrency sslmode=disable"
REDIS="127.0.0.6379"

## build: builds the binary
build:
	@echo "Building..."
	env CGO_ENABLED=0 go build -ldflags="-s -w" -o ${BINARY_NAME} ./cmd/web
	@echo "Built!"

## run: buillds and runs the application
run: build
	@echo "Starting..."
	@env DSN=${DSN} REDIS=${REDIS} ./${BINARY_NAME} &
	@echo "Started!"

## clean: runs go clean and delteles the binaries
clean:
	@echo "Cleaninng..."
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

## restart: restarts the app with stop & start
restart: stop start

## test: runs tests
test:
	@echo "Running tests..."
	@go test -v ./..