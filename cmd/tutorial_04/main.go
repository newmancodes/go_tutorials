package main

import (
	"fmt"
	"time"
)

func main() {
	var intArr [3]int32 = [3]int32{1, 2, 3}
	fmt.Println(intArr[0])
	fmt.Println(intArr[1:3])

	fmt.Println(&intArr[0])
	fmt.Println(&intArr[1])
	fmt.Println(&intArr[2])

	var intSlice []int32 = []int32{4, 5, 6}
	fmt.Printf("The length is %v with capacity %v.\n", len(intSlice), cap(intSlice))
	fmt.Println(intSlice)
	intSlice = append(intSlice, 7)
	fmt.Printf("The length is %v with capacity %v.\n", len(intSlice), cap(intSlice))
	fmt.Println(intSlice)

	var intSlice2 []int32 = []int32{8, 9}
	intSlice = append(intSlice, intSlice2...)
	fmt.Println(intSlice)

	var intSlice3 []int32 = make([]int32, 3, 8)
	fmt.Printf("The length is %v with capacity %v.\n", len(intSlice3), cap(intSlice3))

	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)

	var myMap2 = map[string]uint8{"Adam": 23, "Sarah": 45}
	fmt.Println(myMap2["Adam"])
	var age, ok = myMap2["NotPresentSoDefaultOfValueReturned"]
	if ok {
		fmt.Printf("The age is %v.\n", age)
	} else {
		fmt.Printf("Invalid Name.\n")
	}

	for name, age := range myMap2 {
		fmt.Printf("Name: %v, Age: %v.\n", name, age)
	}

	for i, v := range intArr {
		fmt.Printf("Index: %v, Value: %v.\n", i, v)
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	var n int = 1000000
	var testSlice = []int{}
	var testSlice2 = make([]int, 0, n)

	fmt.Printf("Total time without preallocation: %v.\n", timeLoop(testSlice, n))
	fmt.Printf("Total time with preallocation: %v.\n", timeLoop(testSlice2, n))
}

func timeLoop(slice []int, n int) time.Duration {
	var t0 = time.Now()
	for len(slice) < n {
		slice = append(slice, 1)
	}
	return time.Since(t0)
}
