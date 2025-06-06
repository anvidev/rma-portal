export type User = {
	id: number
	name: string
	email: string
	isActive: boolean
	inserted: string
}

export type Contact = {
	name: string
	company: string
	email: string
	phone: string
	street: string
	city: string
	zip: string
	country: string
}

export type Ticket = {
	id: string
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
	ticket_id: number
	status: string
	initiator: string
	external_comment: string
	internal_comment?: string
	inserted: string
	files?: File[]
}

export type File = {
	id: number
	file_name: string
	file_url: string
	file_domain: 'tickets' | 'logs'
	reference_id: string
	mime_type: string
	inserted: string
}

export interface TicketWithLogs extends Ticket {
	logs: Log[]
}
