# Go Coverage HTML Reporter

A powerful Go library and CLI tool to generate beautiful, interactive HTML reports from Go coverage files.

## Features

- 📊 Parse Go coverage files (`coverage.out`)
- 🎨 Generate interactive HTML reports with syntax highlighting
- 🗂️ File tree navigation similar to GitHub
- 🎯 Visual coverage highlighting (covered vs uncovered lines)
- 📈 Coverage statistics per file and overall
- 🚀 Easy to use CLI tool
- 📦 Library for programmatic usage

## Installation

### Install from GitHub (once published)

```bash
go install github.com/rayque/go-coverage/cmd/go-coverage@latest
```

### Build from Source

```bash
git clone https://github.com/rayque/go-coverage.git
cd go-coverage
make build
# Binary will be in ./bin/go-coverage

# Or install globally
make install
```

## Quick Start

1. **Generate coverage data from your tests:**
   ```bash
   cd your-go-project
   go test -coverprofile=coverage.out ./...
   ```

2. **Generate HTML report:**
   ```bash
   go-coverage
   # Or specify files
   go-coverage -input=coverage.out -output=coverage.html
   ```

3. **Open the report:**
   ```bash
   # Linux/WSL
   xdg-open coverage.html
   
   # macOS
   open coverage.html
   
   # Windows
   start coverage.html
   ```

## CLI Usage

```bash
go-coverage [options]

Options:
  -input string
        Path to the coverage file (default "coverage.out")
  -output string
        Path to the output HTML file (default "coverage.html")
  -version
        Show version information
  -quiet
        Suppress output messages
```

### Examples

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

## Library Usage

You can also use this as a library in your Go programs:

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
    log.Printf("Overall Coverage: %.1f%% (%d/%d statements)", pct, covered, total)

    // Get statistics for individual files
    for path, fileCoverage := range report.Files {
        total, covered, pct := fileCoverage.GetCoverageStats()
        log.Printf("File: %s - Coverage: %.1f%% (%d/%d)", path, pct, covered, total)
    }

    // Generate HTML report
    err = coverage.GenerateHTMLReport(report, "coverage.html")
    if err != nil {
        log.Fatal(err)
    }
}
```

## Testing Locally

Here's a complete example of how to test this tool with your own Go project:

### Step 1: Create a Sample Go Project

```bash
# Create a new directory for testing
mkdir ~/test-coverage-demo
cd ~/test-coverage-demo

# Initialize a Go module
go mod init example.com/demo
```

### Step 2: Create Sample Go Files

Create `calculator.go`:
```go
package main

func Add(a, b int) int {
    return a + b
}

func Subtract(a, b int) int {
    return a - b
}

func Multiply(a, b int) int {
    return a * b
}

func Divide(a, b int) int {
    if b == 0 {
        return 0
    }
    return a / b
}
```

Create `calculator_test.go`:
```go
package main

import "testing"

func TestAdd(t *testing.T) {
    if Add(2, 3) != 5 {
        t.Error("Add failed")
    }
}

func TestSubtract(t *testing.T) {
    if Subtract(5, 3) != 2 {
        t.Error("Subtract failed")
    }
}

// Note: Multiply and Divide are not tested, so coverage won't be 100%
```

### Step 3: Generate Coverage Data

```bash
go test -coverprofile=coverage.out
```

### Step 4: Use go-coverage Tool

```bash
# If you built from source
/path/to/go-coverage/bin/go-coverage

# Or if installed globally
go-coverage
```

### Step 5: View the Report

```bash
# Open the generated HTML file
xdg-open coverage.html  # Linux
open coverage.html      # macOS
start coverage.html     # Windows
```

You should see:
- ✅ Green highlighting on tested functions (Add, Subtract)
- ❌ Red highlighting on untested functions (Multiply, Divide)
- Coverage percentage in the sidebar
- File navigation
- Overall statistics

### Complete Test Script

Here's a complete bash script to test everything:

```bash
#!/bin/bash

# Create test directory
TEST_DIR=~/go-coverage-test
mkdir -p $TEST_DIR
cd $TEST_DIR

# Create Go module
echo "Creating test module..."
cat > go.mod << 'EOF'
module example.com/coveragetest

go 1.21
EOF

# Create calculator.go
cat > calculator.go << 'EOF'
package main

func Add(a, b int) int {
    return a + b
}

func Subtract(a, b int) int {
    return a - b
}

func Multiply(a, b int) int {
    return a * b
}

func Divide(a, b int) int {
    if b == 0 {
        return 0
    }
    return a / b
}
EOF

# Create test file
cat > calculator_test.go << 'EOF'
package main

import "testing"

func TestAdd(t *testing.T) {
    if Add(2, 3) != 5 {
        t.Error("Add failed")
    }
}

func TestSubtract(t *testing.T) {
    if Subtract(5, 3) != 2 {
        t.Error("Subtract failed")
    }
}
EOF

# Run tests with coverage
echo "Running tests with coverage..."
go test -coverprofile=coverage.out

# Generate HTML report
echo "Generating HTML report..."
/path/to/go-coverage/bin/go-coverage

# Open report
echo "Opening report..."
xdg-open coverage.html 2>/dev/null || open coverage.html 2>/dev/null || start coverage.html

echo "Done! Check coverage.html in your browser."
```

Save this as `test-coverage.sh`, make it executable with `chmod +x test-coverage.sh`, and run it!

## HTML Report Features

The generated HTML report includes:

- **Overall Coverage Badge** - Color-coded based on percentage
  - 🟢 Green (≥80%): Excellent coverage
  - 🟡 Light Green (≥60%): Good coverage  
  - 🟠 Orange (≥40%): Fair coverage
  - 🟠 Deep Orange (≥20%): Poor coverage
  - 🔴 Red (<20%): Very poor coverage

- **Interactive Sidebar** - Click files to navigate
- **Coverage Summary Table** - Overview of all files
- **Line-by-line View** - See exactly what's covered
  - Green background = covered
  - Red background = not covered
  - Line numbers for reference

## Development

```bash
# Clone repository
git clone https://github.com/rayque/go-coverage.git
cd go-coverage

# Run tests
make test

# Build binary
make build

# Generate coverage for this project
go test -coverprofile=coverage.out ./...
./bin/go-coverage
```

## How It Works

1. **Parse** - Reads the `coverage.out` file generated by `go test -coverprofile`
2. **Analyze** - Extracts coverage data for each file and line
3. **Match** - Attempts to read source files from the current directory
4. **Generate** - Creates a self-contained HTML file with embedded CSS/JS
5. **Display** - Beautiful, interactive report you can share

## Requirements

- Go 1.21 or higher
- No external dependencies!

## Coverage File Format

This tool works with coverage files generated by Go's testing tools:

```bash
go test -coverprofile=coverage.out ./...
```

The coverage file format:
```
mode: atomic
path/to/file.go:startLine.startCol,endLine.endCol numStmt count
```

## License

MIT License - See LICENSE file for details

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

See [CONTRIBUTING.md](CONTRIBUTING.md) for details.

## Author

Created by Rayque Oliveira

## Links

- [GitHub Repository](https://github.com/rayque/go-coverage)
- [Issue Tracker](https://github.com/rayque/go-coverage/issues)
- [Documentation](https://github.com/rayque/go-coverage/blob/main/USAGE.md)

