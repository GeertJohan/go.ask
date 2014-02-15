package main

import (
	"fmt"
	"github.com/GeertJohan/go.ask"
)

func main() {
	if ask.MustDefaultAsk(false, "Are you a go coder?") {
		fmt.Println("Great! Nice to meet you 😊")
	} else {
		fmt.Println("Owh,.. Well you should become one! 😼")
	}
}
