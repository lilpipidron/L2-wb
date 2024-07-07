package pattern

import "fmt"

// Применимость: добавление новых операций не изменяя структуры старых классов
// Плюсы: Проще добавлять операции в уже готовые классы: OCP соблюдается
// Минусы: не оправдан, если иерархия часто меняется (изменяется то, что мы хоти посещать)

type Visitor interface {
	DoSomethingForStruct1()
	DoSomethingForStruct2()
}

type Visitor1 struct{}
type Visitor2 struct{}

func (v *Visitor1) DoSomethingForStruct1() {
	fmt.Println("DoSomethingForStruct1")
}

func (v *Visitor1) DoSomethingForStruct2() {
	fmt.Println("DoSomethingForStruct2")
}

func (v *Visitor2) DoSomethingForStruct1() {
	fmt.Println("DoSomethingForStruct1_visitor2")
}

func (v *Visitor2) DoSomethingForStruct2() {
	fmt.Println("DoSomethingForStruct2_visitor2")
}

type Something interface {
	Accept(visitor Visitor)
}

type Struct1 struct{}
type Struct2 struct{}

func (s *Struct1) Accept(visitor Visitor) {
	visitor.DoSomethingForStruct1()
}

func (s *Struct2) Accept(visitor Visitor) {
	visitor.DoSomethingForStruct2()
}
