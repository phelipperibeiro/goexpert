package main

func main() {

	a := 10
	b := 20

	if a > b {
		println("a é maior que b")
	} else {
		println("b é maior que a")
	}

	switch a {
	case 10:
		println("a é 10")
	case 20:
		println("a é 20")
	case 30:
		println("a é 30")
	case 40:
		println("a é 40")
	default:
		println("nenhum dos casos")
	}

	// ## OBS:
	// there is no ternary operator in Go
	// there is no else if in Go
}
