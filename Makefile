build:
	@echo "Building api"
	go build -o app cmd/server/main.go

test:
	@echo "Running tests"
	go test -v ./...

lint:
	@echo "Running Lint"
	golangci-lint run

local-run:
	@echo "Running docker images instantly!"
	docker-compose up --build

run: build
	@echo "Running docker images on the background"
	docker-compose up --build -d

start: run

stop:
	@echo "Stopped the images"
	docker-compose stop

test-fields:
	curl --location --request GET 'http://localhost:8080/api/v1/fields'