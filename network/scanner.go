package network // Mantenha o pacote como "network"

import (
	"bufio"
	"bytes"
	"fmt"
	"os/exec"
	"regexp"
	"strings"
	"wifi-monitor/notifier" // Importa o pacote notifier
)

var connectedDevices = make(map[string]string) // Armazena dispositivos conectados atualmente

// ListenIPs escaneia a rede e atualiza a lista de dispositivos conectados
func ListenIPs() {
	cmd := exec.Command("cmd", "/C", "arp -a") // Para Windows
	// cmd := exec.Command("arp", "-a") // Para Linux/Mac
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Println("Erro ao executar o comando arp:", err)
		return
	}

	// Processa a saída para capturar endereços IP e MAC
	currentDevices := make(map[string]string) // Dispositivos no escaneamento atual (MAC -> IP)
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
				currentDevices[mac] = ipAddress

				// Notifica se um novo dispositivo foi encontrado
				if connectedDevices[mac] == "" {
					notifier.Notify("Novo Dispositivo Conectado", fmt.Sprintf("IP: %s | MAC: %s", ipAddress, mac)) // Chama a função do pacote notifier
				}
			}
		}
	}

	// Detecta desconexões
	detectDisconnections(currentDevices)

	// Atualiza o estado atual dos dispositivos conectados
	connectedDevices = currentDevices

	if err := scanner.Err(); err != nil {
		fmt.Println("Erro ao processar saída:", err)
	}
}

// detectDisconnections verifica se algum dispositivo foi desconectado
func detectDisconnections(currentDevices map[string]string) {
	for mac, ip := range connectedDevices {
		if currentDevices[mac] == "" {
			notifier.Notify("Dispositivo Desconectado", fmt.Sprintf("IP: %s | MAC: %s", ip, mac)) // Chama a função do pacote notifier
		}
	}
}
