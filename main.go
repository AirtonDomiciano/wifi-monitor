package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os/exec"
	"regexp"
	"strings"
	"time"

	"github.com/go-toast/toast"
)

// Variável global para armazenar dispositivos conhecidos
var knownDevices = make(map[string]bool)

func main() {
	knownDevices := make(map[string]bool)

	for {
		fmt.Println("Escaneando a rede...")

		// Executa o comando arp no Windows
		cmd := exec.Command("cmd", "/C", "arp -a")
		var out bytes.Buffer
		cmd.Stdout = &out
		err := cmd.Run()

		fmt.Println("Monitorando o arp:", cmd)

		if err != nil {
			fmt.Println("Erro ao executar o comando arp:", err)
			return
		}

		// Processa a saída para capturar endereços IP e MAC
		scanner := bufio.NewScanner(&out)
		for scanner.Scan() {
			line := scanner.Text()

			// Procura padrões de endereço MAC
			macRegex := regexp.MustCompile(`([0-9A-Fa-f]{2}[:-]){5}([0-9A-Fa-f]{2})`)
			mac := macRegex.FindString(line)

			// Captura o endereço IP relacionado
			if mac != "" {
				fields := strings.Fields(line)
				if len(fields) > 0 {
					ipAddress := fields[0]

					if !knownDevices[mac] {
						knownDevices[mac] = true
						fmt.Printf("Novo dispositivo detectado: IP %s, MAC %s\n", ipAddress, mac)
						// Aqui você pode adicionar lógica para notificar

						notify(fmt.Sprintf("Novo dispositivo conectado!"), fmt.Sprintf("IP: %s | MAC: %s", ipAddress, mac))
					}
				}
			}
		}

		if err := scanner.Err(); err != nil {
			fmt.Println("Erro ao processar saída:", err)
			return
		}

		// Espera 30 segundos antes de escanear novamente
		time.Sleep(5 * time.Second)
	}
}

func addNewIPAddress(channel string, ipAddress string) {

}

func notify(title, message string) {
	notification := toast.Notification{
		AppID:   "WiFi Monitor",
		Title:   title,
		Message: message,
		Icon:    "", // Você pode adicionar um caminho para um ícone aqui
	}
	err := notification.Push()
	if err != nil {
		fmt.Println("Erro ao enviar notificação:", err)
	}
}
