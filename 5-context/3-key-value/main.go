package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "token", "senha")
	bookHotel(ctx, "Blue Lagoon Resort")
}

func bookHotel(ctx context.Context, name string) {
	token := ctx.Value("token")
	fmt.Println(token)
	fmt.Printf("Hotel %s booked\n", name)
}

// go run 5-context/3-key-value/main.go
