// Package gen has functions that help with generating code and test templates for new day
package gen

import (
	_ "embed"
	"fmt"
	"io"
	"os"
	"text/template"
)

//go:embed code.txt
var code string

//go:embed test.txt
var test string

func writeFromTemplate(target io.Writer, name string, content string, year int, day int) {
	payload := struct {
		Year string
		Day  string
	}{
		fmt.Sprintf("%d", year),
		fmt.Sprintf("%02d", day),
	}

	tpl := template.Must(template.New(name).Parse(content))
	err := tpl.Execute(target, payload)
	if err != nil {
		panic("Template does not exist")
	}
}

// GenerateSources generates the source code stub given day and year
func GenerateSources(year int, day int) {
	yearDir := fmt.Sprintf("year%d", year%2000)
	dayDir := fmt.Sprintf("day%02d", day)

	err := os.Mkdir(yearDir+"/"+dayDir, os.ModePerm)
	check(err)

	inputFileFormat := fmt.Sprintf("%s/%s/input.txt", yearDir, dayDir)
	codeFileFormat := fmt.Sprintf("%s/%s/main.go", yearDir, dayDir)
	testFileFormat := fmt.Sprintf("%s/%s/main_test.go", yearDir, dayDir)

	inputFile, err := os.Create(inputFileFormat)
	check(err)

	codeFile, err := os.Create(codeFileFormat)
	check(err)

	testFile, err := os.Create(testFileFormat)
	check(err)

	defer checkCloseFile(inputFile)
	defer checkCloseFile(codeFile)
	defer checkCloseFile(testFile)

	writeFromTemplate(codeFile, "code", code, year, day)
	writeFromTemplate(testFile, "test", test, year, day)

	fmt.Printf(
		"[SUCCESS] Generated Files:\n\t%s\n\t%s\n\t%s\n",
		inputFile.Name(),
		codeFile.Name(),
		testFile.Name())
}

func checkCloseFile(file *os.File) {
	err := file.Close()
	if err != nil {
		panic(err)
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
