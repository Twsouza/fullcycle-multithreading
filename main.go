package main

import (
	"context"
	"fmt"
	"os"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run *.go <CEP>")
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	ceps := BuildMultipleCEP()
	cep := os.Args[1]
	ShowCEP(ctx, cep, ceps)
}
