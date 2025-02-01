.PHONY: test test-all test-one

# Run all tests
test-all:
	go test ./...

# Run specific test with pattern
# Usage: make test p=Problem1
test:
	go test ./problems_test -ginkgo.focus="$(p)"

# Run test by problem number
# Usage: make test-one n=1
test-one:
	go test ./problems_test -ginkgo.focus="Problem$(n)"
