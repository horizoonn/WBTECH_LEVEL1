package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	temps := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	step := 10.0

	groups := make(map[int][]float64, len(temps))
	for _, t := range temps {
		key := int(math.Floor(t / step) * step)
		groups[key] = append(groups[key], t)
	}

	keys := make([]int, 0, len(groups))
	for g := range groups {
		keys = append(keys, g)
	}
	sort.Ints(keys)

	for _, k := range keys {
		label := k
		if k < 0 {
			label = k + 10
		}
		vs := groups[k]
		sort.Float64s(vs)
		fmt.Printf("%d:{", label)
		for i, x := range vs {
			if i > 0 {
				fmt.Print(", ")
			}
			fmt.Printf("%.1f", x)
		}
		fmt.Println("}")
	}
}