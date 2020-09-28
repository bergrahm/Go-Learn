package exztra

import "fmt"

func sliceFunc() {
	slice := []int{1, 2, 3}
	for i, v := range slice {
		print(i, v)
	}

	println()

	wellKnownPorts := map[string]int{"http": 80, "https": 443}
	for k, va := range wellKnownPorts {
		fmt.Println(k, va)
	}
}

func loopTilCond() {
	var i int
	for i < 5 {
		fmt.Println(i)
		i++
	}
}

func roon() {
	loopTilCond()
	sliceFunc()
}
