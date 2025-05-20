export type User = {
	id: number
	name: string
	email: string
	isActive: boolean
	inserted: string
}

export type Contact = {
	name: string
	email: string
	phone: string
	street: string
	city: string
	zip: string
	country: string
}

export type Ticket = {
	id: number
	status: string
	categories: string[]
	issue: string
	model: string | null
	serial_number: string | null
	quote: 'yes' | 'no'
	warranty: 'yes' | 'no' | 'unknown'
	inserted: string
	updated: string
	sender: Contact
	billing: Contact
}

export type Log = {
	id: number
	ticketId: number
	status: string
	initiator: string
	external_comment: string
	internal_comment?: string
	inserted: string
}

export interface TicketWithLogs extends Ticket {
	logs: Log[]
}
