package main

import (
	"fmt"

	"github.com/leekichang/LearnGo/Dictionary/mydict"
)

func main() {
	dictionary := mydict.Dictionary{"first": "First Word"}
	definition, err := dictionary.Search("second")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(definition)
	}
}
