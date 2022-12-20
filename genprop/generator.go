package genprop

import (
	"fmt"
	"go/ast"
	"go/token"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"github.com/hidori/go-tools/astutil"
	"github.com/pkg/errors"
)

type GeneratorConfig struct {
	TagName    string
	Initialism []string
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
	switch decl.Tok {
	case token.IMPORT:
		return []ast.Decl{decl}, nil

	case token.TYPE:
		return g.fromTypeDecl(decl)

	default:
		return nil, nil
	}
}

func (g *Generator) fromTypeDecl(decl *ast.GenDecl) ([]ast.Decl, error) {
	var decls []ast.Decl

	for _, spec := range decl.Specs {
		_decls, err := g.fromTypeSpec(spec.(*ast.TypeSpec))
		if err != nil {
			return nil, errors.Wrap(err, "fail to g.fromTypeSpec()")
		}

		decls = append(decls, _decls...)
	}

	return decls, nil
}

func (g *Generator) fromTypeSpec(spec *ast.TypeSpec) ([]ast.Decl, error) {
	structType, ok := spec.Type.(*ast.StructType)
	if !ok {
		return nil, nil
	}

	decls, err := g.fromStructType(spec.Name.Name, structType)
	if err != nil {
		return nil, errors.Wrap(err, "fail to g.fromStructType()")
	}

	return decls, nil
}

func (g *Generator) fromStructType(structName string, structType *ast.StructType) ([]ast.Decl, error) {
	decls, err := g.fromFieldList(structName, structType.Fields)
	if err != nil {
		return nil, errors.Wrap(err, "fail to g.fromFieldList()")
	}

	return decls, nil
}

func (g *Generator) fromFieldList(structName string, fieldList *ast.FieldList) ([]ast.Decl, error) {
	var decls []ast.Decl

	for _, field := range fieldList.List {
		_decls, err := g.fromField(structName, field)
		if err != nil {
			return nil, errors.Wrap(err, "fail to g.fromField()")
		}

		decls = append(decls, _decls...)
	}

	return decls, nil
}

func (g *Generator) fromField(structName string, field *ast.Field) ([]ast.Decl, error) {
	directives, err := g.fromTag(field.Tag)
	if err != nil {
		return nil, errors.Wrap(err, "fail to g.fromTag()")
	}

	var decls []ast.Decl

	for _, directive := range directives {
		if directive == "-" || directive == "" {
			return nil, nil
		}

		var decl ast.Decl

		switch directive {
		case "get":
			decl = g.newGetterFuncDecl(structName, field)

		case "set":
			decl = g.newSetterFuncDecl(structName, field)

		default:
			return nil, fmt.Errorf("invalid tag value '%s'", directive)
		}

		decls = append(decls, decl)
	}

	return decls, nil
}

func (g *Generator) fromTag(tag *ast.BasicLit) ([]string, error) {
	if tag == nil {
		return nil, nil
	}

	tagValue, err := strconv.Unquote(tag.Value)
	if err != nil {
		return nil, errors.Wrapf(err, "fail to strconv.Unquote() s=%s", tag.Value)
	}

	directives := strings.Split(strings.Trim(reflect.StructTag(tagValue).Get(g.config.TagName), " "), ",")

	return directives, nil
}

func (g *Generator) newGetterFuncDecl(structName string, field *ast.Field) ast.Decl {
	recv := astutil.NewFieldList(
		[]*ast.Field{
			astutil.NewField(
				[]*ast.Ident{
					astutil.NewIdent("t"),
				},
				astutil.NewStarExpr(astutil.NewIdent(structName)),
			),
		},
	)
	name := astutil.NewIdent(
		fmt.Sprintf("Get%s", g.prepareFieldName(field.Names[0].Name)),
	)
	funcType := astutil.NewFuncType(
		nil,
		nil,
		astutil.NewFieldList([]*ast.Field{
			astutil.NewField(nil, field.Type),
		}),
	)
	body := astutil.NewBlockStmt([]ast.Stmt{
		astutil.NewReturnStmt([]ast.Expr{
			astutil.NewSelectorExpr(astutil.NewIdent("t"), astutil.NewIdent(field.Names[0].Name)),
		}),
	})

	return astutil.NewFuncDecl(recv, name, funcType, body)
}

func (g *Generator) newSetterFuncDecl(structName string, field *ast.Field) ast.Decl {
	recv := astutil.NewFieldList(
		[]*ast.Field{
			astutil.NewField(
				[]*ast.Ident{
					astutil.NewIdent("t"),
				},
				astutil.NewStarExpr(astutil.NewIdent(structName)),
			),
		},
	)
	name := astutil.NewIdent(
		fmt.Sprintf("Set%s", g.prepareFieldName(field.Names[0].Name)),
	)
	funcType := astutil.NewFuncType(
		nil,
		astutil.NewFieldList([]*ast.Field{
			astutil.NewField(
				[]*ast.Ident{
					ast.NewIdent("v"),
				},
				field.Type,
			),
		}),
		nil,
	)
	body := astutil.NewBlockStmt([]ast.Stmt{
		astutil.NewAssignStmt(
			[]ast.Expr{
				astutil.NewSelectorExpr(astutil.NewIdent("t"), astutil.NewIdent(field.Names[0].Name)),
			},
			token.ASSIGN,
			[]ast.Expr{
				astutil.NewIdent("v"),
			},
		),
	})

	return astutil.NewFuncDecl(recv, name, funcType, body)
}

var camelHeadPattern = regexp.MustCompile(`^[a-z]+`)

func (g *Generator) prepareFieldName(name string) string {
	head := camelHeadPattern.FindString(name)

	if len(head) > 0 {
		head = strings.Title(head)

		for _, s := range g.config.Initialism {
			if strings.Title(s) == head {
				head = strings.ToUpper(head)

				break
			}
		}

		name = camelHeadPattern.ReplaceAllString(name, head)
	}

	return name
}
