package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/parser"
)

type PageData struct {
	Title   string
	Content string
}

const defaultTemplate = `<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>{{.Title}}</title>
  <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/github-markdown-css/github-markdown.css">
  <style>
    body {
      margin: 0;
      padding: 0;
    }

    .markdown-body {
      box-sizing: border-box;
      min-width: 200px;
      max-width: 980px;
      margin: 0 auto;
      padding: 45px;
    }

    @media (max-width: 767px) {
      .markdown-body {
        padding: 15px;
      }
    }
  </style>
</head>
<body>
  <div class="markdown-body">
    {{.Content}}
  </div>
</body>
</html>`

func minifyHTML(html string) string {
	commentRegex := regexp.MustCompile(`<!--.*?-->`)
	html = commentRegex.ReplaceAllString(html, "")

	whitespaceRegex := regexp.MustCompile(`>\s+<`)
	html = whitespaceRegex.ReplaceAllString(html, "><")

	spaceRegex := regexp.MustCompile(`\s{2,}`)
	html = spaceRegex.ReplaceAllString(html, " ")

	html = strings.ReplaceAll(html, "\n", "")

	return html
}

func extractFirstTitle(mdContent string) string {
	atxRegex := regexp.MustCompile(`(?m)^#\s+(.+)$`)
	matches := atxRegex.FindStringSubmatch(mdContent)
	if len(matches) > 1 {
		return strings.TrimSpace(matches[1])
	}

	setextRegex := regexp.MustCompile(`(?m)^(.+)\n={2,}$`)
	matches = setextRegex.FindStringSubmatch(mdContent)
	if len(matches) > 1 {
		return strings.TrimSpace(matches[1])
	}

	return ""
}

func main() {
	inputDir := flag.String("input", ".", "Input directory for Markdown files")
	outputDir := flag.String("output", "./output", "Output directory for HTML files")
	compress := flag.Bool("compress", false, "Compress HTML output")
	flag.Parse()

	if err := os.MkdirAll(*outputDir, 0755); err != nil {
		fmt.Printf("Failed to create output directory: %v\n", err)
		os.Exit(1)
	}

	tmpl, err := template.New("html").Parse(defaultTemplate)
	if err != nil {
		fmt.Printf("Failed to parse built-in template: %v\n", err)
		os.Exit(1)
	}

	fileCount := 0
	err = filepath.Walk(*inputDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && (strings.HasSuffix(path, ".md") || strings.HasSuffix(path, ".markdown")) {
			fileCount++
			if err := convertMarkdownToHTML(path, *outputDir, tmpl, *inputDir, *compress); err != nil {
				fmt.Printf("Conversion failed: %s - %v\n", path, err)
			}
		}
		return nil
	})

	if err != nil {
		fmt.Printf("Failed to traverse directory: %v\n", err)
		os.Exit(1)
	}

	if fileCount == 0 {
		fmt.Println("No Markdown files found in the specified directory.")
	} else {
		fmt.Printf("Conversion completed! Processed %d Markdown files.\n", fileCount)
	}
}

func convertMarkdownToHTML(mdPath, outputDir string, tmpl *template.Template, inputDir string, compress bool) error {
	mdContent, err := os.ReadFile(mdPath)
	if err != nil {
		return fmt.Errorf("failed to read file: %v", err)
	}

	title := extractFirstTitle(string(mdContent))
	if title == "" {
		baseName := filepath.Base(mdPath)
		title = strings.TrimSuffix(baseName, filepath.Ext(baseName))
	}

	extensions := parser.CommonExtensions | parser.AutoHeadingIDs
	parser := parser.NewWithExtensions(extensions)
	html := markdown.ToHTML(mdContent, parser, nil)

	relPath, err := filepath.Rel(inputDir, mdPath)
	if err != nil {
		relPath = filepath.Base(mdPath)
	}
	outputPath := filepath.Join(outputDir, strings.TrimSuffix(relPath, filepath.Ext(relPath))+".html")
	outputDir = filepath.Dir(outputPath)

	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return fmt.Errorf("failed to create output subdirectory: %v", err)
	}

	outFile, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create output file: %v", err)
	}
	defer outFile.Close()

	data := PageData{
		Title:   title,
		Content: string(html),
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return fmt.Errorf("failed to render template: %v", err)
	}
	htmlContent := buf.String()

	if compress {
		htmlContent = minifyHTML(htmlContent)
	}

	if _, err := outFile.WriteString(htmlContent); err != nil {
		return fmt.Errorf("failed to write file: %v", err)
	}

	return nil
}
