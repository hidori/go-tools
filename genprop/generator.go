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
	"github.com/hidori/go-tools/linqutil"
	"github.com/makiuchi-d/linq/v2"
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
	case token.IMPORT:
		return linq.FromSlice([]ast.Decl{decl}), nil

	case token.TYPE:
		return g.fromTypeGenDecl(decl), nil

	default:
		return linq.Empty[ast.Decl](), nil
	}
}

func (g *Generator) fromTypeGenDecl(decl *ast.GenDecl) linq.Enumerable[ast.Decl] {
	e1 := linq.FromSlice(decl.Specs)
	e2 := linq.Select(e1, func(v ast.Spec) (*ast.TypeSpec, error) {
		return linqutil.AsOrEmpty[*ast.TypeSpec](v)
	})
	e3 := linq.Where(e2, func(v *ast.TypeSpec) (bool, error) {
		return v != nil, nil
	})
	e4 := linq.SelectMany(e3, g.fromTypeSpec, linqutil.PassThrough[ast.Decl])

	return e4
}

func (g *Generator) fromTypeSpec(typeSpec *ast.TypeSpec) (linq.Enumerable[ast.Decl], error) {
	structType, ok := typeSpec.Type.(*ast.StructType)
	if !ok {
		return linq.Empty[ast.Decl](), nil
	}

	return g.fromFieldList(typeSpec.Name.Name, structType.Fields), nil
}

func (g *Generator) fromFieldList(structName string, fieldList *ast.FieldList) linq.Enumerable[ast.Decl] {
	e1 := linq.FromSlice(fieldList.List)
	e2 := linq.SelectMany(e1,
		func(v *ast.Field) (linq.Enumerable[ast.Decl], error) {
			return g.fromField(structName, v)
		},
		linqutil.PassThrough[ast.Decl])

	return e2
}

func (g *Generator) fromField(structName string, field *ast.Field) (linq.Enumerable[ast.Decl], error) {
	directives := g.fromTag(field.Tag)
	if len(directives) == 0 || (len(directives) == 1 && (directives[0] == "" || directives[0] == "-")) {
		return linq.Empty[ast.Decl](), nil
	}

	e1 := linq.FromSlice(directives)
	e2 := linq.Select(e1, func(v string) (ast.Decl, error) {
		switch v {
		case "get":
			return g.newGetterFuncDecl(structName, field), nil

		case "set":
			return g.newSetterFuncDecl(structName, field), nil

		default:
			return nil, fmt.Errorf("invalid tag value '%s'", v)
		}
	})

	return e2, nil
}

func (g *Generator) fromTag(tag *ast.BasicLit) []string {
	if tag == nil {
		return []string{}
	}

	t1, _ := strconv.Unquote(tag.Value)
	t2 := reflect.StructTag(t1).Get(g.config.TagName)

	return strings.Split(t2, ",")
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
