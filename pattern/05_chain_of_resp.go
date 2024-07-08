package pattern

import "fmt"

// Применимость: В системе  изначально заложено k проверок, потом захотели добавить еще m, вместо добавления if else, просто создаем новый элемент цепочки и вклиниваем его внутрь
// Плюсы: OCP + SRP позволяет не нарушать
// Минусы: Запрос может остаться необработанным

type Chain interface {
	SetNext(chain Chain) Chain
	Handle(request string)
}

type BaseChain struct {
	next Chain
}

func (bc *BaseChain) SetNext(chain Chain) Chain {
	bc.next = chain
	return chain
}

func (bc *BaseChain) Handle(request string) {
	if bc.next != nil {
		bc.next.Handle(request)
	}
}

type Chain1 struct {
	BaseChain
}

type Chain2 struct {
	BaseChain
}

func (c *Chain1) Handle(request string) {
	if request == "Chain1" {
		fmt.Println("Chain1")
	} else {
		c.BaseChain.Handle(request)
	}
}

func (c *Chain2) Handle(request string) {
	if request == "Chain2" {
		fmt.Println("Chain2")
	} else {
		c.BaseChain.Handle(request)
	}
}
