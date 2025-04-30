package main

import (
	"github.com/anvidev/rma-portal/internal/apidoc"
	"github.com/anvidev/rma-portal/internal/store"
)

func (api *api) addDocs() {
	docs := api.documentation

	tickets := docs.AddTag("Tickets")
	tickets.AddEndpoint(apidoc.MethodGet, "/tickets", "Get all tickets",
		apidoc.WithQuery(
			"limit",
			"int64",
			apidoc.WithDefault(50),
			apidoc.WithMin(1),
			apidoc.WithMax(1000),
		))
	tickets.AddEndpoint(apidoc.MethodPost, "/tickets", "Create new ticket",
		apidoc.WithBody(store.Ticket{}),
	)
	tickets.AddEndpoint(apidoc.MethodPut, "/tickets/{id}", "Update a ticket",
		apidoc.WithQuery("id", "int64"),
		apidoc.WithBody(store.Ticket{}),
	)

	tickets.AddEndpoint(apidoc.MethodGet, "/users", "Get all users",
		apidoc.WithQuery(
			"limit",
			"int64",
			apidoc.WithDefault(50),
			apidoc.WithMin(1),
			apidoc.WithMax(1000),
		),
	)
}
