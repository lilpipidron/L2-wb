package pattern

import "fmt"

// Паттерн очень похож на стратегию, но в отличие от него состояния могут знать друг о друге
// Применимость: когда объект в зависимости от состояния меняю функционал: кнопки телефона в зависимости от крана делают разные дейтсвия: громкость, фото и тд
// Плюсы: Избавляемся от кучи операторов ветвления
// Минусы: Если состояний мало и они редко изменяются, это просто усложнение кода

type State interface {
	DoSomething(str string) State
}

type State1 struct{}
type State2 struct{}
type State3 struct{}

func (state State1) DoSomething(str string) State {
	fmt.Println(str)
	return &State2{}
}

func (state State2) DoSomething(str string) State {
	fmt.Println(str + "1")
	return &State3{}
}

func (state State3) DoSomething(str string) State {
	fmt.Println(str + "2")
	return state
}

type SomethingWithState struct {
	state State
}

func NewSomethingWithState(state State) *SomethingWithState {
	return &SomethingWithState{state: state}
}

func (smth *SomethingWithState) DoSomething() {
	smth.state = smth.state.DoSomething("string")
}

func (smth *SomethingWithState) ChangeState(newState State) {
	smth.state = newState
}
