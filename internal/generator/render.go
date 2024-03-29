package generator

import (
	"strings"

	"golang.org/x/tools/imports"
)

func RenderInterface(
	packageName string,
	interfaceName string,
	sourceStructName string,
	sourcePkgName string,
	modulePath string,
	outputFilename string,
	receivers []Receiver,
) (
	[]byte,
	error,
) {
	var b strings.Builder

	// generated comment
	b.WriteString("// Package ")
	b.WriteString(packageName)
	b.WriteString(" generated with github.com/denisdubovitskiy/ifacemaker, DO NOT EDIT.\n")

	// header
	b.WriteString("package ")
	b.WriteString(packageName)
	b.WriteString("\n")

	b.WriteString("//go:generate ifacemaker")
	b.WriteString(" --source-pkg ")
	b.WriteString(sourcePkgName)
	b.WriteString(" --module-path ")
	b.WriteString(modulePath)
	b.WriteString(" --result-pkg ")
	b.WriteString(packageName)
	b.WriteString(" --struct-name ")
	b.WriteString(sourceStructName)
	b.WriteString(" --interface-name ")
	b.WriteString(interfaceName)
	b.WriteString(" --output ")
	b.WriteString(outputFilename)
	b.WriteString("\n")

	// interface header
	b.WriteString("type ")
	b.WriteString(interfaceName)
	b.WriteString(" interface {\n")

	for _, receiver := range receivers {
		b.WriteString(receiver.String())
		b.WriteString("\n")
	}

	// interface footer
	b.WriteString("}\n")

	return formatCodeWithGoImports(b.String())
}

func formatCodeWithGoImports(code string) ([]byte, error) {
	return imports.Process("", []byte(code), &imports.Options{
		TabIndent: true,
		TabWidth:  4,
		Fragment:  true,
		Comments:  true,
	})
}
