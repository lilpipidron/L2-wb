```go
package main
 
import (
    "fmt"
    "os"
)
 
func Foo() error {
    var err *os.PathError = nil
    return err
}
 
func main() {
    err := Foo()
    fmt.Println(err)
    fmt.Println(err == nil)
}
```
Ответ: nil, false. Получаем false потому что интерфейсное значение в Go состоит из 2х частей: типа и значения. И переменная считается nil, только, если и тип, и значение внутри равны nil. А в данном случае err имеет тип *os.PathError 

```go
iface это обычный интерфейс. Содержит в себе itab - таблицу методов, который нам нужно реализовать.
eface это пустой интерфейс. Он не имеет методов и внутри нет itab.
```