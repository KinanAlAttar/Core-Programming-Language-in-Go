package scanner

import (
	"bufio"
	"compiler-go/core"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

type pair struct {
    token core.Core
    value string
}

type Scanner struct {
    tokens []pair
    br *bufio.Reader
    currentTokenVal string
}

const (
    min_const int = 0
    max_const int = 1023
)

var (
    specials = []string{";","(",")",",","=","!","==","<","<=","+","-","*"}

    keywords = []string{
    "program","begin","end","new","int","define","endfunc",
    "class","extends","endclass","if","then","else","while",
    "endwhile","endif","or","input","output"}
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func NewScanner(filename string) *Scanner {
    f, e := os.Open(filename)
    check(e)

    scnr := Scanner {[]pair{}, bufio.NewReader(f), ""}

    firstToken := findToken(&scnr)

    p := pair {firstToken, scnr.currentTokenVal}
    scnr.tokens = append(scnr.tokens, p)

    return &scnr
}

func (scnr *Scanner) NextToken() {
    token := findToken(scnr)

    p := pair {token, scnr.currentTokenVal}
    scnr.tokens = append(scnr.tokens, p)
}

func (scnr *Scanner) CurrentToken() core.Core {
    if len(scnr.tokens) > 0 {
        return scnr.tokens[len(scnr.tokens) - 1].token
    }
    // return error otherwise
    fmt.Println("ERROR: Failed to find current token")
    return core.Error
}

func (scnr *Scanner) GetConst() int {
    result := 0
    if scnr.CurrentToken() == core.Const {
        c, e := strconv.Atoi(scnr.currentTokenVal)
        if e != nil {
            fmt.Println("ERROR: Failed scanning int")
            panic(e)
        }

        // flush the currentTokenVal
        scnr.currentTokenVal = ""
        result = c
    }
    return result
}

func (scnr *Scanner) GetId() string {
    val := ""
    if scnr.CurrentToken() == core.Id {
        val = scnr.currentTokenVal

        // flush currentTokenVal
        scnr.currentTokenVal = ""
    }
    return val
}


func isInvalidInput(char rune) bool {
    return (!isLetterOrDigit(char) &&
            !isSpecialChar(char)   &&
            !unicode.IsSpace(char))
}

func isSpecialChar(char rune) bool {
    for _, s := range specials {
        if s == string(char) {
            return true
        }
    }
    return false
}

func isLetterOrDigit(char rune) bool {
    return (
        unicode.IsDigit(char) ||
        unicode.IsLetter(char))
}

func isKeyword(token string) bool {
    for _,s := range keywords {
        if token == s {
            return true
        }
    }
    return false
}

func isId(token string) bool {
    if _,e := regexp.MatchString(
      "([a-zA-Z])([a-zA-Z0-9])*", token); e == nil {
        return true
    }
    return false
}

func isConst(token string) bool {
    if  c, e := strconv.Atoi(token); e == nil {
        return c <= max_const && c >= min_const
    }
    return false
}

func keywordToCore(token string) core.Core {
    for _,c := range core.Cores {
        if c.String() == strings.Title(token) {
            return c;
        }
    }
    return core.Error
}

func specCharToCore(token string) core.Core {
    switch token {
        case ";":
            return core.Semicolon
        case "(":
            return core.Lparen
        case ")":
            return core.Rparen
        case ",":
            return core.Comma
        case "=":
            return core.Assign
        case "!":
            return core.Negation
        case "==":
            return core.Equal
        case "<":
            return core.Less
        case "<=":
           return core.Lessequal
        case "+":
            return core.Add
        case "-":
            return core.Sub
        case "*":
            return core.Mult
        default:
            return core.Error
    }
}

func findToken(scnr *Scanner) core.Core {
    for {
        char,_,e := scnr.br.ReadRune();

        if e == io.EOF {
            return core.Eof
        }
        // check if character is invalid
        if isInvalidInput(char) {
            fmt.Printf("ERROR: Invalid symbol \"%s\"\n", string(char))
            return core.Error
        }

        // check if character is a letter
        if unicode.IsLetter(char) {
            token := ""
            token += string(char)

            for {
                char,_,e := scnr.br.ReadRune()
                if (e == io.EOF || isSpecialChar(char) ||
                    unicode.IsSpace(char)) {

                    if (isKeyword(token)) {
                        check(scnr.br.UnreadRune())
                        return keywordToCore(token)
                    } else if isId(token) {
                        check(scnr.br.UnreadByte())
                        scnr.currentTokenVal = token
                        return core.Id
                    }
                } else if isInvalidInput(char) {
                    fmt.Println("ERROR: Invalid symbol" + "\"" +
                        string(char) +   "\"")
                    return core.Error
                }

                token += string(char)
            }
        }

        // check if character is a const
        if unicode.IsDigit(char) {
            token := ""
            token += string(char)

            for {
                char,_,e := scnr.br.ReadRune()
                if (e == io.EOF || isSpecialChar(char) ||
                    unicode.IsSpace(char) || unicode.IsLetter(char)) {

                    if isConst(token) {
                        check(scnr.br.UnreadRune());
                        scnr.currentTokenVal = token
                        return core.Const
                    } else {
                        fmt.Println("ERROR: Constant must be within [0, 1023]")
                        return core.Error
                    }
                } else if isInvalidInput(char) {
                    fmt.Println("ERROR: Invalid symbol" + "\"" + string(char) +
                        "\"")
                    return core.Error
                }

                token += string(char)
            }
        }

        // check if character is a spec char
        if isSpecialChar(char) {
            token := ""
            token += string(char)

            nxtCh,_,_ := scnr.br.ReadRune()
            if string(nxtCh) == "=" {
                return specCharToCore(token + string(nxtCh))
            } else if (isInvalidInput(nxtCh)) {
                fmt.Println("ERROR: Invalid symbol" + "\"" + string(char) +
                        "\"")
                return core.Error
            } else {
                scnr.br.UnreadRune()
                return specCharToCore(token)
            }
        }
    }
}
