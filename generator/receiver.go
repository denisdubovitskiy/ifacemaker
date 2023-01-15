package generator

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/printer"
	"go/token"
	"strings"
)

type Receiver struct {
	Params  []*Param
	Results []*Param
	Name    string
	Comment string
}

func (r Receiver) String() string {
	var (
		params  []string
		results []string
	)

	for _, p := range r.Params {
		params = append(params, p.String())
	}

	for _, p := range r.Results {
		results = append(results, p.String())
	}

	var comment string

	if r.Comment != "" {
		comment = r.Comment
	}

	if len(r.Results) > 0 {
		return fmt.Sprintf(comment+"%s(%s)(%s)", r.Name, strings.Join(params, ", "), strings.Join(results, ", "))
	}
	return fmt.Sprintf(comment+"%s(%s)", r.Name, strings.Join(params, ", "))
}

func ParseReceivers(
	astFile *ast.File,
	fset *token.FileSet,
	structName string,
	sourcePackageName string,
	declaredTypesMap map[string]struct{},
) []Receiver {
	var receivers []Receiver

	ast.Inspect(astFile, func(node ast.Node) bool {
		funcDecl, ok := node.(*ast.FuncDecl)
		if !ok {
			return true
		}

		// don't care about private ones
		if !isReceiver(funcDecl) || !isFuncExported(funcDecl) {
			return true
		}

		recvTypeName := formatNode(fset, funcDecl.Recv.List[0].Type)

		// remove a star if there is any, so we
		// can make assertions against a user-provided type
		recvTypeName = strings.TrimPrefix(recvTypeName, "*")

		// other type's receiver
		if recvTypeName != structName {
			return true
		}

		name := funcDecl.Name.String()

		receiver := Receiver{
			Comment: parseReceiverDocs(extractComments(funcDecl.Doc)),
			Params:  ParseMany(extractList(funcDecl.Type.Params), declaredTypesMap, sourcePackageName),
			Results: ParseMany(extractList(funcDecl.Type.Results), declaredTypesMap, sourcePackageName),
			Name:    name,
		}

		receivers = append(receivers, receiver)

		return true
	})

	return receivers
}

func extractComments(doc *ast.CommentGroup) []*ast.Comment {
	if doc == nil || doc.List == nil {
		return nil
	}
	return doc.List
}

func isFuncExported(n *ast.FuncDecl) bool {
	return n.Name.IsExported()
}

func isReceiver(n *ast.FuncDecl) bool {
	return n.Recv != nil
}

func formatNode(fs *token.FileSet, node any) string {
	var buf bytes.Buffer
	_ = printer.Fprint(&buf, fs, node)
	return buf.String()
}

func parseReceiverDocs(lines []*ast.Comment) string {
	if len(lines) == 0 {
		return ""
	}

	comments := make([]string, len(lines))

	for i, line := range lines {
		comments[i] = line.Text
	}

	return strings.TrimSuffix(strings.Join(comments, "\n"), "\n") + "\n"
}
