package main

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

// В решении используется паттерн цепочка обязанностей, для удобного расширения

var lastChar rune

// Chain - Интерфейс наших звеньев цепочки
type Chain interface {
	Unpack(r rune, result *strings.Builder) error
	SetNext(chain Chain)
}

// BaseChain - базовая реализация наших звеньев
type BaseChain struct {
	next Chain
}

// SetNext - добавления звена в цепочку
func (bc *BaseChain) SetNext(chain Chain) {
	bc.next = chain
}

// Unpack - попытка передачи обработки следующему, если конкретная реализация не смогла распаковать
func (bc *BaseChain) Unpack(r rune, result *strings.Builder) error {
	if bc.next == nil {
		return errors.New("invalid string")
	}
	return bc.next.Unpack(r, result)
}

// LetterChain - звено для обычных букв
type LetterChain struct {
	BaseChain
}

// Unpack - попытка распаковать букву
func (sc *LetterChain) Unpack(r rune, result *strings.Builder) error {
	if !unicode.IsLetter(r) {
		return sc.BaseChain.Unpack(r, result)
	}
	result.WriteRune(r)
	lastChar = r
	return nil
}

// DigitChain - звено для цифр
type DigitChain struct {
	BaseChain
}

// Unpack - пытаемся правильно распаковать цифру
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

// EscapeChain - звено для escape последовательностей
type EscapeChain struct {
	BaseChain
	pos *int
	str *[]rune
}

// Unpack - пытаемся правильно распаковать escape последовательность
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

// Unpack - функция, которая получает исходную строку и пытается ее распаковать с помощью цепочки
func Unpack(str string) (string, error) {
	var builder strings.Builder
	runes := []rune(str)
	var i int

	sc := &LetterChain{}
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
