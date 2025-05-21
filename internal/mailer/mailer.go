package mailer

import "embed"

//go:embed "templates"
var templates embed.FS

var (
	// templates
	TicketCreatedCustomer = "ticket_created_customer.tmpl"
	TicketCreatedSkancode = "ticket_created_skancode.tmpl"
)

type Mailer interface {
	Send(to []string, email string, data any) error
}
