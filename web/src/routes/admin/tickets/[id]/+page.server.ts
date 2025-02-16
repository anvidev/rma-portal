import type { PageServerLoad } from './$types'

export const load: PageServerLoad = async ({ cookies, params }) => {
	const id = params.id

	const response = await fetch(`http://localhost:8080/v1/admin/tickets/${id}`, {
		headers: {
			Authorization: `Bearer ${cookies.get('token')}`
		}
	})

	if (!response.ok) {
		return {
			ticket: null
		}
	}

	const { ticket } = await response.json()

	return {
		ticket: ticket as Ticket
	}
}
