package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	names := []string{"Han", "John", "Heli", "Beca"}
	for _, name := range names {
		go func() {
			fmt.Println(name)
		}()
		time.Sleep(time.Second)
	}
	runtime.GOMAXPROCS(1)
	runtime.Gosched()

}
