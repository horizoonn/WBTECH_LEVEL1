package main

import (
	"fmt"
	"reflect"
)

func identifyType(v interface{}) string {
	switch v.(type) {
	case bool:
		return "bool"
	case int:
		return "int"
	case string:
		return "string"
	default:
		if reflect.TypeOf(v).Kind() == reflect.Chan {
			return "chan"
		}
		return "unknown"
	}
}


func main() {
	a := "cat"
	b := 1242
	c := true
	d := make(chan int)
	e := 314.31

	values := []interface{}{a, b, c, d, e}

	for _, v := range values {
		fmt.Println(identifyType(v))
	}
}
