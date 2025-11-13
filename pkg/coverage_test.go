package coverage
import (
"os"
"testing"
)
func TestParseCoverageFile(t *testing.T) {
content := []byte(`mode: atomic
shipping-management/internal/application/usecases/create_package.go:17.35,22.2 1 2
shipping-management/internal/application/usecases/get_package.go:13.97,17.2 1 0
`)
tmpfile, err := os.CreateTemp("", "coverage-*.out")
if err != nil {
t.Fatal(err)
}
defer os.Remove(tmpfile.Name())
if _, err := tmpfile.Write(content); err != nil {
t.Fatal(err)
}
if err := tmpfile.Close(); err != nil {
t.Fatal(err)
}
report, err := ParseCoverageFile(tmpfile.Name())
if err != nil {
t.Fatalf("Failed to parse coverage file: %v", err)
}
if report.Mode != "atomic" {
t.Errorf("Expected mode 'atomic', got '%s'", report.Mode)
}
if len(report.Files) != 2 {
t.Errorf("Expected 2 files, got %d", len(report.Files))
}
}
func TestGetCoverageStats(t *testing.T) {
coverage := &FileCoverage{
FileName: "test.go",
Blocks: []CoverageBlock{
{StartLine: 1, EndLine: 5, NumStmt: 3, Count: 2},
{StartLine: 6, EndLine: 10, NumStmt: 2, Count: 0},
{StartLine: 11, EndLine: 15, NumStmt: 5, Count: 1},
},
}
total, covered, pct := coverage.GetCoverageStats()
if total != 10 {
t.Errorf("Expected total 10, got %d", total)
}
if covered != 8 {
t.Errorf("Expected covered 8, got %d", covered)
}
expectedPct := 80.0
if pct != expectedPct {
t.Errorf("Expected percentage %.1f, got %.1f", expectedPct, pct)
}
}
func TestGetCoverageColor(t *testing.T) {
tests := []struct {
pct   float64
color string
}{
{90, "#4caf50"},
{70, "#8bc34a"},
{50, "#ff9800"},
{30, "#ff5722"},
{10, "#f44336"},
}
for _, tt := range tests {
color := GetCoverageColor(tt.pct)
if color != tt.color {
t.Errorf("For %.1f%%, expected color %s, got %s", tt.pct, tt.color, color)
}
}
}
func TestGenerateHTMLReport(t *testing.T) {
report := &CoverageReport{
Mode: "atomic",
Files: map[string]*FileCoverage{
"test.go": {
FileName: "test.go",
Blocks: []CoverageBlock{
{StartLine: 1, EndLine: 5, NumStmt: 3, Count: 2},
},
},
},
}
tmpfile, err := os.CreateTemp("", "coverage-*.html")
if err != nil {
t.Fatal(err)
}
defer os.Remove(tmpfile.Name())
tmpfile.Close()
err = GenerateHTMLReport(report, tmpfile.Name())
if err != nil {
t.Fatalf("Failed to generate HTML report: %v", err)
}
info, err := os.Stat(tmpfile.Name())
if err != nil {
t.Fatalf("Failed to stat output file: %v", err)
}
if info.Size() == 0 {
t.Error("Expected non-empty HTML file")
}
}
