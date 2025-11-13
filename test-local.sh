#!/bin/bash

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

echo -e "${GREEN}════════════════════════════════════════════════════════════════${NC}"
echo -e "${GREEN}   Go Coverage HTML Reporter - Local Testing Script${NC}"
echo -e "${GREEN}════════════════════════════════════════════════════════════════${NC}"
echo

# Create test directory
TEST_DIR=~/go-coverage-test-$(date +%s)
echo -e "${YELLOW}📁 Creating test directory: $TEST_DIR${NC}"
mkdir -p $TEST_DIR
cd $TEST_DIR

# Create Go module
echo -e "${YELLOW}📦 Creating Go module...${NC}"
cat > go.mod << 'EOF'
module example.com/coveragetest

go 1.21
EOF

# Create calculator.go
echo -e "${YELLOW}📝 Creating calculator.go...${NC}"
cat > calculator.go << 'EOF'
package main

// Add returns the sum of two integers
func Add(a, b int) int {
    return a + b
}

// Subtract returns the difference of two integers
func Subtract(a, b int) int {
    return a - b
}

// Multiply returns the product of two integers
func Multiply(a, b int) int {
    return a * b
}

// Divide returns the quotient of two integers
func Divide(a, b int) int {
    if b == 0 {
        return 0
    }
    return a / b
}

// IsEven checks if a number is even
func IsEven(n int) bool {
    return n%2 == 0
}

// Max returns the larger of two integers
func Max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
EOF

# Create test file (only testing some functions)
echo -e "${YELLOW}🧪 Creating calculator_test.go...${NC}"
cat > calculator_test.go << 'EOF'
package main

import "testing"

func TestAdd(t *testing.T) {
    tests := []struct {
        a, b, want int
    }{
        {2, 3, 5},
        {0, 0, 0},
        {-1, 1, 0},
    }

    for _, tt := range tests {
        if got := Add(tt.a, tt.b); got != tt.want {
            t.Errorf("Add(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.want)
        }
    }
}

func TestSubtract(t *testing.T) {
    if got := Subtract(5, 3); got != 2 {
        t.Errorf("Subtract(5, 3) = %d, want 2", got)
    }
}

func TestIsEven(t *testing.T) {
    if !IsEven(4) {
        t.Error("IsEven(4) should be true")
    }
    if IsEven(3) {
        t.Error("IsEven(3) should be false")
    }
}

// Note: Multiply, Divide, and Max are NOT tested
// This will demonstrate the coverage visualization
EOF

# Create main.go for completeness
echo -e "${YELLOW}📝 Creating main.go...${NC}"
cat > main.go << 'EOF'
package main

import "fmt"

func main() {
    fmt.Println("Calculator Demo")
    fmt.Printf("2 + 3 = %d\n", Add(2, 3))
    fmt.Printf("5 - 3 = %d\n", Subtract(5, 3))
    fmt.Printf("4 * 5 = %d\n", Multiply(4, 5))
    fmt.Printf("10 / 2 = %d\n", Divide(10, 2))
}
EOF

# Run tests with coverage
echo
echo -e "${YELLOW}🧪 Running tests with coverage...${NC}"
go test -coverprofile=coverage.out -v

# Display coverage summary
echo
echo -e "${YELLOW}📊 Coverage summary:${NC}"
go tool cover -func=coverage.out

# Get path to go-coverage binary
COVERAGE_BIN="$HOME/projects/go-coverage/bin/go-coverage"
if [ ! -f "$COVERAGE_BIN" ]; then
    echo -e "${RED}❌ Error: go-coverage binary not found at $COVERAGE_BIN${NC}"
    echo -e "${YELLOW}💡 Please build it first:${NC}"
    echo "   cd ~/projects/go-coverage"
    echo "   make build"
    exit 1
fi

# Generate HTML report
echo
echo -e "${YELLOW}🔨 Generating HTML report...${NC}"
$COVERAGE_BIN

if [ $? -eq 0 ]; then
    echo -e "${GREEN}✅ Report generated successfully!${NC}"
    echo
    echo -e "${GREEN}════════════════════════════════════════════════════════════════${NC}"
    echo -e "${GREEN}   Test Complete!${NC}"
    echo -e "${GREEN}════════════════════════════════════════════════════════════════${NC}"
    echo
    echo -e "${YELLOW}📂 Test directory: $TEST_DIR${NC}"
    echo -e "${YELLOW}📄 HTML report: $TEST_DIR/coverage.html${NC}"
    echo
    echo -e "${YELLOW}What you'll see in the report:${NC}"
    echo -e "  ${GREEN}✅ Green (Covered):${NC} Add, Subtract, IsEven"
    echo -e "  ${RED}❌ Red (Not Covered):${NC} Multiply, Divide, Max"
    echo
    echo -e "${YELLOW}🌐 Opening report in browser...${NC}"

    # Try to open the HTML file
    if command -v xdg-open &> /dev/null; then
        xdg-open coverage.html &
    elif command -v open &> /dev/null; then
        open coverage.html &
    elif command -v start &> /dev/null; then
        start coverage.html &
    else
        echo -e "${YELLOW}💡 Please open this file manually:${NC}"
        echo "   $TEST_DIR/coverage.html"
    fi

    echo
    echo -e "${YELLOW}📌 To clean up later:${NC}"
    echo "   rm -rf $TEST_DIR"
else
    echo -e "${RED}❌ Failed to generate report${NC}"
    exit 1
fi

