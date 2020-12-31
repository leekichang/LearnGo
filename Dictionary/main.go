package main

import (
	"fmt"

	"github.com/leekichang/LearnGo/Dictionary/mydict"
)

func main() {
	dictionary := mydict.Dictionary{"first": "First Word"}
	definition, err1 := dictionary.Search("second")
	if err1 != nil {
		fmt.Println(err1)
	} else {
		fmt.Println(definition)
	}
	word := "hello"
	def := "Greeting"
	err2 := dictionary.Add(word, def)
	if err2 != nil {
		fmt.Println(err2)
	}
	hello, _ := dictionary.Search("hello")
	fmt.Println("found", word, "definition", hello)

	err3 := dictionary.Add(word, def)
	if err3 != nil {
		fmt.Println(err3)
	}
}
