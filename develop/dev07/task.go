package main

import (
	"fmt"
	"sync"
	"time"
)

// идем в цикле по всем каналам и создаем для каждого рутинку, как только один из каналов закроется, наш основной канал
// тоже закроется, используем sync.Once, чтобы гарантировать только один вызов close(ch)
func or(channels ...<-chan interface{}) <-chan interface{} {
	ch := make(chan interface{})
	var once sync.Once
	for _, channel := range channels {
		go func(channel <-chan interface{}) {
			select {
			case <-channel:
				once.Do(func() {
					close(ch)
				})
			}
		}(channel)
	}

	return ch
}

func main() {
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()
	<-or(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)

	fmt.Printf("fone after %v", time.Since(start))

}
