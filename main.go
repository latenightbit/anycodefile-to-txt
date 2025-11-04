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
	results := &processResults{}
	combinedContent := bytes.NewBuffer(nil)

	err := filepath.WalkDir(inputDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			results.failed++
			fmt.Printf("⚠️ Error accessing %s: %v\n", path, err)
			return nil
		}

		if d.IsDir() {
			return nil
		}

		if d.Name() == ".DS_Store" {
			return nil
		}

		relPath, err := filepath.Rel(inputDir, path)
		if err != nil {
			results.failed++
			fmt.Printf("⚠️ Error processing %s: %v\n", path, err)
			return nil
		}

		baseName := filepath.Base(relPath)
		displayPath := filepath.ToSlash(relPath)

		content, err := os.ReadFile(path)
		if err != nil {
			results.failed++
			fmt.Printf("⚠️ Error processing %s: %v\n", displayPath, err)
			return nil
		}

		outputPath := filepath.Join(outputDir, relPath) + ".txt"
		if err := os.MkdirAll(filepath.Dir(outputPath), filePerm); err != nil {
			results.failed++
			fmt.Printf("⚠️ Error preparing directory for %s: %v\n", displayPath, err)
			return nil
		}

		fileBuffer := bytes.NewBuffer(nil)
		fmt.Fprintf(fileBuffer, "FILE: %s\nDESTINATION: %s\n\n", baseName, displayPath)
		fileBuffer.Write(content)

		if err := os.WriteFile(outputPath, fileBuffer.Bytes(), fs.FileMode(fileMode)); err != nil {
			results.failed++
			fmt.Printf("⚠️ Error writing %s: %v\n", displayPath, err)
			return nil
		}

		results.success++
		fmt.Printf("✅ Created: %s.txt\n", displayPath)
		writeFileToCombined(combinedContent, displayPath, content)
		return nil
	})

	if err != nil {
		return nil, nil, fmt.Errorf("error walking input directory: %w", err)
	}

	return results, combinedContent, nil
}

func writeFileToCombined(buffer *bytes.Buffer, displayPath string, content []byte) {
	separator := strings.Repeat("=", separatorLength)
	fmt.Fprintf(
		buffer,
		"\n%s\nFILE: %s\nDESTINATION: %s\n%s\n\n",
		separator,
		filepath.Base(displayPath),
		displayPath,
		separator,
	)
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
