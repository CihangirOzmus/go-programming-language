package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Book struct {
	id int
	title string
}

func (b *Book) setTitle(title string)  {
	b.title = title

	// go automatically dereference b
	// (*b).title = title

	// that one below does not work
	// *b.title = title
}

func change(num *int)  {
	*num = 100
}


func testPtrSlice(ptrs *[]*int)  {
	var sb strings.Builder
	values := *ptrs
	for i, val := range values {
		s := strconv.Itoa(*val)
		if i < len(values) - 1 {
			sb.WriteString(s)
			sb.WriteString("-")
		} else {
			sb.WriteString(s)
		}
	}
	fmt.Println(sb.String())
	sb.Reset()
}

func main() {
	// 1st example
	a := 2
	change(&a)
	fmt.Println(a)

	// 2nd example
	mybook := Book{id: 1, title: "Dummy Title"}
	mybook.setTitle("Name of The Wind")
	fmt.Println(mybook)

	// 3rd example
	x := 10
	y := &x
	z := &y
	fmt.Printf("%T %T %T\n", x, y, z) //int *int(ptr) **int(ptr to ptr)
	fmt.Println(x, y ,z) // val addr addr
	fmt.Println(x, *y ,*z) // val val addr
	fmt.Println(x, *y ,**z) // val val val

	// 4rd example
	a, b, c := 1, 2, 3
	ptrSlice := &[]*int{&a, &b, &c}
	testPtrSlice(ptrSlice)
}
