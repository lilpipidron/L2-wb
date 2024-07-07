```go
package main
 
import (
    "fmt"
)
 
func test() (x int) {
    defer func() {
        x++
    }()
    x = 1
    return
}
 
 
func anotherTest() int {
    var x int
    defer func() {
        x++
    }()
    x = 1
    return x
}
 
 
func main() {
    fmt.Println(test())
    fmt.Println(anotherTest())
}
```

Ответ: 2, 1, в первом примере defer выполнится после того как вызовется return, но до фактического возврата и он изменит значение именованной переменной, в случае 2 функции, будет точно также, но уже нельзя будет повлиять на x, так как он уже был передан в return