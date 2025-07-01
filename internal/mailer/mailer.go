package mailer

import "embed"

//go:embed "templates"
var templates embed.FS

var (
	// mails
	TicketCreatedCustomer = "ticket_created_customer.tmpl"
	TicketClosedCustomer  = "ticket_closed_customer.tmpl"
	TicketCreatedSkancode = "ticket_created_skancode.tmpl"
)

type Mailer interface {
	Send(to []string, mail string, data any) error
}
