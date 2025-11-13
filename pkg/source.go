package coverage
import (
"bufio"
"fmt"
"os"
"path/filepath"
"sort"
"strings"
)
type LineCoverage struct {
LineNumber int
Content    string
Count      int
IsCovered  bool
}
type FileWithSource struct {
FileName string
Lines    []LineCoverage
Total    int
Covered  int
}
func GetFileWithSource(filePath string, coverage *FileCoverage) (*FileWithSource, error) {
file, err := os.Open(filePath)
if err != nil {
return &FileWithSource{
FileName: filePath,
Lines:    []LineCoverage{},
}, nil
}
defer file.Close()
lines := []LineCoverage{}
scanner := bufio.NewScanner(file)
lineNum := 0
for scanner.Scan() {
lineNum++
lines = append(lines, LineCoverage{
LineNumber: lineNum,
Content:    scanner.Text(),
Count:      0,
IsCovered:  false,
})
}
if err := scanner.Err(); err != nil {
return nil, err
}
for _, block := range coverage.Blocks {
for i := block.StartLine; i <= block.EndLine; i++ {
if i > 0 && i <= len(lines) {
lines[i-1].Count = block.Count
lines[i-1].IsCovered = block.Count > 0
}
}
}
total, covered, _ := coverage.GetCoverageStats()
return &FileWithSource{
FileName: filePath,
Lines:    lines,
Total:    total,
Covered:  covered,
}, nil
}
type FileNode struct {
Name     string
Path     string
IsDir    bool
Children []*FileNode
Coverage *FileCoverage
}
func BuildFileTree(files map[string]*FileCoverage) *FileNode {
root := &FileNode{
Name:     "root",
Path:     "",
IsDir:    true,
Children: []*FileNode{},
}
for path, coverage := range files {
parts := strings.Split(path, string(filepath.Separator))
current := root
for i, part := range parts {
isLast := i == len(parts)-1
var child *FileNode
for _, c := range current.Children {
if c.Name == part {
child = c
break
}
}
if child == nil {
child = &FileNode{
Name:     part,
Path:     strings.Join(parts[:i+1], "/"),
IsDir:    !isLast,
Children: []*FileNode{},
}
current.Children = append(current.Children, child)
}
if isLast {
child.Coverage = coverage
}
current = child
}
}
sortFileTree(root)
return root
}
func sortFileTree(node *FileNode) {
sort.Slice(node.Children, func(i, j int) bool {
if node.Children[i].IsDir != node.Children[j].IsDir {
return node.Children[i].IsDir
}
return node.Children[i].Name < node.Children[j].Name
})
for _, child := range node.Children {
sortFileTree(child)
}
}
func (n *FileNode) GetCoveragePercentage() float64 {
if !n.IsDir && n.Coverage != nil {
_, _, pct := n.Coverage.GetCoverageStats()
return pct
}
if n.IsDir {
total := 0
covered := 0
for _, child := range n.Children {
if child.Coverage != nil {
t, c, _ := child.Coverage.GetCoverageStats()
total += t
covered += c
} else if child.IsDir {
pct := child.GetCoveragePercentage()
if pct > 0 {
t, c := child.GetTotalStatements()
total += t
covered += c
}
}
}
if total > 0 {
return float64(covered) / float64(total) * 100
}
}
return 0.0
}
func (n *FileNode) GetTotalStatements() (total, covered int) {
if !n.IsDir && n.Coverage != nil {
t, c, _ := n.Coverage.GetCoverageStats()
return t, c
}
if n.IsDir {
for _, child := range n.Children {
t, c := child.GetTotalStatements()
total += t
covered += c
}
}
return
}
func GetCoverageColor(percentage float64) string {
switch {
case percentage >= 80:
return "#4caf50"
case percentage >= 60:
return "#8bc34a"
case percentage >= 40:
return "#ff9800"
case percentage >= 20:
return "#ff5722"
default:
return "#f44336"
}
}
func FormatPercentage(pct float64) string {
return fmt.Sprintf("%.1f%%", pct)
}
