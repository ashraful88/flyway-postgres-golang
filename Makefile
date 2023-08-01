# run server 
run:
	@echo "Running server..."
	@go run main.go
run-linter:
	@echo "Running linter..."
	@golangci-lint run
run-test:
	@echo "Running test..."
	@go test -v ./...
run-test-coverage:
	@echo "Running test verbose with coverage..."
	@go test -v -coverprofile=coverage.out ./...
run-docker:
	@echo "Running docker..."
	@docker-compose up -d
run-docker-down:
	@echo "Running docker down..."
	@docker-compose down
run-migrate:
	@echo "Running migrate..."
	@sh ./run_migration.sh