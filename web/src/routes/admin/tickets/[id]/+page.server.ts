import { API_URL } from '$lib/server/env'
import type { AdminTicket } from '$lib/types'
import type { PageServerLoad } from './$types'

export const load: PageServerLoad = async ({ cookies, params }) => {
	const id = params.id

	const response = await fetch(`${API_URL}/v1/admin/tickets/${id}`, {
		headers: {
			Authorization: `Bearer ${cookies.get('token')}`,
		},
	})

	if (!response.ok) {
		return {
			ticket: null,
		}
	}

	const { ticket } = await response.json()

	return {
		ticket: ticket as AdminTicket,
	}
}
