package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := 5.35
	b := 'a'
	c := false
	d := 100
	e := "I am a code monkey"
	f := true
	g := 17.3
	h := 'c'
	i := "derp"

	fmt.Println(reflect.TypeOf(a))
	fmt.Println(reflect.TypeOf(b))
	fmt.Println(reflect.TypeOf(c))
	fmt.Println(reflect.TypeOf(d))
	fmt.Println(reflect.TypeOf(e))
	fmt.Println(reflect.TypeOf(f))
	fmt.Println(reflect.TypeOf(g))
	fmt.Println(reflect.TypeOf(h))
	fmt.Println(reflect.TypeOf(i))
}
