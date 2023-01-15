package generator

import "go/ast"

func extractList(fieldList *ast.FieldList) []*ast.Field {
	if fieldList == nil || fieldList.List == nil {
		return nil
	}
	return fieldList.List
}

func identName(node ast.Node) string {
	return node.(*ast.Ident).Name
}
