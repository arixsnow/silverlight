# Testing Guide

This document describes how to run tests for the Silverlight project.

## Running Tests

### Local Testing

You can run tests using the following commands:

```bash
# Run all tests in the project
go test ./...

# Run tests with verbose output
go test -v ./...

# Run tests for a specific package
go test ./pkg/openflow10/...

# Run a specific test
go test -v ./pkg/openflow10/... -run TestDecodePortName

# Run tests with coverage
go test -cover ./...

# Generate coverage report
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out  # View coverage in browser
```

### Test Structure

Tests are organized following Go's standard testing conventions:

- Test files are named with `_test.go` suffix
- Test functions are prefixed with `Test`
- Table-driven tests are used for comprehensive test cases
- Each package has its own tests in the same directory as the source code

## Continuous Integration

Tests are automatically run on every push and pull request using GitHub Actions.

The workflow:

- Runs on Ubuntu latest
- Tests with multiple Go versions
- Reports test coverage
- Ensures code formatting follows Go standards

You can view the test results in the GitHub Actions tab of the repository.

## Writing Tests

When adding new features, please ensure:

1. Write tests for new functionality
2. Follow table-driven test patterns when applicable
3. Test both success and error cases
4. Maintain test coverage above 80%

## Test Dependencies

All test dependencies are managed through `go.mod`.
o additional setup is required to run the tests.