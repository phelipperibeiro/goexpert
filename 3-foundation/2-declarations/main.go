package main

const a = "Hello, World!"

/**
 * Global variables
 */

var (
	b bool    = true
	c int     = 10
	d string  = "Wesley"
	e float64 = 1.2
)

func main() {
	/**
	 * Local variables
	 */
	a := "X" // string
	println(a)
}

func x() {
}

// go run 3-foundation/2-declarations/*.go
