package builder

import "fmt"

func main() {

	var director *Director

	// receive an gnb
	ranBuilder := GetBuilder(0)

	if director == nil {
		director = NewDirector(ranBuilder)
	}

	ran := director.BuildRan("127.0.0.1", 22)

	fmt.Println(ran)
}
