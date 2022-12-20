# printast

Prints AST (Abstract Syntax Tree)

## INSTALL

```bash
go install github.com/hidori/go-tools/cmd/printast@latest
```
## USAGE

```sh
usage: printast <FILE>
```

## EXAMPLES

### Example #1

`example.go`
```go
package example

import "fmt"

func main() {
    fmt.Println("Hello, world")
}
```

run CLI
```bash
printast example.go > example.txt
```

`example.txt`
```text
     0  *ast.File {
     1  .  Package: -
     2  .  Name: *ast.Ident {
     3  .  .  NamePos: -
     4  .  .  Name: "example"
     5  .  }
     6  .  Decls: []ast.Decl (len = 2) {
     7  .  .  0: *ast.GenDecl {
     8  .  .  .  TokPos: -
     9  .  .  .  Tok: import
    10  .  .  .  Lparen: -
    11  .  .  .  Specs: []ast.Spec (len = 1) {
    12  .  .  .  .  0: *ast.ImportSpec {
    13  .  .  .  .  .  Path: *ast.BasicLit {
    14  .  .  .  .  .  .  ValuePos: -
    15  .  .  .  .  .  .  Kind: STRING
    16  .  .  .  .  .  .  Value: "\"fmt\""
    17  .  .  .  .  .  }
    18  .  .  .  .  .  EndPos: -
    19  .  .  .  .  }
    20  .  .  .  }
    21  .  .  .  Rparen: -
    22  .  .  }
    23  .  .  1: *ast.FuncDecl {
    24  .  .  .  Name: *ast.Ident {
    25  .  .  .  .  NamePos: -
    26  .  .  .  .  Name: "main"
    27  .  .  .  .  Obj: *ast.Object {
    28  .  .  .  .  .  Kind: func
    29  .  .  .  .  .  Name: "main"
    30  .  .  .  .  .  Decl: *(obj @ 23)
    31  .  .  .  .  }
    32  .  .  .  }
    33  .  .  .  Type: *ast.FuncType {
    34  .  .  .  .  Func: -
    35  .  .  .  .  Params: *ast.FieldList {
    36  .  .  .  .  .  Opening: -
    37  .  .  .  .  .  Closing: -
    38  .  .  .  .  }
    39  .  .  .  }
(omit)
```
