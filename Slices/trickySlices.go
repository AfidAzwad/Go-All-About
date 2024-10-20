package main

import "fmt"

func main() {

	//////// Example 1
	fmt.Println("\n---------a---------:")
	a := make([]int, 3)
	fmt.Println("Length of a :", len(a))  // 3
	fmt.Println("Capacity of a:", cap(a)) // 3
	fmt.Println("Address of a:", &a[0])

	fmt.Println("\n---------b---------:")
	b := append(a, 4)
	fmt.Println("Length of b:", len(b))
	fmt.Println("Capacity of b:", cap(b)) // 6 cause Go internally allocates a new array with a larger capacity to hold both the original elements of a and the new element.
	fmt.Println("Address of b:", &b[0])

	fmt.Println("\n---------c---------:")
	c := append(a, 5)
	fmt.Println("Length of c:", len(c))
	fmt.Println("Capacity of c:", cap(c)) // 6 cause Go internally allocates a new array with a larger capacity to hold both the original elements of a and the new element.
	fmt.Println("Address of c:", &c[0])

	fmt.Println("a : ", a)
	fmt.Println("b : ", b)
	fmt.Println("c : ", c)

	/*
	   Explanation:
	   In Example 1, slice 'a' has an initial capacity of 3. When you append to it (both in 'b' and 'c'), Go needs to allocate a new array because the current capacity (3) is not sufficient. As a result, a new array is created, and both 'b' and 'c' get new arrays with a larger capacity (6). They point to different memory locations, so appending to 'b' or 'c' doesn't affect 'a'.
	*/

	//////// Example 2
	fmt.Println("\nExample 2:---")

	fmt.Println("\n---------d---------:")

	d := make([]int, 3, 8)
	fmt.Println("Length of d :", len(d))
	fmt.Println("Capacity of d:", cap(d))
	fmt.Println("Address of d:", &d[0])

	fmt.Println("\n---------e---------:")

	e := append(d, 4)
	fmt.Println("Length of e:", len(e))
	fmt.Println("Capacity of e:", cap(e))
	fmt.Println("Address of e:", &e[0])

	fmt.Println("\n---------f---------:")
	f := append(d, 5)
	fmt.Println("Length of f:", len(f))
	fmt.Println("Capacity of f:", cap(f))
	fmt.Println("Address of f:", &f[0])

	fmt.Println("d : ", d)
	fmt.Println("e : ", e)
	fmt.Println("f : ", f)

	/*
		   Explanation:
		   In Example 2, slice 'd' has an initial capacity of 8 (more than its length of 3). This means that when you append to it (both in 'e' and 'f'),
			the append operation doesn't need to allocate a new array because there's already enough capacity (8) to accommodate the new elements. Both 'e' and 'f'
			point to the same underlying array as 'd', so appending to 'e' or 'f' modifies the same array. That's why 'e' and 'f' share the same memory address and underlying data.

		   The key difference between Example 1 and Example 2 is the initial capacity. In Example 1, 'a' had insufficient capacity (3), which triggered new allocations,
			while in Example 2, 'd' had a larger capacity (8), so no new allocation was needed, and the slices share the same memory.
	*/
}
