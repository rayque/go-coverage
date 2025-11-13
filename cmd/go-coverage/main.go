package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	coverage "github.com/rayque/go-coverage/pkg"
)

var version = "1.0.0"

func main() {
	inputFile := flag.String("input", "coverage.out", "Path to the coverage file")
	outputFile := flag.String("output", "coverage.html", "Path to the output HTML file")
	showVersion := flag.Bool("version", false, "Show version information")
	quiet := flag.Bool("quiet", false, "Suppress output messages")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Go Coverage HTML Reporter v%s\n\n", version)
		fmt.Fprintf(os.Stderr, "Usage: go-coverage [options]\n\n")
		fmt.Fprintf(os.Stderr, "Options:\n")
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\nExamples:\n")
		fmt.Fprintf(os.Stderr, "  go-coverage\n")
		fmt.Fprintf(os.Stderr, "  go-coverage -input=coverage.out -output=report.html\n")
	}
	flag.Parse()
	if *showVersion {
		fmt.Printf("go-coverage v%s\n", version)
		os.Exit(0)
	}
	if _, err := os.Stat(*inputFile); os.IsNotExist(err) {
		log.Fatalf("Error: Coverage file '%s' does not exist\n", *inputFile)
	}
	if !*quiet {
		fmt.Printf("📊 Parsing coverage file: %s\n", *inputFile)
	}
	report, err := coverage.ParseCoverageFile(*inputFile)
	if err != nil {
		log.Fatalf("Error parsing coverage file: %v\n", err)
	}
	if !*quiet {
		totalStmts, coveredStmts, overallPct := report.GetOverallStats()
		fmt.Printf("📈 Overall coverage: %.1f%% (%d/%d statements)\n", overallPct, coveredStmts, totalStmts)
		fmt.Printf("📁 Files analyzed: %d\n", len(report.Files))

		// Check if source files exist using smart path resolution
		missingFiles := 0
		for filePath := range report.Files {
			found := false
			// Try the same alternatives as GetFileWithSource
			alternatives := []string{filePath}

			// Add just the filename
			alternatives = append(alternatives, filepath.Base(filePath))

			// Try parts after the module name
			parts := strings.Split(filePath, "/")
			if len(parts) > 1 {
				for i := 1; i < len(parts); i++ {
					alternatives = append(alternatives, filepath.Join(parts[i:]...))
				}
			}

			// Check if any alternative exists
			for _, alt := range alternatives {
				if _, err := os.Stat(alt); err == nil {
					found = true
					break
				}
			}

			if !found {
				missingFiles++
			}
		}
		if missingFiles > 0 {
			fmt.Printf("⚠️  Warning: %d source files not found in current directory\n", missingFiles)
			fmt.Printf("   Make sure you run this tool from your project root directory\n")
		}

		fmt.Printf("🔨 Generating HTML report: %s\n", *outputFile)
	}
	err = coverage.GenerateHTMLReport(report, *outputFile)
	if err != nil {
		log.Fatalf("Error generating HTML report: %v\n", err)
	}
	if !*quiet {
		fmt.Printf("✅ Report generated successfully!\n")
		fmt.Printf("🌐 Open %s in your browser to view the report\n", *outputFile)
	}
}
