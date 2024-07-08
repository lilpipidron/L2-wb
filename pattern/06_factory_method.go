package pattern

import "fmt"

// Применимость: Когда заранее неизвестны типы и зависимости объектов, с которыми должен работать код
// Плюсы: OCP + SRP
// Минусы: Для каждого типа создаваемого объекта необходимо создавать новый подкласс и фабричный метод

type SomethingForFactoryMethod interface {
	DoSomething()
}

type SomethingForFactoryMethod1 struct{}
type SomethingForFactoryMethod2 struct{}

func (s *SomethingForFactoryMethod1) DoSomething() {
	fmt.Println("SomethingForFactoryMethod1 do something")
}

func (s *SomethingForFactoryMethod2) DoSomething() {
	fmt.Println("SomethingForFactoryMethod2 do something")
}

type SomethingCreator interface {
	CreateSomething() SomethingForFactoryMethod
}

type SomethingCreator1 struct{}
type SomethingCreator2 struct{}

func (s *SomethingCreator1) CreateSomething() SomethingForFactoryMethod {
	return &SomethingForFactoryMethod1{}
}

func (s *SomethingCreator2) CreateSomething() SomethingForFactoryMethod {
	return &SomethingForFactoryMethod2{}
}

func Create(somethingCreator SomethingCreator) SomethingForFactoryMethod {
	return somethingCreator.CreateSomething()
}
