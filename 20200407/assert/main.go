package main

import (
	"fmt"
	"C"
)

func main() {
	fmt.Println("int:", f(1))

	fmt.Println("string:", f("aaa"))

	fmt.Println("float32:", f(float32(3.2)))
}
func t ()

func f(s interface{}) interface{} {
	if _, ok := s.(int); ok {
		return s
	}

	if _, ok := s.(string); ok {
		return s
	}

	if _, ok := s.(string); !ok{
		t()
	}

	return nil

}
