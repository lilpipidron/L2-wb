package pattern

// Применимость: Когда есть большой сложный объект и его проще собрать пошагово
// Плюсы: Можем создавать сложные объекты пошагово, изолируем код сборки продукта от бизнес логики
// Минусы: Код становится сложнее из-за всего этого кода

type bigStruct struct {
	field1 string
	field2 string
	field3 string
	field4 int
	field5 float64
}

type Builder interface {
	SetField1(string) Builder
	SetField2(string) Builder
	SetField3(string) Builder
	SetField4(int) Builder
	SetField5(float64) Builder
	Build() *bigStruct
}

type BigStructBuilder struct {
	bigStruct bigStruct
}

func NewBigStructBuilder() *BigStructBuilder {
	return &BigStructBuilder{bigStruct: bigStruct{}}
}

func (b *BigStructBuilder) SetField1(field1 string) Builder {
	b.bigStruct.field1 = field1
	return b
}

func (b *BigStructBuilder) SetField2(field2 string) Builder {
	b.bigStruct.field2 = field2
	return b
}

func (b *BigStructBuilder) SetField3(field3 string) Builder {
	b.bigStruct.field3 = field3
	return b
}

func (b *BigStructBuilder) SetField4(field4 int) Builder {
	b.bigStruct.field4 = field4
	return b
}

func (b *BigStructBuilder) SetField5(field5 float64) Builder {
	b.bigStruct.field5 = field5
	return b
}

func (b *BigStructBuilder) Build() *bigStruct {
	return &b.bigStruct
}
