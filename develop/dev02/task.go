package main

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

// В решении используется паттерн цепочка обязанностей, для удобного расширения

var lastChar rune

type Chain interface {
	Unpack(r rune, result *strings.Builder) error
	SetNext(chain Chain)
}

type BaseChain struct {
	next Chain
}

func (bc *BaseChain) SetNext(chain Chain) {
	bc.next = chain
}

func (bc *BaseChain) Unpack(r rune, result *strings.Builder) error {
	if bc.next == nil {
		return errors.New("invalid string")
	}
	return bc.next.Unpack(r, result)
}

type SymbolChain struct {
	BaseChain
}

func (sc *SymbolChain) Unpack(r rune, result *strings.Builder) error {
	if !unicode.IsLetter(r) {
		return sc.BaseChain.Unpack(r, result)
	}
	result.WriteRune(r)
	lastChar = r
	return nil
}

type DigitChain struct {
	BaseChain
}

func (dc *DigitChain) Unpack(r rune, result *strings.Builder) error {
	if !unicode.IsDigit(r) {
		return dc.BaseChain.Unpack(r, result)
	}
	if lastChar == 0 {
		return errors.New("invalid string")
	}
	amount, _ := strconv.Atoi(string(r))
	for range amount - 1 {
		result.WriteRune(lastChar)
	}
	lastChar = 0
	return nil
}

type EscapeChain struct {
	BaseChain
	pos *int
	str *[]rune
}

func (ec *EscapeChain) Unpack(r rune, result *strings.Builder) error {
	if r != 92 {
		return ec.BaseChain.Unpack(r, result)
	}
	if *ec.pos+1 >= len(*ec.str) {
		return errors.New("invalid escape sequence")
	}
	*ec.pos++
	lastChar = (*ec.str)[*ec.pos]
	result.WriteRune(lastChar)
	return nil
}

func Unpack(str string) (string, error) {
	var builder strings.Builder
	runes := []rune(str)
	var i int

	sc := &SymbolChain{}
	dc := &DigitChain{}
	ec := &EscapeChain{
		pos: &i,
		str: &runes,
	}
	dc.SetNext(ec)
	sc.SetNext(dc)

	for i < len(runes) {
		if err := sc.Unpack(runes[i], &builder); err != nil {
			return "", err
		}
		i++
	}

	return builder.String(), nil
}
