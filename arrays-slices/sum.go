package arrays

import "fmt"

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	var sums []int
	for _, numberSlice := range numbersToSum {
		sums = append(sums, Sum(numberSlice))
	}
	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}
	return sums
}

func main() {

	// arrays
	// the zero value is ready to use
	// meaning every position for the capacity of the array is 0
	var arr [4]int

	fmt.Printf("%v\n\r", arr[3] == 0) // true

	// the compiler can tell the size of the array for us during initialization
	arrShortDeclaration := [...]string{"Ann", "Marie", "Summer", "Riza"}
	fmt.Printf("%v\n\r", arrShortDeclaration)

	// any position can be mutated
	arr[0] = 5
	fmt.Printf("%v\n\r", arr[0] == 5)

	// can be resized to accomodate more values using append
	// arrOne := [...]int{}

	// slices
	// same declaration as array but leaving size unspecified
	var slice []int
	fmt.Printf("%v", slice)

	// can also be declared with make
	// make(t Type, size int, capacity int (optional))
	madeSlice := make([]int, 5, 10)
	fmt.Printf("%v", madeSlice)
	fmt.Printf("%v %v", len(slice), cap(slice))

	// creating a slice from an existing slice or array
	sliceOfArr := arr[0:3]
	fmt.Printf("%v", sliceOfArr)

	sliceOfSlice := madeSlice[2:4]
	fmt.Printf("%var", sliceOfSlice)

	sliceOfFullArr := arr[:]
	fmt.Printf("%var", sliceOfFullArr)
	// a slice is a pointer to a segment of memory
	// it is not like other values in memory, it keeps a ptr to the memory, a length and a capacity
	// because it is a pointer when we mutate values in the slice we are effectively mutating the original array
	// making slice operations as efficient as manipulating array indices

	sliceOfFullArr[1] = 8
	fmt.Printf("%v", sliceOfFullArr[1] == arr[1]) // true

	// length of a slice is the amount of elements it contains and it can't exceed capacity, capacity is the original capacity of the array
	// from the start of the slice
	nums := [...]int{1, 2, 3, 4, 5, 6, 7, 8}
	sliceOfNums := nums[2:5]
	fmt.Printf("%v len: %v cap: %v", sliceOfNums, len(sliceOfNums), cap(sliceOfNums)) // { 3, 4, 5 } len: 3 cap: 6 (6, 7, 8)

	// if we wanted to grow a slice to its capacity
	sliceOfNums = sliceOfNums[:cap(sliceOfNums)]
	fmt.Printf("%v len: %v cap: %v", sliceOfNums, len(sliceOfNums), cap(sliceOfNums)) // { 3, 4, 5, 6, 7, 8 } len: 6 cap: 6

}
