package pattern

import "fmt"

// Применимость: Запрос -> объект, чтобы удобно передавать в качестве аргументов при вызове методов, логировать и тд
// Плюсы: Реализация отложенного запуска операций, сбор сложных команд из простых
// Минусы: Усложнение кода программы из-за введения большого количества доп. классов

type Command interface {
	Execute()
}

type TextEditor struct {
	text string
}

func (c *TextEditor) Copy() {
	fmt.Println("copy command", c.text)
}

func (c *TextEditor) Paste(text string) {
	fmt.Println("paste command", text)
}

type CopyCommand struct {
	editor *TextEditor
}

func (c *CopyCommand) Execute() {
	c.editor.Copy()
}

type PasteCommand struct {
	editor *TextEditor
	text   string
}

func (c *PasteCommand) Execute() {
	c.editor.Paste(c.text)
}
