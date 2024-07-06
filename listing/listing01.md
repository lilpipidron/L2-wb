```go
package main

import (
	"fmt"
)

func main() {
	a := [5]int{76, 77, 78, 79, 80}
	var b []int = a[1:4]
	fmt.Println(b)
}
```

Ответ: 77, 78, 78
Потому что, когда мы делаем подобный срез мы правую границу не включаем