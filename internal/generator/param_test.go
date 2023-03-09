package generator

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"

	assert "github.com/stretchr/testify/require"
)

func testParseAST(t *testing.T, src string) *ast.File {
	t.Helper()
	f, err := parser.ParseFile(token.NewFileSet(), "", []byte(`package awesomepkg; func some(`+src+`)`), parser.ParseComments)
	assert.NoError(t, err, "unable to parse ast: %v", err)
	return f
}

func testParseType(t *testing.T, src string) *ast.Field {
	f := testParseAST(t, src)
	return f.Decls[0].(*ast.FuncDecl).Type.Params.List[0]
}

func TestParamString(t *testing.T) {
	t.Run("ident", func(t *testing.T) {
		f := testParseType(t, `a int`)

		// act
		param := Parse(f, nil, "awesomepkg")

		// assert
		assert.Equal(t, "a int", param[0].String())
	})

	t.Run("selector", func(t *testing.T) {
		f := testParseType(t, `a somepackage.A`)

		// act
		param := Parse(f, nil, "awesomepkg")

		// assert
		assert.Equal(t, "a somepackage.A", param[0].String())
	})

	t.Run("star selector", func(t *testing.T) {
		f := testParseType(t, `a *somepackage.A`)

		// act
		param := Parse(f, nil, "awesomepkg")

		// assert
		assert.Equal(t, "a *somepackage.A", param[0].String())
	})

	t.Run("ellipsis", func(t *testing.T) {
		f := testParseType(t, `a ...int`)

		// act
		param := Parse(f, nil, "awesomepkg")

		// assert
		assert.Equal(t, "a ...int", param[0].String())
	})

	t.Run("star ellipsis", func(t *testing.T) {
		f := testParseType(t, `a ...*int`)

		// act
		param := Parse(f, nil, "awesomepkg")

		// assert
		assert.Equal(t, "a ...*int", param[0].String())
	})

	t.Run("array", func(t *testing.T) {
		f := testParseType(t, `a []int`)

		// act
		param := Parse(f, nil, "awesomepkg")

		// assert
		assert.Equal(t, "a []int", param[0].String())
	})

	t.Run("star array", func(t *testing.T) {
		f := testParseType(t, `a *[]int`)

		// act
		param := Parse(f, nil, "awesomepkg")

		// assert
		assert.Equal(t, "a *[]int", param[0].String())
	})

	t.Run("array star", func(t *testing.T) {
		f := testParseType(t, `a []*int`)

		// act
		param := Parse(f, nil, "awesomepkg")

		// assert
		assert.Equal(t, "a []*int", param[0].String())
	})

	t.Run("func single result", func(t *testing.T) {
		f := testParseType(t, `a func(m int, d bool) error`)

		// act
		param := Parse(f, nil, "awesomepkg")

		// assert
		assert.Equal(t, "a func(m int, d bool) error", param[0].String())
	})

	t.Run("func multiple results", func(t *testing.T) {
		f := testParseType(t, `a func(m int, d bool) (string, error)`)

		// act
		param := Parse(f, nil, "awesomepkg")

		// assert
		assert.Equal(t, "a func(m int, d bool) (string, error)", param[0].String())
	})
}

