package main

import (
	"fmt"

	"github.com/phelipperibeiro/goexpert/7-packaging/1-starting/math"
)


func main() {
    m := math.NewMath(1, 2)
    m.C = 3
    fmt.Println(m.C)
    // fmt.Println(m.Add())
    // fmt.Println(math.X)
}