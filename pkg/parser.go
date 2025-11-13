package coverage
import (
"bufio"
"fmt"
"os"
"strconv"
"strings"
)
type CoverageBlock struct {
StartLine int
StartCol  int
EndLine   int
EndCol    int
NumStmt   int
Count     int
}
type FileCoverage struct {
FileName string
Blocks   []CoverageBlock
}
type CoverageReport struct {
Mode  string
Files map[string]*FileCoverage
}
func ParseCoverageFile(filename string) (*CoverageReport, error) {
file, err := os.Open(filename)
if err != nil {
return nil, fmt.Errorf("failed to open coverage file: %w", err)
}
defer file.Close()
report := &CoverageReport{
Files: make(map[string]*FileCoverage),
}
scanner := bufio.NewScanner(file)
lineNum := 0
for scanner.Scan() {
line := scanner.Text()
lineNum++
if lineNum == 1 {
if strings.HasPrefix(line, "mode:") {
report.Mode = strings.TrimSpace(strings.TrimPrefix(line, "mode:"))
}
continue
}
parts := strings.Fields(line)
if len(parts) != 3 {
continue
}
fileAndLines := strings.SplitN(parts[0], ":", 2)
if len(fileAndLines) != 2 {
continue
}
fileName := fileAndLines[0]
lineInfo := fileAndLines[1]
positions := strings.Split(lineInfo, ",")
if len(positions) != 2 {
continue
}
startParts := strings.Split(positions[0], ".")
endParts := strings.Split(positions[1], ".")
if len(startParts) != 2 || len(endParts) != 2 {
continue
}
startLine, _ := strconv.Atoi(startParts[0])
startCol, _ := strconv.Atoi(startParts[1])
endLine, _ := strconv.Atoi(endParts[0])
endCol, _ := strconv.Atoi(endParts[1])
numStmt, _ := strconv.Atoi(parts[1])
count, _ := strconv.Atoi(parts[2])
block := CoverageBlock{
StartLine: startLine,
StartCol:  startCol,
EndLine:   endLine,
EndCol:    endCol,
NumStmt:   numStmt,
Count:     count,
}
if _, exists := report.Files[fileName]; !exists {
report.Files[fileName] = &FileCoverage{
FileName: fileName,
Blocks:   []CoverageBlock{},
}
}
report.Files[fileName].Blocks = append(report.Files[fileName].Blocks, block)
}
if err := scanner.Err(); err != nil {
return nil, fmt.Errorf("error reading coverage file: %w", err)
}
return report, nil
}
func (fc *FileCoverage) GetCoverageStats() (totalStmts, coveredStmts int, percentage float64) {
for _, block := range fc.Blocks {
totalStmts += block.NumStmt
if block.Count > 0 {
coveredStmts += block.NumStmt
}
}
if totalStmts > 0 {
percentage = float64(coveredStmts) / float64(totalStmts) * 100
}
return
}
func (r *CoverageReport) GetOverallStats() (totalStmts, coveredStmts int, percentage float64) {
for _, file := range r.Files {
total, covered, _ := file.GetCoverageStats()
totalStmts += total
coveredStmts += covered
}
if totalStmts > 0 {
percentage = float64(coveredStmts) / float64(totalStmts) * 100
}
return
}
