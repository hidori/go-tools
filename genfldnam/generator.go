package genfldnam

import (
	"fmt"
	"go/ast"
	"go/token"
	"reflect"
	"strconv"
	"strings"

	"github.com/barweiss/go-tuple"
	"github.com/hidori/go-tools/astutil"
	"github.com/hidori/go-tools/linqutil"
	"github.com/makiuchi-d/linq/v2"
	"github.com/pkg/errors"
)

type GeneratorConfig struct {
	TagName        string
	Skewer         string
	AllNames       bool
	AllNamesSuffix string
}

type Generator struct {
	config *GeneratorConfig
}

func NewGenerator(config *GeneratorConfig) *Generator {
	return &Generator{
		config: config,
	}
}

func (g *Generator) Generate(fset *token.FileSet, f *ast.File) ([]ast.Decl, error) {
	e1 := linq.FromSlice(f.Decls)
	e2 := linq.Select(e1, func(v ast.Decl) (*ast.GenDecl, error) {
		return linqutil.AsOrEmpty[*ast.GenDecl](v)
	})
	e3 := linq.Where(e2, func(v *ast.GenDecl) (bool, error) {
		return v != nil, nil
	})
	e4 := linq.SelectMany(e3, g.fromGenDecl, linqutil.PassThrough[ast.Decl])

	decls, err := linq.ToSlice(e4)
	if err != nil {
		return nil, errors.Wrap(err, "fail to linq.ToSlice()")
	}

	return decls, nil
}

func (g *Generator) fromGenDecl(decl *ast.GenDecl) (linq.Enumerable[ast.Decl], error) {
	switch decl.Tok {
	case token.TYPE:
		return g.fromTypeGenDecl(decl)

	default:
		return linq.Empty[ast.Decl](), nil
	}
}

func (g *Generator) fromTypeGenDecl(decl *ast.GenDecl) (linq.Enumerable[ast.Decl], error) {
	e1 := linq.FromSlice(decl.Specs)
	e2 := linq.Select(e1, func(v ast.Spec) (*ast.TypeSpec, error) {
		return linqutil.AsOrEmpty[*ast.TypeSpec](v)
	})
	e3 := linq.Where(e2, func(v *ast.TypeSpec) (bool, error) {
		return v != nil, nil
	})
	e4 := linq.SelectMany(e3, g.fromTypeSpec, linqutil.PassThrough[ast.Decl])

	return e4, nil
}

func (g *Generator) fromTypeSpec(typeSpec *ast.TypeSpec) (linq.Enumerable[ast.Decl], error) {
	structType, ok := typeSpec.Type.(*ast.StructType)
	if !ok {
		return linq.Empty[ast.Decl](), nil
	}

	decls, err := g.fromFieldList(typeSpec.Name.Name, structType.Fields)
	if err != nil {
		return nil, errors.Wrap(err, "fail to g.fromFieldList()")
	}

	return decls, nil
}

func (g *Generator) fromFieldList(structName string, fieldList *ast.FieldList) (linq.Enumerable[ast.Decl], error) {
	e1 := linq.FromSlice(fieldList.List)
	e2 := linq.Select(e1, func(v *ast.Field) (*tuple.T2[string, string], error) {
		return g.fromField(structName, v)
	})
	e3 := linq.Where(e2, func(v *tuple.T2[string, string]) (bool, error) {
		return v != nil, nil
	})

	pair, err := linq.Aggregate(e3, &tuple.T2[[]ast.Spec, []ast.Expr]{}, g.fromStringStringPair)
	if err != nil {
		return nil, errors.Wrap(err, "fail to linq.Aggregate()")
	}

	return g.fromSpecExprPair(structName, pair), nil
}

func (g *Generator) fromField(structName string, field *ast.Field) (*tuple.T2[string, string], error) {
	directive := g.fromTag(field.Tag)
	if len(directive) == 0 || directive == "-" {
		return nil, nil
	}

	if directive != "+" {
		return nil, fmt.Errorf("invalid tag value '%s'", directive)
	}

	return &tuple.T2[string, string]{
		V1: fmt.Sprintf("%s%s%s", structName, g.config.Skewer, field.Names[0].Name),
		V2: strconv.Quote(field.Names[0].Name),
	}, nil
}

func (g *Generator) fromTag(tag *ast.BasicLit) string {
	if tag == nil {
		return ""
	}

	tagValue, _ := strconv.Unquote(tag.Value)
	directive := strings.Trim(reflect.StructTag(tagValue).Get(g.config.TagName), " ")

	return directive
}

func (g *Generator) fromStringStringPair(acc *tuple.T2[[]ast.Spec, []ast.Expr], v *tuple.T2[string, string]) (*tuple.T2[[]ast.Spec, []ast.Expr], error) {
	value := astutil.NewBasicLit(
		token.STRING,
		v.V2,
	)
	spec := astutil.NewValueSpec(
		[]*ast.Ident{
			astutil.NewIdent(v.V1),
		},
		[]ast.Expr{
			value,
		},
	)
	pair := tuple.New2(
		append(acc.V1, spec),
		append(acc.V2, value),
	)

	return &pair, nil
}

func (g *Generator) fromSpecExprPair(structName string, pair *tuple.T2[[]ast.Spec, []ast.Expr]) linq.Enumerable[ast.Decl] {
	decls := linq.FromSlice([]ast.Decl{})

	if len(pair.V1) < 1 {
		return decls
	}

	decls = linqutil.Append(decls, ast.Decl(astutil.NewGenDecl(token.CONST, pair.V1)))

	if g.config.AllNames {
		decl := astutil.NewGenDecl(token.VAR, []ast.Spec{
			astutil.NewValueSpec(
				[]*ast.Ident{
					astutil.NewIdent(fmt.Sprintf("%s%s", structName, g.config.AllNamesSuffix)),
				},
				[]ast.Expr{
					astutil.NewCompositeLit(astutil.NewArrayType(astutil.NewIdent("string")), pair.V2),
				},
			),
		})
		decls = linqutil.Append(decls, ast.Decl(decl))
	}

	return decls
}
