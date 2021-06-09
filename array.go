package array

import (
	"fmt"
	"reflect"
)

func ArrayMain() {
	simpleArray()
	twoDimensionalArray()
}

func twoDimensionalArray(){
	td := [2][3] int {{1, 2, 5},{3, 4}}
	fmt.Println(td)

	var twoDarr [2][3] int
	for i := 0; i < 2; i++ { // i --0, 1
		for j:= 0; j < 3; j++{ // j --0, 1,2
			twoDarr[i][j] = i + j // two Darr [0][0]
		}

	}
}

func simpleArray(){
}

	var arr [5]int // [0, 1, 2, 3, 4]
	fmt.Printf("Value of array %v\n", arr)
	fmt.Printf("Type of array %T\n", arr)
	fmt.Println(reflect.TypeOf(arr).Kind())

	arr[4] = 4
	fmt.Println("Setting the value", arr[4])
	fmt.Println("len of arr = ", len(arr))

	arr2 := [5]int{4, 5, 2, 6, 8}
	fmt.Println("Value of aarr2", arr2)

}
}