# Serve Go web server
serve:
	go run .

# Run all integration tests
test_integration:
	go test -v ./tests/integration

# Run integration tests for a spesific test case
test_integration_custom:
	go test -v ./tests/integration -run="$(CASE)"