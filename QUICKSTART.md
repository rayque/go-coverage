# Quick Start Guide - Testing go-coverage Locally

## Method 1: Run the Automated Test Script

The easiest way to test the tool:

```bash
cd ~/projects/go-coverage
./test-local.sh
```

This script will:
1. Create a test Go project with sample code
2. Run tests with coverage
3. Generate the HTML report
4. Open it in your browser

## Method 2: Manual Step-by-Step

### 1. Create a Test Project

```bash
mkdir ~/my-test-project
cd ~/my-test-project
go mod init example.com/mytest
```

### 2. Create a Go File

Create `math.go`:
```go
package main

func Add(a, b int) int {
    return a + b
}

func Multiply(a, b int) int {
    return a * b
}
```

### 3. Create a Test File

Create `math_test.go`:
```go
package main

import "testing"

func TestAdd(t *testing.T) {
    if Add(2, 3) != 5 {
        t.Fail()
    }
}

// Note: Multiply is NOT tested - it will show as uncovered!
```

### 4. Generate Coverage

```bash
go test -coverprofile=coverage.out
```

### 5. Generate HTML Report

```bash
~/projects/go-coverage/bin/go-coverage
```

### 6. View the Report

```bash
# Linux/WSL
xdg-open coverage.html

# macOS
open coverage.html

# Windows
start coverage.html
```

## What You'll See

- **Green lines**: Code that is covered by tests (Add function)
- **Red lines**: Code NOT covered by tests (Multiply function)
- **Statistics**: Coverage percentage for each file
- **Interactive navigation**: Click files in the sidebar

## Expected Output

When you run the coverage tool, you should see:

```
📊 Parsing coverage file: coverage.out
📈 Overall coverage: 50.0% (1/2 statements)
📁 Files analyzed: 1
🔨 Generating HTML report: coverage.html
✅ Report generated successfully!
🌐 Open coverage.html in your browser to view the report
```

## Tips

1. **Test with your own project**: Just run `go test -coverprofile=coverage.out ./...` in your project root
2. **100% coverage example**: Add `TestMultiply` to see all green
3. **Multiple files**: The sidebar will show all files with their coverage percentages

## Troubleshooting

### "Binary not found"
```bash
cd ~/projects/go-coverage
make build
```

### "Coverage file not found"
Make sure you ran:
```bash
go test -coverprofile=coverage.out
```

### "No source code shown"
The tool needs to find the source files in your current directory. Make sure you run it from your project root.

## Next Steps

Once you've tested locally:
1. Install globally: `cd ~/projects/go-coverage && make install`
2. Use in any project: Just run `go-coverage` 
3. Share your coverage reports with your team!

