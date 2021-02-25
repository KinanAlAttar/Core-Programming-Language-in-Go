package main

import (
    "compiler-go/scanner"
    "compiler-go/core"
    "os"
    "fmt"
    "strconv"
)

func main() {
    S := scanner.NewScanner(os.Args[1])

    for (S.CurrentToken() != core.Eof &&
            S.CurrentToken() != core.Error) {

        fmt.Print(S.CurrentToken().String())
        if S.CurrentToken() == core.Id {
            val := S.GetId()
            fmt.Print("[" + val + "]")
        } else if S.CurrentToken() == core.Const {
          var val int = S.GetConst()
          fmt.Print("[" + strconv.Itoa(val) + "]")
        }
        fmt.Print("\n")
        S.NextToken()
    }
}
