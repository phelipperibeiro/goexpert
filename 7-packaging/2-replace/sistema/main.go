package main

import "github.com/phelipperibeiro/goexpert/7-packaging/2-replace/math"

func main() {
    m := math.NewMath(1, 2)
    println(m.Add())
}

// used when you want to replace a package that is not deployed yet and you want to test it locally
// go mod edit -replace github.com/phelipperibeiro/goexpert/7-packaging/2-replace/math=../math