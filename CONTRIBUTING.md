# Contributing to Go Coverage HTML Reporter

Thank you for your interest in contributing! This document provides guidelines for contributing to this project.

## Getting Started

1. Fork the repository
2. Clone your fork: `git clone https://github.com/your-username/go-coverage.git`
3. Create a new branch: `git checkout -b feature/your-feature-name`
4. Make your changes
5. Run tests: `go test ./...`
6. Commit your changes: `git commit -am 'Add some feature'`
7. Push to the branch: `git push origin feature/your-feature-name`
8. Submit a pull request

## Development Setup

### Prerequisites

- Go 1.21 or higher
- Git

### Building

```bash
# Build the CLI tool
go build -o go-coverage ./cmd/go-coverage

# Run tests
go test ./...

# Run tests with coverage
go test -coverprofile=coverage.out ./...
./go-coverage
```

## Code Style

- Follow standard Go conventions and idioms
- Use `gofmt` to format your code
- Write clear, descriptive commit messages
- Add tests for new functionality
- Update documentation as needed

## Testing

- Write unit tests for new functionality
- Ensure all tests pass before submitting a PR
- Aim for high test coverage

## Pull Request Process

1. Update the README.md with details of changes if applicable
2. Update the documentation with any new features or changes
3. Ensure all tests pass
4. Your PR will be reviewed by maintainers
5. Address any feedback from reviewers
6. Once approved, your PR will be merged

## Reporting Bugs

When reporting bugs, please include:

- Go version
- Operating system
- Steps to reproduce
- Expected behavior
- Actual behavior
- Any relevant error messages or logs

## Feature Requests

We welcome feature requests! Please open an issue describing:

- The problem you're trying to solve
- Your proposed solution
- Any alternatives you've considered
- Additional context

## Code of Conduct

- Be respectful and inclusive
- Welcome newcomers
- Focus on constructive feedback
- Maintain a positive environment

## Questions?

Feel free to open an issue for any questions about contributing!

Thank you for contributing! 🎉

