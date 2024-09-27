package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/exp/slices"
)

type (
	Token  string
	Syntax string
)

const (
	OBJ_OP  Token = "OBJ_OP"
	OBJ_CLO Token = "OBJ_CLO"
	DOU_CLO Token = "DOU_QUO"
	COLON   Token = "COLON"
	COMMA   Token = "COMMA"
	ARR_OP  Token = "ARR_OP"
	ARR_CLO Token = "ARR_CLO"
)

type TokenMap struct {
	name   Token
	symbol rune
}

var ignore = []rune{' ', '\n', '\t'}

var lexer = []TokenMap{
	{OBJ_OP, '{'},
	{OBJ_CLO, '}'},
	{DOU_CLO, '"'},
	{COLON, ':'},
	{COMMA, ','},
	{ARR_OP, '['},
	{ARR_CLO, ']'},
}

func main() {
	if v, err := process(os.Args[1:]); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	} else {
		fmt.Println(v)
	}
}

func process(args []string) (bool, error) {
	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "Error: %v\n", "missing input file")
		os.Exit(1)
	}
	data, err := getFileData(args[0])
	if err != nil {
		return false, err
	}
	lex, err := lexData(data)
	fmt.Println(lex)
	return true, nil
}

func lexData(data []byte) ([]string, error) {
	tokens := []string{}
	lit := ""
	for _, s := range strings.TrimSpace(string(data)) {
		if slices.Index(ignore, s) != -1 {
			continue
		}
		idx := slices.IndexFunc(lexer, func(t TokenMap) bool { return t.symbol == s })
		if idx != -1 {
			if len(lit) > 0 {
				tokens = append(tokens, "STR="+lit)
				lit = ""
			}
			tokens = append(tokens, string(lexer[idx].name))
		} else {
			lit = lit + string(s)
		}
	}
	return tokens, nil
}

func getFileData(path string) ([]byte, error) {
	return os.ReadFile(path)
}
