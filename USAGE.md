# Usage Guide
## Quick Start
### Generate Coverage Report
1. Run your tests with coverage:
```bash
go test -coverprofile=coverage.out ./...
```
2. Generate HTML report:
```bash
go-coverage
```
3. Open `coverage.html` in your browser
## Command Line Options
```bash
go-coverage [options]
```
### Options:
- `-input=<file>` - Path to coverage file (default: "coverage.out")
- `-output=<file>` - Path to HTML output (default: "coverage.html")
- `-version` - Show version information
- `-quiet` - Suppress output messages
### Examples:
```bash
# Use default files
go-coverage
# Custom input and output
go-coverage -input=my-coverage.out -output=report.html
# Quiet mode
go-coverage -quiet
# Show version
go-coverage -version
```
## Using as a Library
```go
package main
import (
    "log"
    coverage "github.com/rayque/go-coverage/pkg"
)
func main() {
    // Parse coverage file
    report, err := coverage.ParseCoverageFile("coverage.out")
    if err != nil {
        log.Fatal(err)
    }
    // Get overall statistics
    total, covered, pct := report.GetOverallStats()
    log.Printf("Coverage: %.1f%% (%d/%d)", pct, covered, total)
    // Get file-specific statistics
    for path, fileCov := range report.Files {
        total, covered, pct := fileCov.GetCoverageStats()
        log.Printf("%s: %.1f%%", path, pct)
    }
    // Generate HTML report
    err = coverage.GenerateHTMLReport(report, "coverage.html")
    if err != nil {
        log.Fatal(err)
    }
}
```
## HTML Report Features
The generated HTML report includes:
- **Overall Coverage**: Summary statistics at the top
- **File Tree**: Navigate files easily in the sidebar
- **Coverage Summary Table**: Quick overview of all files
- **Detailed View**: Line-by-line coverage with:
  - Green highlighting for covered lines
  - Red highlighting for uncovered lines
  - Line numbers for easy reference
- **Interactive Navigation**: Click on files to jump to their details
- **Color-coded Badges**:
  - Green (≥80%): Excellent coverage
  - Light Green (≥60%): Good coverage
  - Orange (≥40%): Fair coverage
  - Deep Orange (≥20%): Poor coverage
  - Red (<20%): Very poor coverage
## Building from Source
```bash
# Clone the repository
git clone https://github.com/rayque/go-coverage.git
cd go-coverage
# Build the binary
go build -o bin/go-coverage ./cmd/go-coverage
# Or use make
make build
# Install globally
make install
```
## Running Tests
```bash
# Run all tests
go test ./...
# Run tests with coverage
go test -coverprofile=coverage.out ./...
# Generate coverage report for the library itself
go-coverage
# Or use make
make coverage-html
```
## Integration with CI/CD
### GitHub Actions
```yaml
name: Coverage
on: [push, pull_request]
jobs:
  coverage:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
      - name: Set up Go
        uses: actions/setup-go@v4
        with:
          go-version: '1.21'
      - name: Install go-coverage
        run: go install github.com/rayque/go-coverage/cmd/go-coverage@latest
      - name: Run tests with coverage
        run: go test -coverprofile=coverage.out ./...
      - name: Generate HTML report
        run: go-coverage
      - name: Upload coverage report
        uses: actions/upload-artifact@v3
        with:
          name: coverage-report
          path: coverage.html
```
### GitLab CI
```yaml
coverage:
  stage: test
  script:
    - go test -coverprofile=coverage.out ./...
    - go install github.com/rayque/go-coverage/cmd/go-coverage@latest
    - go-coverage
  artifacts:
    paths:
      - coverage.html
    expire_in: 30 days
```
## Troubleshooting
### "Coverage file does not exist"
Make sure you run `go test -coverprofile=coverage.out` first to generate the coverage file.
### "Source file not found"
The tool tries to read source files from the current directory. Make sure you run it from your project root, or the source files may not be found (coverage statistics will still be shown).
### Empty or incorrect coverage
Ensure your coverage file is in the correct format. It should start with `mode:` and contain coverage blocks.
## Tips
1. **Regular Monitoring**: Generate coverage reports regularly to track your project's test coverage over time
2. **Coverage Goals**: Aim for at least 80% coverage for critical code paths
3. **Ignore Generated Code**: Add `//go:generate` directives to exclude generated files from coverage
4. **Focus on Quality**: High coverage doesn't always mean good tests - focus on meaningful test cases
