package productions

import (
	"compiler-go/core"
	"compiler-go/scanner"
	"fmt"
	"strings"
)

type DeclSeq struct {
    DeclList []Decl
}

func NewDeclSeq() *DeclSeq {
    ds := DeclSeq {}
    return &ds
}

func (ds *DeclSeq) Parse(scnr *scanner.Scanner) {
    // must begin with decl
    d := NewDecl()
    d.Parse(scnr)
    ds.DeclList = append(ds.DeclList, d)

    // if next token is int call parse recursively
    scnr.NextToken()
    if (scnr.CurrentToken() == core.Int) {
        ds.Parse(scnr)
    }
}

func (ds *DeclSeq) Print(indent int) {
    for _,d := range ds.DeclList {
        fmt.Println(strings.Repeat("\t", indent))
        d.Print()
    }
}
