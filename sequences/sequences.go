package main

import "fmt"

func addOne(a int) int {
	return a + 1
}

func square(a int) int {
	return a * a
}

func double(slice []int) {
	slice = append(slice, slice...)
}

func mapSlice(f func(a int) int, slice []int) {
	for index, value := range slice {
		slice[index] = f(value)
	}
}

// ? Question 3c
func mapArray(f func(a int) int, array *[3]int) {
	for index, value := range *array {
		ptr := &array[index]
		*ptr = f(value)
	}
}

func main() {
	var intSlice = []int{1, 2, 3}
	mapSlice(addOne, intSlice)
	fmt.Println("intSlice:", intSlice)

	var intArray = [3]int{1, 2, 3}
	// ? Question 3b
	// ? An array is a sequence of (immutable) elements so
	// ? reassigning the storage location with mapped values
	// ? should "update" the array with new values
	mapArray(addOne, &intArray)
	fmt.Println("intArray:", intArray)

	// ? Question 3c
	// ? You can no longer pass intArray to mapArray since it is
	// ? no longer the right size or type
	intSlice = []int{1, 2, 3, 4, 5}
	// intArray = [5]int{1, 2, 3, 4, 5}

	// ? Question 3d
	newSlice := intSlice[1:3]
	mapSlice(square, newSlice)
	fmt.Println("intSlice after newSlice square:", intSlice)
	fmt.Println("newSlice after square:", newSlice)

	// ? Question 3e
	double(intSlice)
	fmt.Println("intSlice after double:", intSlice)

	// ? Question 3f
	// ? Observations:
	// * Arrays seem to be passed in by value to functions
	// * with no reference to the memory location that it
	// * was stored in. To update, use pointers :D

	// * Slices seem to pass references to memory locations
	// * similar to pointers in C which allows overwriting of
	// * what is stored there
}
