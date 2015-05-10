package main

import "fmt"

func main() {
	//https://tour.golang.org/moretypes/7

	slice := make([]int,3, 5)
	fmt.Printf("slice len: %d, and cap: %d \n",len(slice),cap(slice))
    for i, v := range slice {
        fmt.Printf("var slice - the index: %d, has this value: %d \n",i,v)
    }

    fmt.Println("Assign values with append:")

    slice[0] = 1
    slice[1] = 2
    slice[2] = 3

    slice = append(slice,4)
    slice = append(slice,5)
    slice = append(slice,6)
    slice = append(slice,7)
    slice = append(slice,8)


    for i, v := range slice {
        fmt.Printf("var slice - the index: %d, has this value: %d \n",i,v)
    }

    fmt.Println("What the point of that slices has capacity?")



	p := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("p ==", p)

	for i := 0; i < len(p); i++ {
		fmt.Printf("p[%d] == %d\n", i, p[i])
	}


	fmt.Println("p[1:4] ==", p[1:4])

	// missing low index implies 0
	fmt.Println("p[:3] ==", p[:3])

	// missing high index implies len(s)
	fmt.Println("p[4:] ==", p[4:])


	a := make([]int, 5)
	printSlice("a", a)
	b := make([]int, 0, 5)
	printSlice("b", b)
	c := b[:2]
	printSlice("c", c)
	d := c[2:5]
	printSlice("d", d)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
