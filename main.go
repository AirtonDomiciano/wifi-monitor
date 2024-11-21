package main

import (
	"fmt"
	"time"
	"wifi-monitor/network" // Altere para o caminho correto do seu pacote
)

func main() {

	for {
		fmt.Println("Escaneando a rede...")
		// network.GetGatewayDefault()
		network.ListenIPs()

		// Espera 10 segundos antes de escanear novamente
		time.Sleep(10 * time.Second)
	}
}
