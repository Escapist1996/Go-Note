package main

import (
	"fmt"

	"github.com/Go-Note/Go/MockTest/service"
)

func main() {
	d := service.NewDB()

	g := service.NewGreeter(d, "en")
	fmt.Println(g.Greeter)

}
