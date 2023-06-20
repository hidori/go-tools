package genfldnam

import (
	"fmt"
	"go/ast"
	"go/token"
	"reflect"
	"strconv"
	"strings"

	"github.com/hidori/go-tools/astutil"
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
	var decls []ast.Decl

	for _, decl := range f.Decls {
		if genDecl, ok := decl.(*ast.GenDecl); ok {
			_decls, err := g.fromGenDecl(genDecl)
			if err != nil {
				return nil, err
			}

			decls = append(decls, _decls...)
		}
	}

	return decls, nil
}

func (g *Generator) fromGenDecl(decl *ast.GenDecl) ([]ast.Decl, error) {
	var decls []ast.Decl

	for _, spec := range decl.Specs {
		typeSpec, ok := spec.(*ast.TypeSpec)
		if !ok {
			continue
		}

		_decls, err := g.fromTypeSpec(typeSpec)
		if err != nil {
			return nil, errors.Wrap(err, "fail to g.fromTypeSpec()")
		}

		decls = append(decls, _decls...)
	}

	return decls, nil
}

func (g *Generator) fromTypeSpec(typeSpec *ast.TypeSpec) ([]ast.Decl, error) {
	structType, ok := typeSpec.Type.(*ast.StructType)
	if !ok {
		return nil, nil
	}

	decls, err := g.fromStructType(typeSpec.Name.Name, structType)
	if err != nil {
		return nil, errors.Wrap(err, "fail to g.fromStructType()")
	}

	return decls, nil
}

func (g *Generator) fromStructType(structName string, structType *ast.StructType) ([]ast.Decl, error) {
	specs, values, err := g.fromFieldList(structName, structType.Fields)
	if err != nil {
		return nil, errors.Wrap(err, "fail to g.fromFieldList()")
	}

	if len(specs) < 1 {
		return nil, nil
	}

	decls := []ast.Decl{
		g.newFieldNameGenDecl(specs),
	}

	if !g.config.AllNames {
		return decls, nil
	}

	decls = append(decls, g.newAllFieldNamesGenDecl(structName, values))

	return decls, nil
}

func (g *Generator) fromFieldList(structName string, fieldList *ast.FieldList) ([]ast.Spec, []ast.Expr, error) {
	var (
		specs  []ast.Spec
		values []ast.Expr
	)

	for _, field := range fieldList.List {
		spec, value, err := g.fromField(structName, field)
		if err != nil {
			return nil, nil, errors.Wrap(err, "fail to g.fromField()")
		}

		if spec == nil {
			continue
		}

		specs = append(specs, spec)
		values = append(values, value)
	}

	return specs, values, nil
}

func (g *Generator) fromField(structName string, field *ast.Field) (ast.Spec, ast.Expr, error) {
	directive := g.fromTag(field.Tag)
	if len(directive) == 0 || directive == "-" {
		return nil, nil, nil
	}

	if directive != "+" {
		return nil, nil, fmt.Errorf("invalid tag value '%s'", directive)
	}

	value := g.newFieldNameExpr(field)
	spec := g.newFieldNameSpec(structName, field, value)

	return spec, value, nil
}

func (g *Generator) fromTag(tag *ast.BasicLit) string {
	if tag == nil {
		return ""
	}

	tagValue, _ := strconv.Unquote(tag.Value)
	directive := strings.Trim(reflect.StructTag(tagValue).Get(g.config.TagName), " ")

	return directive
}

func (g *Generator) newFieldNameGenDecl(specs []ast.Spec) *ast.GenDecl {
	return astutil.NewGenDecl(token.CONST, specs)
}

func (g *Generator) newAllFieldNamesGenDecl(structName string, values []ast.Expr) *ast.GenDecl {
	return astutil.NewGenDecl(token.VAR, []ast.Spec{
		astutil.NewValueSpec(
			[]*ast.Ident{
				astutil.NewIdent(fmt.Sprintf("%s%s", structName, g.config.AllNamesSuffix)),
			},
			[]ast.Expr{
				astutil.NewCompositeLit(astutil.NewArrayType(astutil.NewIdent("string")), values),
			},
		),
	})
}

func (g *Generator) newFieldNameExpr(field *ast.Field) ast.Expr {
	return astutil.NewBasicLit(
		token.STRING,
		strconv.Quote(field.Names[0].Name),
	)
}

func (g *Generator) newFieldNameSpec(structName string, field *ast.Field, value ast.Expr) ast.Spec {
	return astutil.NewValueSpec(
		[]*ast.Ident{
			astutil.NewIdent(fmt.Sprintf("%s%s%s", structName, g.config.Skewer, field.Names[0].Name)),
		},
		[]ast.Expr{
			value,
		},
	)
}
