package coverage
import (
"fmt"
"html/template"
"os"
"path/filepath"
"sort"
)
type HTMLReport struct {
Report *CoverageReport
}
type FileInfo struct {
Path       string
Name       string
Coverage   float64
Total      int
Covered    int
Color      string
Lines      []LineCoverage
HasSource  bool
}
func GenerateHTMLReport(report *CoverageReport, outputPath string) error {
htmlGen := &HTMLReport{Report: report}
return htmlGen.Generate(outputPath)
}
func (h *HTMLReport) Generate(outputPath string) error {
file, err := os.Create(outputPath)
if err != nil {
return fmt.Errorf("failed to create output file: %w", err)
}
defer file.Close()
tree := BuildFileTree(h.Report.Files)
totalStmts, coveredStmts, overallPct := h.Report.GetOverallStats()
fileInfos := []FileInfo{}
for path, coverage := range h.Report.Files {
total, covered, pct := coverage.GetCoverageStats()
fileWithSource, _ := GetFileWithSource(path, coverage)
fileInfos = append(fileInfos, FileInfo{
Path:      path,
Name:      filepath.Base(path),
Coverage:  pct,
Total:     total,
Covered:   covered,
Color:     GetCoverageColor(pct),
Lines:     fileWithSource.Lines,
HasSource: len(fileWithSource.Lines) > 0,
})
}
sort.Slice(fileInfos, func(i, j int) bool {
return fileInfos[i].Path < fileInfos[j].Path
})
data := map[string]interface{}{
"Mode":         h.Report.Mode,
"TotalStmts":   totalStmts,
"CoveredStmts": coveredStmts,
"OverallPct":   overallPct,
"OverallColor": GetCoverageColor(overallPct),
"Files":        fileInfos,
"FileTree":     tree,
}
tmpl, err := template.New("coverage").Funcs(template.FuncMap{
"formatPct": FormatPercentage,
"getCoverageColor": GetCoverageColor,
}).Parse(getHTMLTemplate())
if err != nil {
return fmt.Errorf("failed to parse template: %w", err)
}
if err := tmpl.Execute(file, data); err != nil {
return fmt.Errorf("failed to execute template: %w", err)
}
return nil
}
func getHTMLTemplate() string {
return htmlTemplateContent
}
