package broadcasting

import "github.com/renanmedina/dcp-broadcaster/internal/accounts"

type MessageService interface {
	Send(message string, user accounts.User) error
	Broadcast(message string) error
}
