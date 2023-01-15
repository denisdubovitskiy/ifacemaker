package generator

import (
	"go/ast"
)

type Param struct {
	Name string
	Type *Type
}

func (p Param) String() string {
	if p.Name == "" {
		return p.Type.String()
	}
	return p.Name + " " + p.Type.String()
}

func ParseMany(list []*ast.Field, declaredTypesMap map[string]struct{}, sourcePackageName string) []*Param {
	if len(list) == 0 {
		return nil
	}

	params := make([]*Param, 0, len(list))

	for _, p := range list {
		parsed := Parse(p, declaredTypesMap, sourcePackageName)
		params = append(params, parsed...)
	}

	return params
}

func Parse(
	field *ast.Field,
	typesMap map[string]struct{},
	sourcePackageName string,
) []*Param {
	params := make([]*Param, 0, len(field.Names))

	if field.Names == nil {
		param := &Param{
			Name: "",
			Type: ParseType(
				field.Type,
				typesMap,
				sourcePackageName,
			),
		}
		params = append(params, param)
	}

	for _, name := range field.Names {
		param := &Param{
			Name: name.Name,
			Type: ParseType(
				field.Type,
				typesMap,
				sourcePackageName,
			),
		}

		params = append(params, param)
	}

	return params
}
