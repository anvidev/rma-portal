type User = {
	id: number
	name: string
	email: string
	isActive: boolean
	inserted: string
}

type Ticket = {
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
