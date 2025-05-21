package mailer

import "embed"

//go:embed "templates"
var templates embed.FS

var (
	// templates
	TicketCreated = "ticket_created.tmpl"
)

type Mailer interface {
	Send(to []string, email string, data any) error
}
