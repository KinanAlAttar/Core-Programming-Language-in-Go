package productions

import (
	"compiler-go/core"
	"compiler-go/scanner"
	"fmt"
)

type Prog struct {
    ds *DeclSeq;
    ss *StmtSeq;
}

func NewProg() *Prog {
    p := Prog {}
    return &p
}

func (p *Prog) Parse(scnr *scanner.Scanner) {
    // program must start with the keyword program
    scnr.CurrentToken().ThrowErrorIfNot(core.Program)

    // if next token is notbegin Parse decl-seq
    scnr.NextToken()
    if (scnr.CurrentToken() != core.Begin) {
        p.ds = NewDeclSeq()
        p.ds.Parse(scnr)
    }

    // if last token was a decl-seq then next token must be
    // begin, also there is no need to call NextToken() as
    // it was implicitly called in ds
    if (len(p.ds.DeclList) > 0) {
        scnr.CurrentToken().ThrowErrorIfNot(core.Begin)
        // grab next token which must be a stmt-seq
        scnr.NextToken()
        p.ss = NewStmtSeq()
        p.ss.Parse(scnr)
    } else { // else next token must be a stmt-seq
        scnr.NextToken()
        p.ss = NewStmtSeq()
        p.ss.Parse(scnr)
    }

    // next token must be end
    // no need to call NextToken() as it is called 
    // in ss
    scnr.CurrentToken().ThrowErrorIfNot(core.End)

    // next token must be eof
    scnr.CurrentToken().ThrowErrorIfNot(core.Eof)
}

func (p *Prog) Print() {
    fmt.Println("program")

    if (len(p.ds.DeclList) > 0) {
        ps.ds.Print(1)
    }

    fmt.Println("begin")

    if (len(p.ss.StmtList) > 0) {
        p.ss.Print(1)
    }

    fmt.Println("end")
}

func (p *Prog) Exec() {
}

