# Serve Go web server
run:
	go run .

# Run all integration tests
test_integration:
	go test -v ./tests/integration

# Run integration tests for a spesific test case
test_integration_custom:
	go test -v ./tests/integration -run="$(CASE)"

# Run all unit tests
test_unit:
	go test -v ./tests/unit

# Run integration tests for a spesific test case
test_unit_custom:
	go test -v ./tests/unit -run="$(CASE)"