# Project Structure
```
go-coverage/
├── bin/
│   └── go-coverage              # Compiled binary (ready to use!)
│
├── cmd/
│   └── go-coverage/
│       └── main.go              # CLI application entry point
│
├── pkg/
│   ├── coverage_test.go         # Unit tests (all passing!)
│   ├── html.go                  # HTML report generator
│   ├── parser.go                # Coverage file parser
│   ├── source.go                # Source code handler & file tree builder
│   └── template.go              # HTML template with CSS/JavaScript
│
├── .gitignore                   # Git ignore rules
├── CONTRIBUTING.md              # Contribution guidelines
├── go.mod                       # Go module definition
├── LICENSE                      # MIT License
├── Makefile                     # Build automation
├── PROJECT_STRUCTURE.md         # This file
├── QUICKSTART.md                # Quick start guide
├── README.md                    # Main documentation
├── test-local.sh                # Automated local testing script
├── USAGE.md                     # Detailed usage guide
│
├── coverage.out                 # Example coverage file (for testing)
└── coverage.html                # Example HTML report (generated)
```
## File Descriptions
### Core Library (`pkg/`)
- **parser.go** - Parses Go coverage files
  - `ParseCoverageFile()` - Main parsing function
  - `CoverageReport` - Data structure for coverage data
  - `GetCoverageStats()` - Calculate coverage percentages
- **source.go** - Source code handling
  - `GetFileWithSource()` - Read and annotate source files
  - `BuildFileTree()` - Create hierarchical file structure
  - `GetCoverageColor()` - Color coding for coverage levels
- **html.go** - HTML generation
  - `GenerateHTMLReport()` - Main HTML generation function
  - `HTMLReport` - Report generator struct
- **template.go** - HTML template
  - Contains the complete HTML/CSS/JavaScript template
  - Self-contained, no external dependencies
- **coverage_test.go** - Unit tests
  - Tests for all major functions
  - 100% passing
### CLI Tool (`cmd/go-coverage/`)
- **main.go** - Command-line interface
  - Flag parsing
  - User-friendly output
  - Error handling
### Documentation
- **README.md** - Complete project documentation
- **QUICKSTART.md** - Quick start guide for testing
- **USAGE.md** - Detailed usage examples and integration guides
- **CONTRIBUTING.md** - How to contribute to the project
### Build & Development
- **Makefile** - Build automation
  - `make build` - Build the binary
  - `make install` - Install globally
  - `make test` - Run tests
  - `make coverage-html` - Generate coverage for this project
- **go.mod** - Go module configuration
  - Module path: `github.com/rayque/go-coverage`
  - Go version: 1.21
### Testing
- **test-local.sh** - Automated testing script
  - Creates sample project
  - Generates coverage
  - Opens HTML report
## Key Features
✅ **Zero Dependencies** - Pure Go, no external packages
✅ **Self-Contained** - HTML files include all CSS/JS
✅ **Tested** - Comprehensive unit tests
✅ **Documented** - README, usage guide, examples
✅ **Production Ready** - Error handling, logging, validation
## Build Output
- **bin/go-coverage** - The compiled executable
- Built with `go build -o bin/go-coverage ./cmd/go-coverage`
- Can be installed globally with `make install`
## Size
- Total Go code: ~1000 lines
- Binary size: ~3-4 MB
- HTML template: Responsive, modern design
## Next Steps
1. Test locally: `./test-local.sh`
2. Read the quickstart: `cat QUICKSTART.md`
3. Publish to GitHub
4. Share with the community!
