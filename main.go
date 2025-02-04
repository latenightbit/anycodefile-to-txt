package main

import (
	"bytes"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

const (
	inputDir         = "./input"
	outputDir        = "./text_output"
	filePerm         = 0755
	fileMode         = 0644
	combinedFileName = "everything_combined.txt"
	separatorLength  = 40
)

func main() {
	if err := setupDirectories(); err != nil {
		exitWithError(err)
	}

	results, combinedContent, err := processFiles()
	if err != nil {
		exitWithError(err)
	}

	if err := createCombinedFile(combinedContent); err != nil {
		exitWithError(err)
	}

	printSummary(results, combinedContent)
}

type processResults struct {
	success, failed int
}

func processFiles() (*processResults, *bytes.Buffer, error) {
	entries, err := os.ReadDir(inputDir)
	if err != nil {
		return nil, nil, fmt.Errorf("error reading input directory: %w", err)
	}

	results := &processResults{}
	combinedContent := bytes.NewBuffer(nil)

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		filename := entry.Name()
		content, err := os.ReadFile(filepath.Join(inputDir, filename))
		if err != nil {
			results.failed++
			fmt.Printf("⚠️ Error processing %s: %v\n", filename, err)
			continue
		}

		outputPath := filepath.Join(outputDir, filename+".txt")
		if err := os.WriteFile(outputPath, content, fs.FileMode(fileMode)); err != nil {
			results.failed++
			fmt.Printf("⚠️ Error writing %s: %v\n", filename, err)
			continue
		}

		results.success++
		fmt.Printf("✅ Created: %s.txt\n", filename)
		writeFileToCombined(combinedContent, filename, content)
	}

	return results, combinedContent, nil
}

func writeFileToCombined(buffer *bytes.Buffer, filename string, content []byte) {
	separator := strings.Repeat("=", separatorLength)
	fmt.Fprintf(buffer, "\n%s\nFILE: %s\n%s\n\n", separator, filename, separator)
	buffer.Write(content)
}

func createCombinedFile(content *bytes.Buffer) error {
	if content.Len() == 0 {
		return nil
	}

	return os.WriteFile(
		filepath.Join(outputDir, combinedFileName),
		content.Bytes(),
		fs.FileMode(fileMode),
	)
}

func printSummary(results *processResults, content *bytes.Buffer) {
	fmt.Printf("\nProcessed %d files (%d successes, %d failures)\n",
		results.success+results.failed,
		results.success,
		results.failed,
	)

	if content.Len() > 0 {
		fmt.Printf("📚 Combined output: %s/%s\n", outputDir, combinedFileName)
	}
}

func setupDirectories() error {
	if _, err := os.Stat(inputDir); os.IsNotExist(err) {
		if err := os.Mkdir(inputDir, filePerm); err != nil {
			return fmt.Errorf("failed to create input directory: %w", err)
		}
		fmt.Printf("Created input directory: %s\n", inputDir)
	}

	if _, err := os.Stat(outputDir); !os.IsNotExist(err) {
		os.RemoveAll(outputDir)
	}
	if err := os.Mkdir(outputDir, filePerm); err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	return nil
}

func exitWithError(err error) {
	fmt.Fprintf(os.Stderr, "\n🚨 Critical error: %v\n", err)
	os.Exit(1)
}
