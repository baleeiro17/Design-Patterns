package main

import (
	"Design-Patterns/builder"
	"fmt"
)

func main() {

	var director *builder.Director

	// receive an gnb
	ranBuilder := builder.GetBuilder(0)

	if director == nil {
		director = builder.NewDirector(ranBuilder)
	}

	ran := director.BuildRan("127.0.0.1", 22)

	fmt.Println(ran)
}
