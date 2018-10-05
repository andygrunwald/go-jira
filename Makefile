test:
	go test -v ./...

test_cover: ## Run application unit tests with coverage and generate global code coverage report
	@go test ./... -parallel 4 -failfast -cover -coverprofile=coverage.out

covercli: ## Generate code coverage report
	@go tool cover -func=coverage.out

coverhtml: ## Generate global code coverage report in HTML
	@go tool cover -html=coverage.out

coverage: test_cover coverhtml

coverage_cli: test_cover covercli