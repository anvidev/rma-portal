export type User = {
	id: number
	name: string
	email: string
	isActive: boolean
	inserted: string
}

export type Ticket = {
	id: number
	status: string
	categories: string[]
	issue: string
	inserted: string
	updated: string
	sender: {
		name: string
		email: string
		address: string
	}
}

export interface AdminTicket extends Ticket {
	logs: {
		id: number
		ticketId: number
		status: string
		initiator: string
		external_comment: string
		internal_comment?: string
		inserted: string
	}[]
}
