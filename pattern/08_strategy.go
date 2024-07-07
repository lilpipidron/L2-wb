package pattern

import "fmt"

// Применимость: когда один объект может использовать какой-то алгоритм в разных вариациях: навигатор (пешком, на велосипеде и тд)
// Плюсы: соблюдение OCP, упрощение тестирования
// Минусы: нужно знать разницу между стратегиями, чтобы выбрать подходящую

type Strategy interface {
	DoSomething()
}

type Strategy1 struct{}
type Strategy2 struct{}
type Strategy3 struct{}

func (strategy Strategy1) DoSomething() {
	fmt.Println("strategy 1.DoSomething")
}

func (strategy Strategy2) DoSomething() {
	fmt.Println("strategy 2.DoSomething")
}

func (strategy Strategy3) DoSomething() {
	fmt.Println("strategy 3.DoSomething")
}

type SomethingWithStrategy struct {
	strategy Strategy
}

func NewSomethingWithStrategy(strategy Strategy) *SomethingWithStrategy {
	return &SomethingWithStrategy{strategy: strategy}
}

func (smth *SomethingWithStrategy) DoSmth() {
	smth.strategy.DoSomething()
}
