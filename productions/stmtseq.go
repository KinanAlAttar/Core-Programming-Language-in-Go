package productions

import (
	"compiler-go/core"
	"compiler-go/scanner"
	"fmt"
	"strings"
)

type StmtSeq struct {
    StmtList []Stmt
}

func NewStmtSeq() *StmtSeq {
  ss := StmtSeq {}
  return &ss
}

func (ss *StmtSeq) Parse(scnr *scanner.Scanner) {
    // must begin with stmt
    s := NewStmt()
    s.Parse(scnr)
    ss.StmtList = append(ss.StmtList, s)

    // if next token is a stmt call parse recursively
    scnr.NextToken()
    if (isStmt(scnr.CurrentToken())) {
        ss.Parse(scnr)
    }
}

func (ss *StmtSeq) Print(indent int) {
    for _,stmt := range ss.StmtList {
        // add indentation accordingly
        fmt.Println(strings.Repeat("\t", indent))
        stmt.Print(indent)
    }
}

func isStmt(token core.Core) bool {
    return (
        token == core.If ||
        token == core.Id ||
        token == core.While ||
        token == core.Input ||
        token == core.Output ||
        token == core.Int)
}
