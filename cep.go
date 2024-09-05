package main

import (
	"context"
	"errors"
	"fmt"
	"fullcycle-multithreading/brasilapi"
	"fullcycle-multithreading/viacep"
	"sync"
)

type CEP interface {
	Get(ctx context.Context, cep string) (string, error)
}

func BuildMultipleCEP() []CEP {
	return []CEP{
		&viacep.Viacep{},
		&brasilapi.BrasilAPI{},
	}
}

func ShowCEP(ctx context.Context, cep string, ceps []CEP) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	res := make(chan string)

	wg := sync.WaitGroup{}
	wg.Add(len(ceps))

	for _, c := range ceps {
		go func(c CEP) {
			defer wg.Done()

			r, err := c.Get(ctx, cep)
			if err != nil {
				if errors.Is(err, context.DeadlineExceeded) {
					fmt.Printf("Timeout: %s\n", err)
				}

				return
			}
			res <- r
		}(c)
	}

	go func() {
		wg.Wait()
		close(res)
	}()

	for r := range res {
		fmt.Printf("CEP %s => %s\n", cep, r)
		cancel()
	}
}
