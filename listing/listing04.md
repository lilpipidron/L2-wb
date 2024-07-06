```go
package main
 
func main() {
    ch := make(chan int)
    go func() {
        for i := 0; i < 10; i++ {
            ch <- i
        }
    }()
 
    for n := range ch {
        println(n)
    }
}
```
Ответ: 0-9 и потом deadlock, потому что мы не закрыли канал, и основная горутина будет ждать пока в него что-то запишут, но этого никогда не случится