package productions

import (
    "compiler-go/core"
	"compiler-go/scanner"
	"fmt"
)

type Stmt struct {
    asgn   AssignStmt
    ifstmt IfStmt
    loop   LoopStmt
    in     InStmt
    out    OutStmt
    d      Decl
}

func NewStmt() *Stmt {
    s := Stmt {}
    return &s
}

func (s *Stmt) Parse(scnr *scanner.Scanner) {
    token := scnr.CurrentToken()
    switch token {
        case core.If:
            s.ifstmt = NewIfStmt()
            s.ifstmt.Parse(scnr)
        case core.Id:
            s.asgn = NewAssignStmt()
            s.asgn.Parse(scnr)
        case core.While
            s.loop = NewLoopStmt()
            s.loop.Parse(scnr)
        case 
    }
}
