package main

import (
	"cow/interpreter"
	"os"
	"path/filepath"
)

func main() {
	filepath := filepath.Join("cow_examples", "hello.cow")
	source, _ := os.ReadFile(filepath)
	var i = interpreter.Interpreter()
	i.Interpret(string(source))
}
