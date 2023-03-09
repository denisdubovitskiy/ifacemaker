package generator

import (
	"strings"

	"golang.org/x/tools/imports"
)

func RenderInterface(
	packageName string,
	interfaceName string,
	receivers []Receiver,
) (
	[]byte,
	error,
) {
	var b strings.Builder

	// generated comment
	b.WriteString("// Package ")
	b.WriteString(packageName)
	b.WriteString(" generated with github.com/densdubovitskiy/ifacemaker, DO NOT EDIT.\n")

	// header
	b.WriteString("package ")
	b.WriteString(packageName)
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
