```go
package main
 
type customError struct {
     msg string
}
 
func (e *customError) Error() string {
    return e.msg
}
 
func test() *customError {
     {
         // do something
     }
     return nil
}
 
func main() {
    var err error
    err = test()
    if err != nil {
        println("error")
        return
    }
    println("ok")
}
```
Ответ: "error". Все из-за устройства интерфейсов, в данном случае у нас будет лежать информация о типе customError, а как было сказано в 3 вопросе, чтобы был nil, надо чтобы оба поля внутри были nil, а в данном случае одно из полей содержит информацию о типе. 