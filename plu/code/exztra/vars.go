package main

import "fmt"

const (
	fi = iota + 6
	se = iota << 2
	th
)

func structuuu() {
	type usah struct {
		ID        int
		FirstName string
		LastName  string
	}

	var u usah
	u.ID = 10
	u.FirstName = "Mike"
	u.LastName = "Hawk"
	fmt.Println(u)

	u2 := usah{ID: 1,
		FirstName: "arthas",
		LastName:  "Menethil",
	}
	fmt.Println(u2)
}

func mappoo() {
	m := map[string]int{"foo": 42, "fook": 3}
	fmt.Println(m)
	fmt.Println(m["foo"])
	fmt.Println(m["fook"])

}

func arrSlo() {
	var arr [3]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	fmt.Println(arr)

	arr2 := [3]int{3, 2, 1}
	fmt.Println(arr2)

	slice := arr[:]

	arr[1] = 42
	slice[2] = 27
	fmt.Println(slice, arr)

	slice2 := []int{1, 2, 3, 4, 5}
	fmt.Println(slice2)

	slice2 = append(slice2, 4)
	fmt.Println(slice2)
}
func varslol() {
	var i int
	i = 42
	fmt.Println(i)

	var f float32 = 3.14
	fmt.Println(f)

	firstName := "Mike"
	fmt.Println(firstName)

	b := true
	fmt.Println(b)

	c := complex(3, 4)
	fmt.Println(c)

	r, im := real(c), imag(c)
	println(r, im)

	var fName *string = new(string)
	fmt.Println(fName)
	*fName = "HAWK"
	fmt.Println(*fName)

	secName := "ARt"
	fmt.Println(secName)

	ptr := &secName
	secName = "safasfaashfajshfaklshfalsjhflaskjhfals"
	fmt.Println(ptr, *ptr)
}

func constaL() {
	const pi = 3.145
	fmt.Println(pi)

	const c int = 3
	fmt.Println(c + 3)
	//fmt.Println(c + 1.2)
	println(fi, se, th)
}

func main() {
	varslol()
	constaL()
	arrSlo()
	mappoo()
	structuuu()
}
