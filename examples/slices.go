package main

import "fmt"

func main() {
	//https://tour.golang.org/moretypes/7

	slice := make([]int, 3, 5)
	fmt.Printf("slice len: %d, and cap: %d \n", len(slice), cap(slice))
	for i, v := range slice {
		fmt.Printf("var slice - the index: %d, has this value: %d \n", i, v)
	}

	fmt.Println("")

	fmt.Println("Assign values with append:")

	slice[0] = 1
	slice[1] = 2
	slice[2] = 3

	slice = append(slice, 4)
	slice = append(slice, 5)
	slice = append(slice, 6)
	slice = append(slice, 7)
	slice = append(slice, 8)

	for i, v := range slice {
		fmt.Printf("var slice - the index: %d, has this value: %d \n", i, v)
	}

	fmt.Println("")

	fmt.Println("One slice of 6 elements:")

	p := []int{1, 2, 3, 4, 5, 6}
	fmt.Println("p ==", p)

	for i := 0; i < len(p); i++ {
		fmt.Printf("p[%d] == %d\n", i, p[i])
	}

	fmt.Println("")

	fmt.Printf("p[1:4] of the %d shoulby [2 3 4] and we get this: %d \n", p, p[1:4])

	fmt.Printf("p[:3] of the %d shoulby [1 2 3] and we get this: %d \n", p, p[:3])

	fmt.Printf("p[4:] of the %d shoulby [5 6] and we get this: %d \n", p, p[4:])

	fmt.Printf("p[5:6] of the %d shoulby [6] and we get this: %d \n", p, p[5:6])

	fmt.Println("")

	fmt.Printf("p[1:4] means we start count from 2 to 4 \n")
	fmt.Printf("p[:3] means we start count from 1 to 3 \n")
	fmt.Printf("p[4:] means we start count from 5 to 6 \n")
	fmt.Printf("p[5:] means we start count from 6 to 6 \n")

	fmt.Println("")

	var a = make([]int, 5)
	fmt.Printf("var a = make([]int, 5) is a slice with a length of %d and capacity of %d \n", len(a), cap(a))
	fmt.Printf("var a; has this value: %d \n", a)

	fmt.Println("")

	var b = make([]int, 5, 10)
	fmt.Printf("var b = make([]int, 5, 10) is a slice with a length of %d and capacity of the %d \n", len(b), cap(b))
	fmt.Printf("var b; has this value: %d \n", b)

	b[0] = 1
	b[1] = 2
	b[2] = 3
	b[3] = 4
	b[4] = 5

	c := b[:10]

	fmt.Printf("var c; has this value: %d after give values to d[0] ... d[4] and set c := b[:10] \n", c)
	fmt.Printf("var c; is now a new slice with a length of %d and capacity of the %d \n", len(c), cap(c))

	c[5] = 6
	c[6] = 7
	c[7] = 8
	c[8] = 9
	c[9] = 10

	fmt.Printf("After give values to c[5] ... c[9]; c now has this value: %d\n", c)

	fmt.Println("")

	fmt.Printf("Expand a slice size: \n")
	var z = make([]int, 2, 7)
	fmt.Printf("var z = make([]int, 2, 7) is a slice with a length of %d and capacity of %d \n", len(z), cap(z))
	fmt.Printf("var z; has this value: %d \n", z)

	z[0] = 1
	z[1] = 2
	z = z[:7]
	fmt.Printf("after add values to z[0],z[1] and expand a slice size: z = z[:7] the slice var z; has the length of %d \n", len(z))
	fmt.Printf("var z; has now this value: %d \n", z)

	fmt.Println("")
	fmt.Println("Remove capacity of a slice:")
	fmt.Println("")

	var x = []int{1, 2, 3, 4, 5}
	var x1 = x[1:]
	var x2 = x[2:]
	var x3 = x[3:]
	var x4 = x[4:]

	fmt.Printf("var x = []int{1,2,3,4,5} is a slice with a length of %d and capacity of %d \n", len(x), cap(x))
	fmt.Printf("var x; has this value: %d \n", x)

	fmt.Printf("var x1 = x[1:] is a slice with a length of %d and capacity of %d \n", len(x1), cap(x1))
	fmt.Printf("var x1; has this value: %d \n", x1)

	fmt.Printf("var x2 = x[2:] is a slice with a length of %d and capacity of %d \n", len(x2), cap(x2))
	fmt.Printf("var x2; has this value: %d \n", x2)

	fmt.Printf("var x3 = x[3:] is a slice with a length of %d and capacity of %d \n", len(x3), cap(x3))
	fmt.Printf("var x3; has this value: %d \n", x3)

	fmt.Printf("var x4 = x[4:] is a slice with a length of %d and capacity of %d \n", len(x4), cap(x4))
	fmt.Printf("var x4; has this value: %d \n", x4)

	fmt.Println("")

	var y = make([]int, 5, 10)

	y[0] = 1
	y[1] = 2
	y[2] = 3
	y[3] = 4
	y[4] = 5

	var y1 = y[1:]
	var y2 = y[2:]
	var y3 = y[3:]
	var y4 = y[4:]

	fmt.Printf("var y = make([]int,5,10) is a slice with a length of %d and capacity of %d \n", len(y), cap(y))
	fmt.Printf("var y; has this value: %d \n", y)

	fmt.Printf("var x1 = x[1:] is a slice with a length of %d and capacity of %d \n", len(y1), cap(y1))
	fmt.Printf("var y1; has this value: %d \n", y1)

	fmt.Printf("var y2 = y[2:] is a slice with a length of %d and capacity of %d \n", len(y2), cap(y2))
	fmt.Printf("var y2; has this value: %d \n", y2)

	fmt.Printf("var y3 = x[3:] is a slice with a length of %d and capacity of %d \n", len(y3), cap(y3))
	fmt.Printf("var y3; has this value: %d \n", y3)

	fmt.Printf("var y4 = x[4:] is a slice with a length of %d and capacity of %d \n", len(y4), cap(y4))
	fmt.Printf("var y4; has this value: %d \n", y4)

	fmt.Println("")
	fmt.Println("Array and Slices")

	var arr1 [10]int

	arr1[0] = 1
	arr1[1] = 2
	arr1[2] = 3
	arr1[3] = 4
	arr1[4] = 5
	arr1[5] = 6
	arr1[6] = 7
	arr1[7] = 8
	arr1[8] = 9
	arr1[9] = 10

	var slice1 []int = arr1[2:]
	var slice2 []int = arr1[2:10]
	var slice3 []int = arr1[2:5]

	fmt.Println("arr1: ", arr1)
	fmt.Println("slice1: ", slice1)
	fmt.Println("slice2: ", slice2)
	fmt.Println("slice3: ", slice3)



}
