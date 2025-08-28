package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"

	"gopkg.in/yaml.v3"
)

// PathInstruction represents a single path instruction
type PathInstruction struct {
	Path         string `yaml:"path"`
	Instructions string `yaml:"instructions"`
}

// TemplateData holds the data for the template
type TemplateData struct {
	PathInstructions []PathInstruction
}

// Helper function to indent text
func indentText(text string, indent string) string {
	lines := strings.Split(text, "\n")
	var result []string

	for _, line := range lines {
		if strings.TrimSpace(line) != "" {
			result = append(result, indent+line)
		} else {
			result = append(result, "")
		}
	}

	return strings.Join(result, "\n")
}

// generateFilePattern generates a file-based pattern from a folder-based pattern
func generateFilePattern(folderPattern string) string {
	// Extract keywords from folder pattern like "**/{http,client,api,network}*/**"
	// and convert to file pattern like "**/*{http*,client*,api*,network*}*"

	// Match pattern like **/{keyword1,keyword2,keyword3}*/**
	re := regexp.MustCompile(`\*\*/\{([^}]+)\}\*/\*\*`)
	matches := re.FindStringSubmatch(folderPattern)

	if len(matches) != 2 {
		// If pattern doesn't match expected format, return original
		return folderPattern
	}

	// Extract keywords and convert to file pattern
	keywords := strings.Split(matches[1], ",")
	var fileKeywords []string

	for _, keyword := range keywords {
		keyword = strings.TrimSpace(keyword)
		if keyword != "" {
			fileKeywords = append(fileKeywords, keyword+"*")
		}
	}

	// Create file pattern
	filePattern := "**/*{" + strings.Join(fileKeywords, ",") + "}*"
	return filePattern
}

// expandPathInstructions takes folder-based patterns and creates both folder and file patterns
func expandPathInstructions(instructions []PathInstruction) []PathInstruction {
	var expanded []PathInstruction

	for _, instruction := range instructions {
		// Add the original folder-based pattern
		expanded = append(expanded, instruction)

		// Generate and add the file-based pattern
		filePattern := generateFilePattern(instruction.Path)
		if filePattern != instruction.Path { // Only add if different
			fileInstruction := PathInstruction{
				Path:         filePattern,
				Instructions: instruction.Instructions,
			}
			expanded = append(expanded, fileInstruction)
		}
	}

	return expanded
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run cmd/merge/main.go <configs_dir> [output_file]")
		fmt.Println("  configs_dir: Directory containing technology-specific configs")
		fmt.Println("  output_file: Output file (default: .coderabbit.yaml)")
		os.Exit(1)
	}

	configsDir := os.Args[1]
	outputFile := ".coderabbit.yaml"
	if len(os.Args) > 2 {
		outputFile = os.Args[2]
	}

	// Collect all path instructions
	pathInstructions, err := collectPathInstructions(configsDir)
	if err != nil {
		log.Fatalf("Failed to collect path instructions: %v", err)
	}

	// Expand instructions to include both folder and file patterns
	expandedInstructions := expandPathInstructions(pathInstructions)

	// Generate the final config using template
	err = generateConfig(expandedInstructions, outputFile)
	if err != nil {
		log.Fatalf("Failed to generate config: %v", err)
	}

	fmt.Printf("Successfully merged %d path instructions into %s (expanded to %d total patterns)\n",
		len(pathInstructions), outputFile, len(expandedInstructions))
}

// collectPathInstructions collects all path instructions from individual config files
func collectPathInstructions(configsDir string) ([]PathInstruction, error) {
	var allInstructions []PathInstruction

	err := filepath.WalkDir(configsDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// Skip directories and non-yaml files
		if d.IsDir() || !strings.HasSuffix(d.Name(), ".yaml") {
			return nil
		}

		// Skip if not a path_instructions.yaml file
		if d.Name() != "path_instructions.yaml" {
			return nil
		}

		fmt.Printf("Processing: %s\n", path)

		// Read and parse the individual config file
		data, err := os.ReadFile(path)
		if err != nil {
			return fmt.Errorf("failed to read %s: %w", path, err)
		}

		var pathInstructions []PathInstruction
		err = yaml.Unmarshal(data, &pathInstructions)
		if err != nil {
			return fmt.Errorf("failed to parse %s: %w", path, err)
		}

		// Add the path instructions to the collection
		allInstructions = append(allInstructions, pathInstructions...)

		fmt.Printf("  Added %d path instructions\n", len(pathInstructions))
		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to walk configs directory: %w", err)
	}

	fmt.Printf("Total path instructions collected: %d\n", len(allInstructions))
	return allInstructions, nil
}

// generateConfig generates the final config using the base template
func generateConfig(pathInstructions []PathInstruction, outputFile string) error {
	// Read the base template
	baseTemplate, err := os.ReadFile(".coderabbit.base.yaml")
	if err != nil {
		return fmt.Errorf("failed to read .coderabbit.base.yaml: %w", err)
	}

	// Create template data
	data := TemplateData{
		PathInstructions: pathInstructions,
	}

	// Create template with custom functions
	tmpl, err := template.New("coderabbit").Funcs(template.FuncMap{
		"indent": func(text string) string {
			return indentText(text, "        ")
		},
	}).Parse(string(baseTemplate))
	if err != nil {
		return fmt.Errorf("failed to parse template: %w", err)
	}

	// Create output file
	file, err := os.Create(outputFile)
	if err != nil {
		return fmt.Errorf("failed to create %s: %w", outputFile, err)
	}
	defer file.Close()

	// Execute template
	err = tmpl.Execute(file, data)
	if err != nil {
		return fmt.Errorf("failed to execute template: %w", err)
	}

	return nil
}
