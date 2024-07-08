package flags

import "errors"

type Chain interface {
	SetNext(next Chain)
	Parse(args []string, pos *int) error
}

type BaseChain struct {
	next Chain
}

func (bc *BaseChain) SetNext(next Chain) {
	bc.next = next
}

func (bc *BaseChain) Parse(args []string, pos *int) error {
	if bc.next != nil {
		return bc.next.Parse(args, pos)
	}
	return errors.New("failed parse, unknown arg: " + args[*pos])
}
