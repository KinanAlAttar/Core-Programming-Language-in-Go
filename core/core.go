package core

import (
	"fmt"
	"os"
)

type Core int

const (
    Program Core = iota
    Begin
    End
    New
    Define
    Extends
    Class
    Endclass
    Int
    Endfunc
    If
    Then
    Else
    While
    Endwhile
    Endif
    Semicolon
    Lparen
    Rparen
    Comma
    Assign
    Negation
    Or
    Equal
    Less
    Lessequal
    Add
    Sub
    Mult
    Input
    Output
    Const
    Id
    Eof
    Error
)

var Cores = []Core{
    Program,
    Begin,
    End,
    New,
    Define,
    Extends,
    Class,
    Endclass,
    Int,
    Endfunc,
    If,
    Then,
    Else,
    While,
    Endwhile,
    Endif,
    Semicolon,
    Lparen,
    Rparen,
    Comma,
    Assign,
    Negation,
    Or,
    Equal,
    Less,
    Lessequal,
    Add,
    Sub,
    Mult,
    Input,
    Output,
    Const,
    Id,
    Eof,
    Error,
}


func (c Core) String() string {
    return [...]string{
      "Program",
      "Begin",
      "End",
      "New",
      "Define",
      "Extends",
      "Class",
      "Endclass",
      "Int",
      "Endfunc",
      "If",
      "Then",
      "Else",
      "While",
      "Endwhile",
      "Endif",
      "Semicolon",
      "Lparen",
      "Rparen",
      "Comma",
      "Assign",
      "Negation",
      "Or",
      "Equal",
      "Less",
      "Lessequal",
      "Add",
      "Sub",
      "Mult",
      "Input",
      "Output",
      "Const",
      "Id",
      "Eof",
      "Error",
    }[c]
}

func (c1 Core) ThrowErrorIfNot(c2 Core) {
    if c1 != c2 {
        fmt.Println("ERROR: Expected " + "\"" + c1.String() +
            " found " + "\"" + c2.String() + "\"")
        os.Exit(0)
    }
}
