package main

import (
	"fmt"
	"html/template"
	"math/rand"
	"os"
	"path/filepath"

	"app/data"
)

func getRandomLightColor() string {
	return fmt.Sprintf("hsl(%d, %d%%, %d%%)",
		rand.Intn(360),
		60+rand.Intn(10),
		90+rand.Intn(8))
}

func createStaticIndexHTML(outputDir string) error {

	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return err
	}

	pageData := data.GetPageData()

	tmplData := data.PrepareTemplateData(pageData)

	tmpl, err := template.ParseFiles("templates/template.html")
	if err != nil {
		return fmt.Errorf("template parsing failed: %v", err)
	}

	outputPath := filepath.Join(outputDir, "index.html")
	f, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("output file creation failed: %v", err)
	}
	defer f.Close()

	if err := tmpl.Execute(f, tmplData); err != nil {
		return fmt.Errorf("template execution failed: %v", err)
	}

	fmt.Printf("Static HTML generated to %s\n", outputPath)
	return nil
}

func copyStaticFiles(outputDir, staticDir string) error {

	files := []string{
		filepath.Join(staticDir, "style.css"),
		filepath.Join(staticDir, "favicon.ico"),
		filepath.Join(staticDir, "icp.png"),
	}

	for _, file := range files {

		content, err := os.ReadFile(file)
		if err != nil {
			return err
		}

		outFile := filepath.Join(outputDir, filepath.Base(file))
		if err := os.WriteFile(outFile, content, 0644); err != nil {
			return err
		}
		fmt.Printf("Copied %s to %s\n", file, outFile)
	}
	return nil
}

func main() {

	projectRoot := "../../"
	os.Chdir(projectRoot)
	outputDir := "dist"
	staticDir := "static"

	fmt.Println("Starting static file generation...")

	if err := createStaticIndexHTML(outputDir); err != nil {
		fmt.Printf("Static HTML generation failed: %v\n", err)
		os.Exit(1)
	}

	if err := copyStaticFiles(outputDir, staticDir); err != nil {
		fmt.Printf("Static file copying failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Static website generation complete! Files located in ./dist directory")
}
