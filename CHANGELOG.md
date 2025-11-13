# Release Notes

## v0.1.2 (2024-11-13)

### Changes
- ✅ Updated installation documentation
- ✅ Added comprehensive INSTALLATION.md guide
- ✅ Added RELEASE_v0.1.1.md documentation
- ✅ Improved README with clearer installation instructions

### Installation

```bash
go install github.com/rayque/go-coverage/cmd/go-coverage@v0.1.2
```

Or use latest:

```bash
go install github.com/rayque/go-coverage/cmd/go-coverage@latest
```

---

## v0.1.1 (2024-11-13)

### Changes
- ✅ Fixed missing cmd/go-coverage/main.go in repository
- ✅ Fixed go.mod Go version (1.21)
- ✅ Fixed import formatting in main.go
- ✅ First working installable release

### Installation

```bash
go install github.com/rayque/go-coverage/cmd/go-coverage@v0.1.1
```

---

## Initial Release

First public release of go-coverage HTML reporter.

### Features

- 📊 Parse Go coverage files
- 🎨 Generate beautiful HTML reports with syntax highlighting
- 🗂️ GitHub-like file tree navigation
- 🎯 Visual coverage highlighting (covered vs uncovered lines)
- 📈 Coverage statistics per file and overall
- 🚀 Easy to use CLI tool
- 📦 Library for programmatic usage
- ✅ Zero dependencies
- ✅ Self-contained HTML files

### Components

- **CLI Tool**: `cmd/go-coverage/main.go`
- **Core Library**: `pkg/parser.go`, `pkg/source.go`, `pkg/html.go`, `pkg/template.go`
- **Tests**: `pkg/coverage_test.go`
- **Documentation**: README.md, USAGE.md, QUICKSTART.md, CONTRIBUTING.md

### Usage

```bash
# Generate coverage
go test -coverprofile=coverage.out ./...

# Generate HTML report
go-coverage

# Open report
xdg-open coverage.html
```

