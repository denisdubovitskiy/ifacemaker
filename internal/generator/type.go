package generator

import (
	"fmt"
	"go/ast"
	"strings"
)

const (
	TypeKindSelector = "selector"
	TypeKindStar     = "star"
	TypeKindIdent    = "ident"
	TypeKindEllipsis = "ellipsis"
	TypeKindArray    = "array"
	TypeKindFunc     = "func"
	TypeKindMap      = "map"
)

type Type struct {
	Child   *Type
	Package string
	Name    string
	Kind    string

	// For function parameters only
	Results []*Param
	Params  []*Param

	// For maps only
	mapKeyType string
	mapValType string
}

func (t *Type) String() string {
	switch t.Kind {
	case TypeKindStar:
		return "*" + t.Child.String()
	case TypeKindIdent:
		if t.Package == "" {
			return t.Name
		}
		return fmt.Sprintf("%s.%s", t.Package, t.Name)
	case TypeKindEllipsis:
		return fmt.Sprintf("...%s", t.Child.String())
	case TypeKindArray:
		return fmt.Sprintf("[]%s", t.Child.String())
	case TypeKindMap:
		return fmt.Sprintf("map[%s]%s", t.mapKeyType, t.mapValType)
	case TypeKindSelector:
		return fmt.Sprintf("%s.%s", t.Package, t.Name)
	case TypeKindFunc:
		params := make([]string, len(t.Params))
		for i, p := range t.Params {
			params[i] = p.String()
		}

		results := make([]string, len(t.Results))
		for i, r := range t.Results {
			results[i] = r.String()
		}

		resultsTemplate := "%s"
		if len(results) > 1 {
			resultsTemplate = "(%s)"
		}

		return fmt.Sprintf(
			"func(%s) "+resultsTemplate,
			strings.Join(params, ", "),
			strings.Join(results, ", "),
		)
	}
	return ""
}

func ParseType(
	node ast.Node,
	typesMap map[string]struct{},
	sourcePackageName string,
) *Type {
	formatPackage := func(pkg, typeName string) string {
		if pkg != "" {
			return ""
		}
		if _, ok := typesMap[typeName]; ok {
			return sourcePackageName
		}
		return pkg
	}

	switch paramType := node.(type) {
	case *ast.SelectorExpr:
		return &Type{
			Name:    paramType.Sel.Name,
			Package: identName(paramType.X),
			Kind:    TypeKindSelector,
		}
	case *ast.Ident:
		return &Type{
			Name:    identName(paramType),
			Package: formatPackage("", identName(paramType)),
			Kind:    TypeKindIdent,
		}
	case *ast.Ellipsis:
		return &Type{
			Child: ParseType(paramType.Elt, typesMap, sourcePackageName),
			Kind:  TypeKindEllipsis,
		}
	case *ast.StarExpr:
		return &Type{
			Child: ParseType(paramType.X, typesMap, sourcePackageName),
			Kind:  TypeKindStar,
		}
	case *ast.FuncType:
		return &Type{
			Params:  ParseMany(extractList(paramType.Params), typesMap, sourcePackageName),
			Results: ParseMany(extractList(paramType.Results), typesMap, sourcePackageName),
			Kind:    TypeKindFunc,
		}
	case *ast.ArrayType:
		return &Type{
			Child: ParseType(paramType.Elt, typesMap, sourcePackageName),
			Kind:  TypeKindArray,
		}
	case *ast.MapType:
		return &Type{
			Kind:       TypeKindMap,
			mapKeyType: paramType.Key.(*ast.Ident).Name,
			mapValType: paramType.Value.(*ast.Ident).Name,
		}
	default:
		panic("unhandled type")
	}

	return nil
}

func parseTypesFromFile(fileAst *ast.File) []string {
	var types []string

	ast.Inspect(fileAst, func(node ast.Node) bool {
		ts, ok := node.(*ast.TypeSpec)
		if !ok || !ts.Name.IsExported() {
			return true
		}

		types = append(types, ts.Name.Name)

		return true
	})

	return types
}
