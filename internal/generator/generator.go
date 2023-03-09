package generator

import (
	"go/ast"
	"go/parser"
	"go/token"
	"os"
)

type Options struct {
	Files             []string
	StructName        string
	InterfaceName     string
	OutputPackageName string
}

func Generate(options Options) ([]byte, error) {
	var sourcePackageName string
	var interfaceDoc string
	parsedDeclaredTypes := make(map[string]struct{})

	for _, f := range options.Files {
		src, err := os.ReadFile(f)
		if err != nil {
			return nil, err
		}

		parsed, err := parser.ParseFile(token.NewFileSet(), "", src, parser.ParseComments)
		if err != nil {
			return nil, err
		}

		if sourcePackageName == "" {
			sourcePackageName = identName(parsed.Name)
		}

		if interfaceDoc == "" {
			interfaceDoc = parseInterfaceDoc(parsed, options.StructName)
		}

		for _, t := range parseTypesFromFile(parsed) {
			parsedDeclaredTypes[t] = struct{}{}
		}
	}

	var receivers []Receiver
	fileSet := token.NewFileSet()

	for _, f := range options.Files {
		src, err := os.ReadFile(f)
		if err != nil {
			return nil, err
		}

		parsed, err := parser.ParseFile(fileSet, "", src, parser.ParseComments)
		if err != nil {
			return nil, err
		}

		fileReceivers := ParseReceivers(
			parsed,
			fileSet,
			options.StructName,
			sourcePackageName,
			parsedDeclaredTypes,
		)
		receivers = append(receivers, fileReceivers...)
	}

	return RenderInterface(
		options.OutputPackageName,
		options.InterfaceName,
		receivers,
	)
}

func parseInterfaceDoc(parsed *ast.File, structName string) string {
	ast.Inspect(parsed, func(node ast.Node) bool {
		n, ok := node.(*ast.TypeSpec)
		if !ok {
			return true
		}

		if n.Name.String() != structName {
			return true
		}

		if n.Doc == nil {
			return true
		}
		// TODO
		return true
	})

	return ""
}