func TestParseParam(t *testing.T) {
	t.Run("ident", func(t *testing.T) {
		f := testParseType(t, `a int`)

		// act
		param := Parse(f, nil, "awesomepkg")[0]

		// assert
		typ := param.Type
		assert.Equal(t, TypeKindIdent, typ.Kind)
		assert.Equal(t, typ.Name, "int")
		assert.Nil(t, typ.Child)
	})

	t.Run("selector", func(t *testing.T) {
		f := testParseType(t, `a somepackage.A`)

		// act
		param := Parse(f, nil, "awesomepkg")[0]

		// assert
		assert.Equal(t, "a", param.Name)
		typ := param.Type
		assert.Equal(t, TypeKindSelector, typ.Kind)
		assert.Equal(t, typ.Name, "A")
		assert.Equal(t, typ.Package, "somepackage")
		assert.Nil(t, typ.Child)
	})

	t.Run("star selector", func(t *testing.T) {
		f := testParseType(t, `a *somepackage.A`)

		// act
		param := Parse(f, nil, "awesomepkg")[0]

		// assert
		assert.Equal(t, "a", param.Name)
		typ := param.Type
		assert.Equal(t, TypeKindStar, typ.Kind)
		child := typ.Child
		assert.Equal(t, TypeKindSelector, child.Kind)
		assert.Equal(t, child.Name, "A")
		assert.Equal(t, child.Package, "somepackage")
		assert.Nil(t, child.Child)
	})

	t.Run("ellipsis", func(t *testing.T) {
		f := testParseType(t, `a ...int`)

		// act
		param := Parse(f, nil, "awesomepkg")[0]

		// assert
		assert.Equal(t, "a", param.Name)
		typ := param.Type
		assert.Equal(t, TypeKindEllipsis, typ.Kind)
		child := typ.Child
		assert.Equal(t, TypeKindIdent, child.Kind)
		assert.Equal(t, child.Name, "int")
		assert.Nil(t, child.Child)
	})

	t.Run("star ellipsis", func(t *testing.T) {
		f := testParseType(t, `a ...*int`)

		// act
		param := Parse(f, nil, "awesomepkg")[0]

		// assert
		assert.Equal(t, "a", param.Name)
		typ := param.Type
		assert.Equal(t, TypeKindEllipsis, typ.Kind)
		child := typ.Child
		assert.Equal(t, TypeKindStar, child.Kind)
		child = child.Child
		assert.Equal(t, TypeKindIdent, child.Kind)
		assert.Equal(t, child.Name, "int")
		assert.Nil(t, child.Child)
	})

	t.Run("array", func(t *testing.T) {
		f := testParseType(t, `a []int`)

		// act
		param := Parse(f, nil, "awesomepkg")[0]

		// assert
		assert.Equal(t, "a", param.Name)
		typ := param.Type
		assert.Equal(t, TypeKindArray, typ.Kind)
		child := typ.Child
		assert.Equal(t, TypeKindIdent, child.Kind)
		assert.Equal(t, child.Name, "int")
		assert.Nil(t, child.Child)
	})

	t.Run("star array", func(t *testing.T) {
		f := testParseType(t, `a *[]int`)

		// act
		param := Parse(f, nil, "awesomepkg")[0]

		// assert
		assert.Equal(t, "a", param.Name)
		typ := param.Type
		assert.Equal(t, TypeKindStar, typ.Kind)
		child := typ.Child
		assert.Equal(t, TypeKindArray, child.Kind)
		child = child.Child
		assert.Equal(t, TypeKindIdent, child.Kind)
		assert.Equal(t, child.Name, "int")
		assert.Nil(t, child.Child)
	})

	t.Run("array star", func(t *testing.T) {
		f := testParseType(t, `a []*int`)

		// act
		param := Parse(f, nil, "awesomepkg")[0]

		// assert
		assert.Equal(t, "a", param.Name)
		typ := param.Type
		assert.Equal(t, TypeKindArray, typ.Kind)
		child := typ.Child
		assert.Equal(t, TypeKindStar, child.Kind)
		child = child.Child
		assert.Equal(t, TypeKindIdent, child.Kind)
		assert.Equal(t, child.Name, "int")
		assert.Nil(t, child.Child)
	})

	t.Run("func", func(t *testing.T) {
		f := testParseType(t, `a func(m int, d bool) error`)

		// act
		param := Parse(f, nil, "awesomepkg")[0]

		// assert
		assert.Equal(t, "a", param.Name)
		typ := param.Type
		assert.Equal(t, TypeKindFunc, typ.Kind)
		assert.Nil(t, typ.Child)

		m := param.Type.Params[0]
		assert.Equal(t, TypeKindIdent, m.Type.Kind)
		assert.Equal(t, "int", m.Type.Name)

		d := param.Type.Params[1]
		assert.Equal(t, TypeKindIdent, d.Type.Kind)
		assert.Equal(t, "bool", d.Type.Name)

		e := param.Type.Results[0]
		assert.Equal(t, TypeKindIdent, e.Type.Kind)
		assert.Equal(t, "error", e.Type.Name)
	})
}
