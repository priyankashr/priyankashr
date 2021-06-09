package array

import (
	"fmt"
)

func SlicesMain() {
	intro()

}
func intro() {
	// fmt.Printf("Value of array %v\n", arr)
	// fmt.Printf("Type of array %T\n", arr)
	// fmt.Println("Length of array %v\n", len(arr))
	// fmt.Println(reflect.TypeOf(arr).Kind())

	s := make([]string, 6) // make(type, length and capacity)
	fmt.Println("slices:", s)

	s[0] = "g"
	s[1] = "o"
	s[2] = "l"
	s[3] = "a"
	s[4] = "n"
	s[5] = "g"
	fmt.Println("slices:", s)
	fmt.Printf("len = %v, cap = %v\n", len(s), cap(s))

	s := s[2:4]
	fmt.Println("slices:", s)
	fmt.Printf("len = %v, cap = %v\n", len(s), cap(s))

	s = s[2:4]
	fmt.Println(slice)
	fmt.Printf("len = %v, cap = %v", len(s), cap(s))

	s = s[:4]
	fmt.Println(slice)
	fmt.Printf("len = %v, cap = %v", len(s), cap(s))

	s = s[2:]
	fmt.Println(slice)
	fmt.Printf("len = %v, cap = %v", len(s), cap(s))

	s = s[:]
	fmt.Println(slice)
	fmt.Printf("len = %v, cap = %v", len(s), cap(s))

	fmt.Println("slices:", s)
	fmt.Printf("len = %v, cap = %v", len(s), cap(s))
}
