package network

import (
	"fmt"
	"net"

	"github.com/vishvananda/netlink"
)

func GetGatewayDefault() {
	// Obtém as informações de rede
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("Erro ao obter interfaces:", err)
		return
	}

	// Itera sobre as interfaces
	for _, iface := range interfaces {
		// Ignora interfaces down ou loopback
		if iface.Flags&net.FlagUp == 0 || iface.Flags&net.FlagLoopback != 0 {
			continue
		}

		// Obtém as rotas para a interface
		addrs, err := iface.Addrs()
		if err != nil {
			fmt.Println("Erro ao obter endereços da interface:", err)
			continue
		}

		// Itera sobre os endereços
		for _, addr := range addrs {
			// Verifica se o endereço é do tipo IP
			if ipNet, ok := addr.(*net.IPNet); ok {
				// Obtém a rota padrão
				gateway := GetDefaultGateway(ipNet.IP)
				if gateway != nil {
					fmt.Printf("Interface: %s, Gateway: %s\n", iface.Name, gateway)
				}
			}
		}
	}
}

func GetDefaultGateway(ip net.IP) net.IP {
	// Obtém as rotas do sistema
	fmt.Println("getDefaultGateway IP:", ip)
	routes, err := netlink.RouteGet(ip)

	if err != nil {
		fmt.Println("Erro ao obter rotas:", err)
		return nil
	}

	// Retorna o gateway da primeira rota padrão
	if len(routes) > 0 {
		return routes[0].Gw
	}

	return nil
}
