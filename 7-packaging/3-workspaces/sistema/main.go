package main

import (
	"github.com/google/uuid"
	"github.com/phelipperibeiro/goexpert/7-packaging/3-workspaces/math"
)

func main() {
    m := math.NewMath(1, 2)
    println(m.Add())
    println(uuid.New().String())
}


// go work init ./math ./sistema