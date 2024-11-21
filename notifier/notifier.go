package notifier

import (
	"fmt"

	"github.com/go-toast/toast"
)

func Notify(title, message string) {
	notification := toast.Notification{
		AppID:   "WiFi Monitor",
		Title:   title,
		Message: message,
		Icon:    "",
	}
	err := notification.Push()
	if err != nil {
		fmt.Println("Erro ao enviar notificação:", err)
	}
}
