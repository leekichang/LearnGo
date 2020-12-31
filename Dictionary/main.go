package main

import (
	"fmt"

	"github.com/leekichang/LearnGo/Dictionary/mydict"
)

func main() {
	fmt.Println("This is a update example")
	dictionary := mydict.Dictionary{"first": "First Word"}
	baseword := "hello"
	dictionary.Add(baseword, "First")
	err := dictionary.Update(baseword, "Second")
	if err != nil {
		fmt.Println(err)
	}
	word, _ := dictionary.Search(baseword)
	fmt.Println(word)
	fmt.Println("\nThis is a delete example")
	baseword2 := "Bye"
	dictionary.Add(baseword2, "bye")
	dictionary.Search(baseword2)
	dictionary.Delete(baseword2)
	word1, err1 := dictionary.Search(baseword2)
	if err1 != nil {
		fmt.Println(err1)
	} else {
		fmt.Println(word1)
	}
}
