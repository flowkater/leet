# LeetCode Algorithm Solutions in Go

This repository contains my solutions to LeetCode problems implemented in Go.

## Project Structure

- `problems/`: Contains the solution implementations
- `problems_test/`: Contains the test files using Ginkgo framework

## Testing

This project uses [Ginkgo](https://onsi.github.io/ginkgo/) for testing.

### Test Commands

Run all tests:
```bash
make test-all
```

Run test for a specific problem by number:
```bash
make test-one n=1  # Run tests for Problem 1 (Two Sum)
make test-one n=20 # Run tests for Problem 20 (Valid Parentheses)
```

Run test by pattern:
```bash
make test p=Problem1  # Run tests matching "Problem1"
```

## Problems

- [Blind 75 LeetCode Questions](https://leetcode.com/discuss/general-discussion/460599/blind-75-leetcode-questions)
- [Order Grind 75 Questions](https://www.techinterviewhandbook.org/grind75/)

1. Two Sum - [Problem Link](https://leetcode.com/problems/two-sum/)
2. Valid Parentheses - [Problem Link](https://leetcode.com/problems/valid-parentheses/)
