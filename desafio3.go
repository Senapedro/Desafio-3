package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Second) // Cria um ticker que dispara um evento a cada segundo
	defer ticker.Stop()                   // Garante que o ticker seja parado quando a função main terminar

	for {
		select {
		case <-ticker.C: // Quando o ticker disparar um evento, imprime "ping" e espera meio segundo antes de imprimir "pong"
			fmt.Println("ping")
			time.Sleep(time.Millisecond * 500)
			fmt.Println("pong")
		}
	}
}
