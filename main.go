package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/Gerardo02/interpreter-go/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is the cera programming language!\n", user.Username)
	fmt.Println("Feel free to type in commands!")
	repl.Start(os.Stdin, os.Stdout)
}

// WASM build
//
// func interpret(this js.Value, args []js.Value) interface{} {
// 	input := args[0].String()
//
// 	var out strings.Builder
// 	repl.EvalLine(input, &out)
//
// 	return js.ValueOf(out.String())
// }
//
// func main() {
// 	c := make(chan struct{}, 0)
//
// 	js.Global().Set("interpret", js.FuncOf(interpret))
//
// 	<-c // keep Go WASM alive
// }
