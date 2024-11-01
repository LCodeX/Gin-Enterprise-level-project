//go:generate go run generate.go

package main

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

type LayerData struct {
	Name string // The name for each module, such as "Packages" or "User"
}

// Custom function to convert strings to lowercase
func toLower(s string) string {
	return strings.ToLower(s)
}

func main() {
	// Define the modules to be generated
	layerData := []LayerData{
		{Name: "News"},
	}

	// Parse the template file from the generate directory
	tmpl, err := template.New("layer").Funcs(template.FuncMap{"ToLower": toLower}).ParseFiles("layer_template.tmpl")
	if err != nil {
		fmt.Println("Error parsing template:", err)
		return
	}

	// Generate files for each module in layerData
	for _, data := range layerData {
		// Generate controller, service, and DAO files, placing them in ../controllers, ../services, and ../dao
		generateFile(tmpl, data, fmt.Sprintf("../controllers/%s.go", strings.ToLower(data.Name)), "controller")
		generateFile(tmpl, data, fmt.Sprintf("../services/%s.go", strings.ToLower(data.Name)), "service")
		generateFile(tmpl, data, fmt.Sprintf("../dao/%s.go", strings.ToLower(data.Name)), "dao")
	}
}

func generateFile(tmpl *template.Template, data LayerData, filePath string, templateName string) {
	var buf bytes.Buffer
	// Execute the specific template content for controllers, services, or dao
	if err := tmpl.ExecuteTemplate(&buf, templateName, data); err != nil {
		fmt.Printf("Error generating %s: %v\n", filePath, err)
		return
	}

	// Ensure the target directory exists
	if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
		fmt.Printf("Error creating directory for %s: %v\n", filePath, err)
		return
	}

	// Write the generated content to the file
	if err := os.WriteFile(filePath, buf.Bytes(), 0644); err != nil {
		fmt.Printf("Error writing file %s: %v\n", filePath, err)
		return
	}

	fmt.Printf("Generated %s\n", filePath)
}
