package main

import "fmt"

func main() {
	// assign array
	var student [15]string // or

	student[0] = "w"
	student[1] = "q"
	// abbreviate assign array
	employee := [15]string{"a", "b", "c"}

	fmt.Printf("%s \n %v\n", student, employee)
	for i := 0; i < len(employee); i++ {
		fmt.Print(employee[i] + ",")
	}

	arr3 := [...]int{1, 2, 3}
	fmt.Printf("\n %v \n", arr3)
	/*
		//index out of range
		var arr4 []int
		arr4[0] = 10
		arr4[1] = 20
		arr4[2] = 30
		fmt.Printf("\n%v", arr4)
	*/
	fmt.Println("\n***Slice in Golang***")
	var my_slice = []string{"Geeks", "for", "Geeks"}
	fmt.Println(my_slice)

	// another example of slice
	// Creating an array
	arr := []string{"Geeks", "for", "Geeks", "GFG"}

	// Creating slices from the given array
	var my_slice_1 = arr[1:2]
	my_slice_2 := arr[0:]
	my_slice_3 := arr[:2]
	my_slice_4 := arr[:]

	// Display the result
	fmt.Println("My Array: ", arr)
	fmt.Println("My Slice 1: ", my_slice_1)
	fmt.Println("My Slice 2: ", my_slice_2)
	fmt.Println("My Slice 3: ", my_slice_3)
	fmt.Println("My Slice 4: ", my_slice_4)
}
