package productions

import (
    "compiler-go/scanner"
    "compiler-go/core"
)

type Prog struct {
    ds declseq;
    ss stmtseq;
}

func (p *Prog) Parse(scnr *scanner.Scanner) {
    // program must start with the keyword program
    scnr.CurrentToken().ThrowErrorIfNot(core.Program)

    scnr.NextToken()
    if () {

    }
}

func (p *Prog) Print() {
}

func (p *Prog) Exec() {
}

