package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	bookHotel(ctx)
}

func bookHotel(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("Reserva de hotel cancelada. Limite de tempo excedido.")
		return

	case <-time.After(1 * time.Second):
		fmt.Println("Hotel reservado.")
	}
}
