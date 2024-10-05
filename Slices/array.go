package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6}
	fmt.Println("Array : ", arr)

	sliceOfArr := arr[1:4] // 1 to 3
	fmt.Println("Slice of Array : ", sliceOfArr)

	arrFixed := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Fixed Array : ", arrFixed)

	slicedArr := make([]int, 3, 5) // Length: 3, Capacity: 5

	fmt.Println("Sliced Array : ", slicedArr)
	fmt.Println("Sliced Array Length:", len(slicedArr))
	fmt.Println("Sliced Array Capacity:", cap(slicedArr))

	slicedArr = append(slicedArr, 5, 6, 7)
	fmt.Println("Appended Sliced Array:", slicedArr)
	fmt.Println("Appended Sliced Array Length:", len(slicedArr))
	fmt.Println("Appended Sliced Array Capacity:", cap(slicedArr))
}

/* make([]int, 3, 5):

[]int: This specifies that the slice will hold elements of type int.

3: This is the length of the slice. It means that the slice will initially contain 3 elements,
all set to their zero value (which is 0 for integers).

5: This is the capacity of the slice. It indicates that the underlying array can hold up to 5 elements.
The slice can grow beyond its initial length (up to the specified capacity) without needing to allocate a new array.
*/
